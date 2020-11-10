package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	circleci "github.com/tylux/go-circleci"
	cache "github.com/patrickmn/go-cache"
)

type backend struct {
	*framework.Backend

	client     Client
	ProjectMap *framework.PolicyMap

	AttemptsCache *cache.Cache
	CacheExpiry   time.Duration
}

// Client is the interface for clients used to talk to the CircleCI API.
type Client interface {
	GetBuild(project string, buildNum int) (*circleci.Build, error)
	SetBaseURL(baseURL *url.URL)
}

// Factory constructs the plugin instance with the provided BackendConfig.
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend() (*backend, error) {

	var b backend

	b.ProjectMap = &framework.PolicyMap{
		PathMap: framework.PathMap{
			Name: "projects",
		},
		DefaultKey: "default",
	}

	b.AttemptsCache = cache.New(5*time.Hour, cache.NoExpiration)
	b.CacheExpiry = 5 * time.Hour

	// allPaths := append(b.ProjectMap.Paths(), pathConfig(&b), pathLogin(&b))

	b.Backend = &framework.Backend{
		// Help:        strings.TrimSpace(mockHelp),
		BackendType: logical.TypeCredential,
		// AuthRenew:   b.pathAuthRenew,
		PeriodicFunc: b.periodicFunc,
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{
				"login",
			},
		},
		// Paths:       allPaths,
		Paths: framework.PathAppend(
			b.ProjectMap.Paths(),
			[]*framework.Path{
				b.pathConfig(),
				b.pathLogin(),
			},
		),
	}

	return &b, nil
}

func (b *backend) GetClient(token, vcsType, owner string) Client {
	if b.client == nil {
		b.client = NewCCIClient(token, vcsType, owner)
	}

	return b.client
}

func (b *backend) periodicFunc(_ context.Context, _ *logical.Request) error {
	b.Logger().Trace("periodicFunc called")
	b.AttemptsCache.DeleteExpired()

	return nil
}

// This method takes in the TTL and MaxTTL values provided by the user,
// compares those with the SystemView values. If they are empty a value of 0 is
// set, which will cause initial secret or LeaseExtend operations to use the
// mount/system defaults.  If they are set, their boundaries are validated.
func (b *backend) SanitizeTTLStr(ttlStr, maxTTLStr string) (ttl, maxTTL time.Duration, err error) {
	if len(ttlStr) == 0 || ttlStr == "0" {
		ttl = 0
	} else {
		ttl, err = time.ParseDuration(ttlStr)
		if err != nil {
			return 0, 0, fmt.Errorf("Invalid ttl: %s", err)
		}
	}

	if len(maxTTLStr) == 0 || maxTTLStr == "0" {
		maxTTL = 0
	} else {
		maxTTL, err = time.ParseDuration(maxTTLStr)
		if err != nil {
			return 0, 0, fmt.Errorf("Invalid max_ttl: %s", err)
		}
	}

	ttl, maxTTL, err = b.SanitizeTTL(ttl, maxTTL)

	return
}

// Caps the boundaries of ttl and max_ttl values to the backend mount's max_ttl value.
func (b *backend) SanitizeTTL(ttl, maxTTL time.Duration) (time.Duration, time.Duration, error) {
	sysMaxTTL := b.System().MaxLeaseTTL()
	if ttl > sysMaxTTL {
		return 0, 0, fmt.Errorf("\"ttl\" value must be less than allowed max lease TTL value '%s'", sysMaxTTL.String())
	}
	if maxTTL > sysMaxTTL {
		return 0, 0, fmt.Errorf("\"max_ttl\" value must be less than allowed max lease TTL value '%s'", sysMaxTTL.String())
	}
	if ttl > maxTTL && maxTTL != 0 {
		ttl = maxTTL
	}
	return ttl, maxTTL, nil
}
