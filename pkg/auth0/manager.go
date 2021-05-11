package auth0

import (
	"bytes"
	"encoding/json"
	"html/template"
	"os"

	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/aserto-dev/aserto-idp-seed/pkg/counter"
	"github.com/aserto-dev/aserto-idp-seed/pkg/data"
	"github.com/pkg/errors"
	"gopkg.in/auth0.v5/management"
)

const (
	counterInterval = 1
)

type Manager struct {
	config  *config.Auth0
	mgnt    *management.Management
	spew    bool
	exec    bool
	counter counter.Counter
}

// NewManager Auth0 management interface.
func NewManager(cfg *config.Auth0, filename string) *Manager {
	manager := Manager{
		config:  cfg,
		exec:    true,
		counter: counter.Counter{},
	}
	return &manager
}

// Init initialize management connection.
func (m *Manager) Init() error {
	mgnt, err := management.New(
		m.config.Domain,
		management.WithClientCredentials(
			m.config.ClientID,
			m.config.ClientSecret,
		),
	)
	if err != nil {
		return errors.Wrapf(err, "create management instance")
	}

	m.mgnt = mgnt

	return nil
}

// Spew output, default off
func (m *Manager) Spew(f bool) {
	m.spew = f
}

// Dryrun, no execution mode,default off
func (m *Manager) Dryrun(f bool) {
	m.exec = !f
}

// Seed, seed Auth0 with test user identities.
func (m *Manager) Seed(params *config.TemplateParams) error {

	r := bytes.NewBuffer(data.Users)

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(true)
	enc.SetIndent("", "  ")

	m.counter.Print(counter.Init)

	dec := json.NewDecoder(r)
	if _, err := dec.Token(); err != nil {
		return errors.Wrapf(err, "token close")
	}

	for dec.More() {
		var user *management.User
		if err := dec.Decode(&user); err != nil {
			return errors.Wrapf(err, "decode user")
		}

		b, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			return errors.Wrapf(err, "marshal user")
		}

		s := templatize(params, string(b))

		if err := json.Unmarshal([]byte(s), &user); err != nil {
			return errors.Wrapf(err, "unmarshal user")
		}

		if m.spew {
			_ = enc.Encode(user)
		}

		if m.userExists(*user.ID) {
			if err := m.updateUser(*user.ID, user); err != nil {
				continue
			}
		} else {
			if err := m.createUser(*user.ID, user); err != nil {
				continue
			}
		}

		m.counter.IncrRows()
		m.counter.Print(counterInterval)
	}

	m.counter.Print(counter.Last)

	return nil
}

// Reset, remove test user identities from Auth0.
func (m *Manager) Reset() error {
	r := bytes.NewBuffer(data.Users)

	m.counter.Print(counter.Init)

	dec := json.NewDecoder(r)
	if _, err := dec.Token(); err != nil {
		return errors.Wrapf(err, "token close")
	}

	for dec.More() {
		var user *management.User
		if err := dec.Decode(&user); err != nil {
			return errors.Wrapf(err, "decode user")
		}

		if err := m.deleteUser(*user.ID); err != nil {
			continue
		}

		m.counter.IncrRows()
		m.counter.Print(counterInterval)
	}

	m.counter.Print(counter.Last)

	return nil
}

func templatize(params *config.TemplateParams, s string) string {
	t, err := template.New("config").Parse(s)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, params)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func (m *Manager) userExists(id string) bool {
	if !m.exec {
		return true
	}

	if _, err := m.mgnt.User.Read("auth0|" + id); err != nil {
		return false
	}
	return true
}

func (m *Manager) createUser(id string, u *management.User) error {
	if !m.exec {
		return nil
	}

	if err := m.mgnt.User.Create(u); err != nil {
		m.counter.IncrError()
		return err
	}

	return nil
}

func (m *Manager) updateUser(id string, u *management.User) error {
	if !m.exec {
		return nil
	}

	// reset fields which cannot be changed
	u.ID = nil
	u.Password = nil

	if err := m.mgnt.User.Update("auth0|"+id, u); err != nil {
		m.counter.IncrError()
		return err
	}

	return nil
}

func (m *Manager) deleteUser(id string) error {
	if !m.exec {
		return nil
	}

	if m.userExists(id) {
		if err := m.mgnt.User.Delete("auth0|" + id); err != nil {
			m.counter.IncrError()
			return err
		}
	}

	return nil
}