module github.com/jeffwecan/vault-circleci-auth-plugin

go 1.15

require (
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/vault/api v1.0.4
	github.com/hashicorp/vault/sdk v0.1.13
	github.com/marcboudreau/vault-circleci-auth-plugin v0.0.0-20180723043507-12d005998064
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.6.1
	github.com/tylux/go-circleci v0.0.0-20180427201651-57548b38ee3b
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
)

replace gopkg.in/ory-am/dockertest.v3 => github.com/ory/dockertest/v3 v3.6.2

replace github.com/marcboudreau/vault-circleci-auth-plugin => ./
