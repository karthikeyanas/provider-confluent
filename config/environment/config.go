package environment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_environment", func(r *config.Resource) {
		r.ShortGroup = "environment"
	})
}
