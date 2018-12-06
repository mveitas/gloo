// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionVirtualServiceFunc func(original, desired *VirtualService) (bool, error)

type VirtualServiceReconciler interface {
	Reconcile(namespace string, desiredResources VirtualServiceList, transition TransitionVirtualServiceFunc, opts clients.ListOpts) error
}

func virtualServicesToResources(list VirtualServiceList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, virtualService := range list {
		resourceList = append(resourceList, virtualService)
	}
	return resourceList
}

func NewVirtualServiceReconciler(client VirtualServiceClient) VirtualServiceReconciler {
	return &virtualServiceReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type virtualServiceReconciler struct {
	base reconcile.Reconciler
}

func (r *virtualServiceReconciler) Reconcile(namespace string, desiredResources VirtualServiceList, transition TransitionVirtualServiceFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "virtualService_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*VirtualService), desired.(*VirtualService))
		}
	}
	return r.base.Reconcile(namespace, virtualServicesToResources(desiredResources), transitionResources, opts)
}
