package firehose

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/controller"

	"github.com/crossplane-contrib/provider-aws/pkg/controller/firehose/deliverystream"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/setup"
)

// Setup athena controllers.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	return setup.SetupControllers(
		mgr, o,
		deliverystream.SetupDeliveryStream,

	)
}