package controllercontext

import (
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"

	"github.com/giantswarm/aws-operator/client/aws"
)

type ContextClient struct {
	ControlPlane  ContextClientControlPlane
	TenantCluster ContextClientTenantCluster
}

type ContextClientControlPlane struct {
	AWS aws.Clients
}

type ContextClientTenantCluster struct {
	AWS aws.Clients
	K8s k8sclient.Interface
}
