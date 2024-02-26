package customresource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var Gateway = schema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "gateways"}
var VirtualService = schema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "virtualservices"}
