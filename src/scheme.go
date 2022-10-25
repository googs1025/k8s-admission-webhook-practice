package src

import (
	admissionV1 "k8s.io/api/admission/v1"
	admissionV1beta1 "k8s.io/api/admission/v1beta1"
	admissionRegistrationV1 "k8s.io/api/admissionregistration/v1"
	admissionRegistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilRuntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(scheme)

func init() {
	addToScheme(scheme)
}

func addToScheme(scheme *runtime.Scheme) {
	utilRuntime.Must(coreV1.AddToScheme(scheme))
	utilRuntime.Must(admissionV1beta1.AddToScheme(scheme))
	utilRuntime.Must(admissionRegistrationV1beta1.AddToScheme(scheme))
	utilRuntime.Must(admissionV1.AddToScheme(scheme))
	utilRuntime.Must(admissionRegistrationV1.AddToScheme(scheme))
}
