package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/go-hclog"

	// "github.com/hashicorp/vault/helper/pluginutil"
	"os"
	// "github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/sdk/plugin"
	circleci "github.com/tylux/go-circleci"
)

// func main() {
// 	apiClientMeta := &pluginutil.APIClientMeta{}
// 	flags := apiClientMeta.FlagSet()
// 	flags.Parse(os.Args[1:])

// 	tlsConfig := apiClientMeta.GetTLSConfig()
// 	tlsProviderFunc := pluginutil.VaultPluginTLSProvider(tlsConfig)

// 	if err := plugin.Serve(&plugin.ServeOpts{
// 		BackendFactoryFunc: Factory,
// 		TLSProviderFunc:    tlsProviderFunc,
// 	}); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}

var _ logical.Factory = Factory

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

// Client is the interface for clients used to talk to the CircleCI API.
type Client interface {
	GetBuild(project string, buildNum int) (*circleci.Build, error)
	SetBaseURL(baseURL *url.URL)
}
