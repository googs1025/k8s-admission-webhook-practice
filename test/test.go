package main

import (
	"encoding/json"
	"io/ioutil"
	v1 "k8s.io/api/admission/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"k8s_AdmissionControllers_practice/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		var body []byte
		if r.Body != nil {
			if data, err := ioutil.ReadAll(r.Body); err == nil {
				body = data
			}
		}
		//第二步
		reqAdmissionReview := v1.AdmissionReview{} //请求
		rspAdmissionReview := v1.AdmissionReview{ //响应 ---只构建了一部分
			TypeMeta: metaV1.TypeMeta{
				Kind:       "AdmissionReview",
				APIVersion: "admission.k8s.io/v1",
			},
		}
		//第三步。 把body decode 成对象
		deserializer := src.Codecs.UniversalDeserializer()
		if _, _, err := deserializer.Decode(body, nil, &reqAdmissionReview); err != nil {
			klog.Error(err)
			rspAdmissionReview.Response = src.ToV1AdmissionResponse(err)
		} else {
			rspAdmissionReview.Response = src.AdmitPods(reqAdmissionReview) //我们的业务
		}
		rspAdmissionReview.Response.UID = reqAdmissionReview.Request.UID
		respBytes,_ := json.Marshal(rspAdmissionReview)

		w.Write(respBytes)
	})

	http.ListenAndServe(":8080",nil )
}
