module github.com/jeffwecan/vault-circleci-auth-plugin

go 1.15

require (
	github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2
	github.com/armon/go-radix v0.0.0-20170727155443-1fca145dffbc
	github.com/davecgh/go-spew v1.1.0
	github.com/fatih/structs v1.0.0
	github.com/golang/protobuf v1.0.0
	github.com/golang/snappy v0.0.0-20170215233205-553a64147049
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce
	github.com/hashicorp/go-cleanhttp v0.0.0-20171218145408-d5fe4b57a186
	github.com/hashicorp/go-hclog v0.0.0-20180122232401-5bcb0f17e364
	github.com/hashicorp/go-multierror v0.0.0-20171204182908-b7773ae21874
	github.com/hashicorp/go-plugin v0.0.0-20180314222826-8068b0bdcfb7
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90
	github.com/hashicorp/go-uuid v0.0.0-20180228145832-27454136f036
	github.com/hashicorp/go-version v0.0.0-20180322230233-23480c066577
	github.com/hashicorp/golang-lru v0.0.0-20180201235237-0fb14efe8c47
	github.com/hashicorp/hcl v0.0.0-20180320202055-f40e974e75af
	github.com/hashicorp/vault v0.9.6
	github.com/mattn/go-colorable v0.0.9
	github.com/mattn/go-isatty v0.0.3
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/mgutz/logxi v0.0.0-20161027140823-aebf8a7d67ab
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747
	github.com/mitchellh/go-testing-interface v0.0.0-20171004221916-a61a99592b77
	github.com/mitchellh/mapstructure v0.0.0-20180220230111-00c29f56e238
	github.com/oklog/run v1.0.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pmezard/go-difflib v1.0.0
	github.com/ryanuber/go-glob v0.0.0-20160226084822-572520ed46db
	github.com/sethgrid/pester v0.0.0-20180227223404-ed9870dad317
	github.com/stretchr/testify v1.2.1
	github.com/tylux/go-circleci v0.0.0-20171109182250-498a7a967f7d
	golang.org/x/net v0.0.0-20180320002117-6078986fec03
	golang.org/x/sys v0.0.0-20180326154331-13d03a9a82fb
	golang.org/x/text v0.3.0
	google.golang.org/genproto v0.0.0-20180323190852-ab0870e398d5
	google.golang.org/grpc v1.10.0
)

replace gopkg.in/ory-am/dockertest.v3 => github.com/ory/dockertest/v3 v3.6.2
