package kafkacluster

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_cluster", func(r *config.Resource) {
		r.ShortGroup = "kafkacluster"
	})
}
