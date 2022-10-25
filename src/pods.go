package src

import (
	"fmt"
	"k8s.io/api/admission/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func AdmitPods(ar v1.AdmissionReview) *v1.AdmissionResponse {
	// 查看传递来的对象是否为pod
	podResource := metaV1.GroupVersionResource{
		Group: "",
		Version: "v1",
		Resource: "pods",
	}

	// err
	if ar.Request.Resource != podResource {
		err := fmt.Errorf("expect resource to be %s", podResource)
		klog.Error(err)
		return ToV1AdmissionResponse(err)
	}

	// 解析client端传进来的json
	raw := ar.Request.Object.Raw
	pod := coreV1.Pod{}
	deserializer := Codecs.UniversalDeserializer()	// 解析json格式文件
	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
		klog.Error(err)
		return ToV1AdmissionResponse(err)
	}
	//fmt.Println(pod)

	reviewResponse := v1.AdmissionResponse{}

	if pod.Name=="shenyi"{
		reviewResponse.Allowed = false
		reviewResponse.Result = &metaV1.Status{Code:503,Message: "pod name cannot be shenyi"}
	}else{
		reviewResponse.Allowed = true
	}


	return &reviewResponse
}
