/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	environment "github.com/karthikeyanas/provider-confluent/internal/controller/environment/environment"
	cluster "github.com/karthikeyanas/provider-confluent/internal/controller/kafkacluster/cluster"
	topic "github.com/karthikeyanas/provider-confluent/internal/controller/kafkatopic/topic"
	providerconfig "github.com/karthikeyanas/provider-confluent/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.Setup,
		cluster.Setup,
		topic.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
