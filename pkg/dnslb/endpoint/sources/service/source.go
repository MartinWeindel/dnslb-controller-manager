package service

import (
	"fmt"
	"github.com/gardener/dnslb-controller-manager/pkg/dnslb/endpoint/sources"
	"github.com/gardener/lib/pkg/logger"
	"github.com/gardener/lib/pkg/resources"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Source struct {
	*resources.ServiceObject
}

type SourceType struct {
	schema.GroupKind
}

var _ sources.Source = &Source{}

func init() {
	sources.Register(&SourceType{resources.NewGroupKind(api.GroupName,"Service")})
}

func (this *SourceType) GetGroupKind() schema.GroupKind{
	return this.GroupKind
}

func (this *SourceType) Get(obj resources.Object) (sources.Source, error) {
	if obj.GroupKind()!=this.GroupKind {
		return nil,fmt.Errorf("invalid object type %q", obj.GroupKind())
	}
	return &Source{resources.Service(obj)}, nil
}


func (this *Source) GetTargets(logger logger.LogContext, lb resources.Object) (ip, cname string) {
	status:=this.Status()
	for _, i := range status.LoadBalancer.Ingress {
		if i.IP != "" {
			ip=i.IP
		}
		if i.Hostname != "" {
			cname=i.Hostname
		}
	}
	return "",""
}

func (this *Source) Validate(lb resources.Object) (bool, error) {
	ok, err := HasLoadBalancer(this.Service())
	if err != nil {
		return false, err
	}
	if !ok {
		return true, fmt.Errorf("load balancer not yet assigned for '%s'", this.ObjectName())
	}
	return true, nil
}


func HasLoadBalancer(svc *api.Service) (bool, error) {
	if svc.Spec.Type != "LoadBalancer" {
		return false, fmt.Errorf("service %s/%s is not of type LoadBalancer",
			svc.Namespace, svc.Name)
	}
	for _, i := range svc.Status.LoadBalancer.Ingress {
		if i.IP != "" || i.Hostname != "" {
			return true, nil
		}
	}
	return false, nil
}