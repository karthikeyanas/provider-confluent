/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/karthikeyanas/provider-confluent/config/environment"
	"github.com/karthikeyanas/provider-confluent/config/kafkacluster"
	"github.com/karthikeyanas/provider-confluent/config/kafkatopic"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "confluent"
	modulePath     = "github.com/karthikeyanas/provider-confluent"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		kafkacluster.Configure,
		kafkatopic.Configure,
		environment.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
