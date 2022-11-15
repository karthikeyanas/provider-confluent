package kafkatopic

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = "kafkatopic"
	})
}
