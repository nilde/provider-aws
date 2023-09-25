package deliverystream

import (
	"context"
	//"strings"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/firehose"
	//xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	//"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/firehose/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/apis/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/pkg/features"
)

// SetupStage adds a controller that reconciles Stage.
func SetupDeliveryStream(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.DeliveryStreamGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), v1alpha1.StoreConfigGroupVersionKind))
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(resource.DesiredStateChanged()).
		For(&svcapitypes.DeliveryStream{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.DeliveryStreamGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithPollInterval(o.PollInterval),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
			managed.WithConnectionPublishers(cps...)))
}
func preObserve(_ context.Context, cr *svcapitypes.DeliveryStream, obj *svcsdk.DescribeDeliveryStreamInput) error {
	obj.DeliveryStreamName = aws.String(meta.GetExternalName(cr))
	return nil
}

func preCreate(_ context.Context, cr *svcapitypes.DeliveryStream, obj *svcsdk.CreateDeliveryStreamInput) error {
	obj.DeliveryStreamName = aws.String(meta.GetExternalName(cr))
	return nil
}

/**
func preUpdate(_ context.Context, cr *svcapitypes.DeliveryStream, obj *svcsdk.UpdateDeliveryStreamInput) error {
	obj.DeliveryStreamName = aws.String(meta.GetExternalName(cr))
	return nil
}
**/

func preDelete(_ context.Context, cr *svcapitypes.DeliveryStream, obj *svcsdk.DeleteDeliveryStreamInput) error {
	obj.DeliveryStreamName = aws.String(meta.GetExternalName(cr))
	return nil
}