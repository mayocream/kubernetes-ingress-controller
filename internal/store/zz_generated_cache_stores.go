// Code generated by hack/generators/cache-stores/main.go; DO NOT EDIT.
// If you want to add a new type to the cache store, you need to add a new entry to the supportedTypes list in spec.go.
package store

import (
	"fmt"
	"sync"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kongv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
	kongv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	kongv1beta1 "github.com/kong/kubernetes-configuration/api/configuration/v1beta1"
	incubatorv1alpha1 "github.com/kong/kubernetes-configuration/api/incubator/v1alpha1"
	"github.com/kong/kubernetes-ingress-controller/v3/internal/gatewayapi"
)

// CacheStores stores cache.Store for all Kinds of k8s objects that
// the Ingress Controller reads.
type CacheStores struct {
	IngressV1                      cache.Store
	IngressClassV1                 cache.Store
	Service                        cache.Store
	Secret                         cache.Store
	EndpointSlice                  cache.Store
	HTTPRoute                      cache.Store
	UDPRoute                       cache.Store
	TCPRoute                       cache.Store
	TLSRoute                       cache.Store
	GRPCRoute                      cache.Store
	ReferenceGrant                 cache.Store
	Gateway                        cache.Store
	Plugin                         cache.Store
	ClusterPlugin                  cache.Store
	Consumer                       cache.Store
	ConsumerGroup                  cache.Store
	KongIngress                    cache.Store
	TCPIngress                     cache.Store
	UDPIngress                     cache.Store
	KongUpstreamPolicy             cache.Store
	IngressClassParametersV1alpha1 cache.Store
	KongServiceFacade              cache.Store
	KongVault                      cache.Store
	KongCustomEntity               cache.Store

	l *sync.RWMutex
}

// NewCacheStores is a convenience function for CacheStores to initialize all attributes with new cache stores.
func NewCacheStores() CacheStores {
	return CacheStores{
		IngressV1:                      cache.NewStore(namespacedKeyFunc),
		IngressClassV1:                 cache.NewStore(clusterWideKeyFunc),
		Service:                        cache.NewStore(namespacedKeyFunc),
		Secret:                         cache.NewStore(namespacedKeyFunc),
		EndpointSlice:                  cache.NewStore(namespacedKeyFunc),
		HTTPRoute:                      cache.NewStore(namespacedKeyFunc),
		UDPRoute:                       cache.NewStore(namespacedKeyFunc),
		TCPRoute:                       cache.NewStore(namespacedKeyFunc),
		TLSRoute:                       cache.NewStore(namespacedKeyFunc),
		GRPCRoute:                      cache.NewStore(namespacedKeyFunc),
		ReferenceGrant:                 cache.NewStore(namespacedKeyFunc),
		Gateway:                        cache.NewStore(namespacedKeyFunc),
		Plugin:                         cache.NewStore(namespacedKeyFunc),
		ClusterPlugin:                  cache.NewStore(clusterWideKeyFunc),
		Consumer:                       cache.NewStore(namespacedKeyFunc),
		ConsumerGroup:                  cache.NewStore(namespacedKeyFunc),
		KongIngress:                    cache.NewStore(namespacedKeyFunc),
		TCPIngress:                     cache.NewStore(namespacedKeyFunc),
		UDPIngress:                     cache.NewStore(namespacedKeyFunc),
		KongUpstreamPolicy:             cache.NewStore(namespacedKeyFunc),
		IngressClassParametersV1alpha1: cache.NewStore(namespacedKeyFunc),
		KongServiceFacade:              cache.NewStore(namespacedKeyFunc),
		KongVault:                      cache.NewStore(clusterWideKeyFunc),
		KongCustomEntity:               cache.NewStore(namespacedKeyFunc),

		l: &sync.RWMutex{},
	}
}

// Get checks whether or not there's already some version of the provided object present in the cache.
func (c CacheStores) Get(obj runtime.Object) (item interface{}, exists bool, err error) {
	c.l.RLock()
	defer c.l.RUnlock()

	switch obj := obj.(type) {
	case *netv1.Ingress:
		return c.IngressV1.Get(obj)
	case *netv1.IngressClass:
		return c.IngressClassV1.Get(obj)
	case *corev1.Service:
		return c.Service.Get(obj)
	case *corev1.Secret:
		return c.Secret.Get(obj)
	case *discoveryv1.EndpointSlice:
		return c.EndpointSlice.Get(obj)
	case *gatewayapi.HTTPRoute:
		return c.HTTPRoute.Get(obj)
	case *gatewayapi.UDPRoute:
		return c.UDPRoute.Get(obj)
	case *gatewayapi.TCPRoute:
		return c.TCPRoute.Get(obj)
	case *gatewayapi.TLSRoute:
		return c.TLSRoute.Get(obj)
	case *gatewayapi.GRPCRoute:
		return c.GRPCRoute.Get(obj)
	case *gatewayapi.ReferenceGrant:
		return c.ReferenceGrant.Get(obj)
	case *gatewayapi.Gateway:
		return c.Gateway.Get(obj)
	case *kongv1.KongPlugin:
		return c.Plugin.Get(obj)
	case *kongv1.KongClusterPlugin:
		return c.ClusterPlugin.Get(obj)
	case *kongv1.KongConsumer:
		return c.Consumer.Get(obj)
	case *kongv1beta1.KongConsumerGroup:
		return c.ConsumerGroup.Get(obj)
	case *kongv1.KongIngress:
		return c.KongIngress.Get(obj)
	case *kongv1beta1.TCPIngress:
		return c.TCPIngress.Get(obj)
	case *kongv1beta1.UDPIngress:
		return c.UDPIngress.Get(obj)
	case *kongv1beta1.KongUpstreamPolicy:
		return c.KongUpstreamPolicy.Get(obj)
	case *kongv1alpha1.IngressClassParameters:
		return c.IngressClassParametersV1alpha1.Get(obj)
	case *incubatorv1alpha1.KongServiceFacade:
		return c.KongServiceFacade.Get(obj)
	case *kongv1alpha1.KongVault:
		return c.KongVault.Get(obj)
	case *kongv1alpha1.KongCustomEntity:
		return c.KongCustomEntity.Get(obj)
	}
	return nil, false, fmt.Errorf("%T is not a supported cache object type", obj)
}

// Add stores a provided runtime.Object into the CacheStore if it's of a supported type.
// The CacheStore must be initialized (see NewCacheStores()) or this will panic.
func (c CacheStores) Add(obj runtime.Object) error {
	c.l.Lock()
	defer c.l.Unlock()

	switch obj := obj.(type) {
	case *netv1.Ingress:
		return c.IngressV1.Add(obj)
	case *netv1.IngressClass:
		return c.IngressClassV1.Add(obj)
	case *corev1.Service:
		return c.Service.Add(obj)
	case *corev1.Secret:
		return c.Secret.Add(obj)
	case *discoveryv1.EndpointSlice:
		return c.EndpointSlice.Add(obj)
	case *gatewayapi.HTTPRoute:
		return c.HTTPRoute.Add(obj)
	case *gatewayapi.UDPRoute:
		return c.UDPRoute.Add(obj)
	case *gatewayapi.TCPRoute:
		return c.TCPRoute.Add(obj)
	case *gatewayapi.TLSRoute:
		return c.TLSRoute.Add(obj)
	case *gatewayapi.GRPCRoute:
		return c.GRPCRoute.Add(obj)
	case *gatewayapi.ReferenceGrant:
		return c.ReferenceGrant.Add(obj)
	case *gatewayapi.Gateway:
		return c.Gateway.Add(obj)
	case *kongv1.KongPlugin:
		return c.Plugin.Add(obj)
	case *kongv1.KongClusterPlugin:
		return c.ClusterPlugin.Add(obj)
	case *kongv1.KongConsumer:
		return c.Consumer.Add(obj)
	case *kongv1beta1.KongConsumerGroup:
		return c.ConsumerGroup.Add(obj)
	case *kongv1.KongIngress:
		return c.KongIngress.Add(obj)
	case *kongv1beta1.TCPIngress:
		return c.TCPIngress.Add(obj)
	case *kongv1beta1.UDPIngress:
		return c.UDPIngress.Add(obj)
	case *kongv1beta1.KongUpstreamPolicy:
		return c.KongUpstreamPolicy.Add(obj)
	case *kongv1alpha1.IngressClassParameters:
		return c.IngressClassParametersV1alpha1.Add(obj)
	case *incubatorv1alpha1.KongServiceFacade:
		return c.KongServiceFacade.Add(obj)
	case *kongv1alpha1.KongVault:
		return c.KongVault.Add(obj)
	case *kongv1alpha1.KongCustomEntity:
		return c.KongCustomEntity.Add(obj)
	}
	return fmt.Errorf("cannot add unsupported kind %q to the store", obj.GetObjectKind().GroupVersionKind())
}

// Delete removes a provided runtime.Object from the CacheStore if it's of a supported type.
// The CacheStore must be initialized (see NewCacheStores()) or this will panic.
func (c CacheStores) Delete(obj runtime.Object) error {
	c.l.Lock()
	defer c.l.Unlock()

	switch obj := obj.(type) {
	case *netv1.Ingress:
		return c.IngressV1.Delete(obj)
	case *netv1.IngressClass:
		return c.IngressClassV1.Delete(obj)
	case *corev1.Service:
		return c.Service.Delete(obj)
	case *corev1.Secret:
		return c.Secret.Delete(obj)
	case *discoveryv1.EndpointSlice:
		return c.EndpointSlice.Delete(obj)
	case *gatewayapi.HTTPRoute:
		return c.HTTPRoute.Delete(obj)
	case *gatewayapi.UDPRoute:
		return c.UDPRoute.Delete(obj)
	case *gatewayapi.TCPRoute:
		return c.TCPRoute.Delete(obj)
	case *gatewayapi.TLSRoute:
		return c.TLSRoute.Delete(obj)
	case *gatewayapi.GRPCRoute:
		return c.GRPCRoute.Delete(obj)
	case *gatewayapi.ReferenceGrant:
		return c.ReferenceGrant.Delete(obj)
	case *gatewayapi.Gateway:
		return c.Gateway.Delete(obj)
	case *kongv1.KongPlugin:
		return c.Plugin.Delete(obj)
	case *kongv1.KongClusterPlugin:
		return c.ClusterPlugin.Delete(obj)
	case *kongv1.KongConsumer:
		return c.Consumer.Delete(obj)
	case *kongv1beta1.KongConsumerGroup:
		return c.ConsumerGroup.Delete(obj)
	case *kongv1.KongIngress:
		return c.KongIngress.Delete(obj)
	case *kongv1beta1.TCPIngress:
		return c.TCPIngress.Delete(obj)
	case *kongv1beta1.UDPIngress:
		return c.UDPIngress.Delete(obj)
	case *kongv1beta1.KongUpstreamPolicy:
		return c.KongUpstreamPolicy.Delete(obj)
	case *kongv1alpha1.IngressClassParameters:
		return c.IngressClassParametersV1alpha1.Delete(obj)
	case *incubatorv1alpha1.KongServiceFacade:
		return c.KongServiceFacade.Delete(obj)
	case *kongv1alpha1.KongVault:
		return c.KongVault.Delete(obj)
	case *kongv1alpha1.KongCustomEntity:
		return c.KongCustomEntity.Delete(obj)
	}
	return fmt.Errorf("cannot delete unsupported kind %q from the store", obj.GetObjectKind().GroupVersionKind())
}

// ListAllStores returns a list of all cache stores embedded in the struct.
func (c CacheStores) ListAllStores() []cache.Store {
	return []cache.Store{
		c.IngressV1,
		c.IngressClassV1,
		c.Service,
		c.Secret,
		c.EndpointSlice,
		c.HTTPRoute,
		c.UDPRoute,
		c.TCPRoute,
		c.TLSRoute,
		c.GRPCRoute,
		c.ReferenceGrant,
		c.Gateway,
		c.Plugin,
		c.ClusterPlugin,
		c.Consumer,
		c.ConsumerGroup,
		c.KongIngress,
		c.TCPIngress,
		c.UDPIngress,
		c.KongUpstreamPolicy,
		c.IngressClassParametersV1alpha1,
		c.KongServiceFacade,
		c.KongVault,
		c.KongCustomEntity,
	}
}

// SupportedTypes returns a list of supported types for the cache.
func (c CacheStores) SupportedTypes() []client.Object {
	return []client.Object{
		&netv1.Ingress{},
		&netv1.IngressClass{},
		&corev1.Service{},
		&corev1.Secret{},
		&discoveryv1.EndpointSlice{},
		&gatewayapi.HTTPRoute{},
		&gatewayapi.UDPRoute{},
		&gatewayapi.TCPRoute{},
		&gatewayapi.TLSRoute{},
		&gatewayapi.GRPCRoute{},
		&gatewayapi.ReferenceGrant{},
		&gatewayapi.Gateway{},
		&kongv1.KongPlugin{},
		&kongv1.KongClusterPlugin{},
		&kongv1.KongConsumer{},
		&kongv1beta1.KongConsumerGroup{},
		&kongv1.KongIngress{},
		&kongv1beta1.TCPIngress{},
		&kongv1beta1.UDPIngress{},
		&kongv1beta1.KongUpstreamPolicy{},
		&kongv1alpha1.IngressClassParameters{},
		&incubatorv1alpha1.KongServiceFacade{},
		&kongv1alpha1.KongVault{},
		&kongv1alpha1.KongCustomEntity{},
	}
}
