# aserto-idp-seed

Aserto IDP seeding utility, a simple tool for populating users your IDP, the tool only supports Auth0 domains. 

## Installation

To install the seeder utility, there are a couple of options, outlined in this section. The seeder is a simple (golang) single static binary, so it can be placed according to your own liking and removed by simply deleting the `aserto-idp-seed(.exe)` binary to uninstall

### Binary release
You can download the latest binary [release](https://github.com/aserto-dev/aserto-idp-seed/releases) for Windows 10, MacOS and Linux from the projects GitHub release page.

### Homebrew installation

On MacOS and Linux you can use the Aserto Homebrew tap to install our tools, register it using:

	brew tap aserto-dev/tap

To install, execute:

	brew install aserto-idp-seed

or

	brew install aserto-dev/tap/aserto-idp-seed

To update execute:

	brew upgrade aserto-idp-seed

To uninstall execute:

	brew uninstall aserto-idp-seed


### Source based installation

To install from source, you need [golang](https://golang.org/dl/) version **1.16.x or higher**, as the tool depends on the **1.16** introduced embedded file support.

Check if the correct version of golang is installed using:

	go version

Which should provide the conformation you are using version 1.16 or higher

	❯ go version
	go version go1.16.3 darwin/amd64

To install execute the release:

	go install github.com/aserto-dev/aserto-idp-seed/cmd/aserto-idp-seed@v0.0.12

Which will output:
	
	❯ go install github.com/aserto-dev/aserto-idp-seed/cmd/aserto-idp-seed@v0.0.12
	go: downloading github.com/aserto-dev/aserto-idp-seed v0.0.11

`go install` will install the binary in the [GOBIN](https://golang.org/ref/mod#go-install) directory.


## Configuration

The seeder utility uses a `.env` file and/or environment variables for its configuration. 

> **NOTE:** The .env file must resides in the current working directory when executing the seeder. 

A template [.env](https://github.com/aserto-dev/aserto-idp-seed/blob/aea63bb632821351fea419b6a820112cf22f33c5/.env.template) file is available in the root of the source code repository. Copy the template file and rename to `.env`, next adjust the values accordingly using your favorite text editor.

### .env file 

	AUTH0_DOMAIN="mydomain.us.auth0.com"
	AUTH0_CLIENT_ID="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	AUTH0_CLIENT_SECRET="yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"

	TEMPL_CORPORATION="acmecorp"
	TEMPL_EMAIL_DOMAIN="acmecorp.com"
	TEMPL_PASSWORD="V@rySecr#et321!"

### Environment variables

| Variable names        | Descriptions                                       |
| ----------------------| -------------------------------------------------- |
| AUTH0\_DOMAIN         | Auth0 domain value, like "mydomain.us.auth0.com"   |
| AUTH0\_CLIENT\_ID.    | Auth0 client ID of management API connection       |
| AUTH0\_CLIENT\_SECRET | Auth0 client secret of management API connection   |
|                       |                                                    |
| TEMPL\_CORPORATION    | corporation name, used for company role membership |
| TEMPL\_EMAIL\_DOMAIN  | email domain, as user@<$TEMPL_CORPORATION>         |
| TEMPL\_PASSWORD       | password value used for all seeded users           |

For more information about [Auth0 management API settings](https://auth0.com/docs/tokens/management-api-access-tokens) see link.


## Check & Verify

To check if the IDP settings are working correctly and to validate which user accounts are available in the IDP's directory, the seeder tools provides the `users` subcommand.

When execute without any parameters, it will connect to the IDP and enumerate the users in to a simple count.

	❯ aserto-idp-seed users
	row count: 272 skip count 0 error count: 0

If you want to inspect the content currently present in the IDP's directory, execute:

	aserto-idp-seed users --nocount --spew

This will output each user record as a json payload. 

This might be too much information that can easily be reduced using standard `stdout` filtering using using tools like [jq](https://stedolan.github.io/jq/). For example the following command will return a list of just the email names:

	aserto-idp-seed users --nocount --spew | jq .email

Now that we know what is in the IDP's directory, lets seed it with test users.

## Seed

For ease of use and consumption of the tool, the [data file](https://raw.githubusercontent.com/aserto-dev/aserto-idp-seed/main/pkg/data/users.json) containing the seed users is embedded inside the tool. 

Some of the data is templatized, the template values are substituted with the `TEMPL_*` environment variables described above. 

To inspect the data used before sending it to the IDP, you can use the `--dryrun` option.

To get a simple count of the number of user records in the seed data set, execute: 

	❯ aserto-idp-seed seed --dryrun
	row count: 272 skip count 0 error count: 0

To inspect the payloads, execute:

	aserto-idp-seed seed --dryrun --spew

To filter the payload, execute:

	aserto-idp-seed seed --dryrun --nocount --spew | jq .email

To load the seed data, execute:

	aserto-idp-seed seed

## Reset

To remove the seeded user records from the IDP's directory, use the `reset` subcommand

To remove the seeded user records, execute:

	aserto-idp-seed reset

## Example user record

```
{
  "user_id": "67b42b6c-6bd8-40e2-a622-fe69eacd3d47",
  "connection": "Username-Password-Authentication",
  "email": "chrisjohns@acmecorp.com",
  "given_name": "Chris",
  "family_name": "Johnson",
  "nickname": "Chris Johnson [SALES]",
  "password": "**********",
  "user_metadata": {
    "department": "Sales Engagement Management",
    "dn": "cn=chris johnson [sales]",
    "manager": "2bfaa552-d9a5-41e9-a6c3-5be62b4433c8",
    "phone": "+1-206-555-9004",
    "title": "Salesperson",
    "username": "chrisjohns"
  },
  "app_metadata": {
    "roles": [
      "user",
      "acmecorp",
      "sales-engagement-management"
    ]
  },
  "picture": "https://github.com/aserto-demo/contoso-ad-sample/raw/main/UserImages/Chris%20Johnson%20%5BSALES%5D.jpg",
  "email_verified": true
}
```
