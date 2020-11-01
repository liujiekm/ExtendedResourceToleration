package main

import (

	"fmt"
	"k8s.io/api/admission/v1beta1"
	//corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"log"
	"net/http"
	"path/filepath"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/apis/core"

)

const (
	tlsDir      = `/run/secrets/tls`
	tlsCertFile = `tls.crt`
	tlsKeyFile  = `tls.key`
)

var (
	podResource = metav1.GroupVersionResource{Version: "v1", Resource: "pods"}
)


func applyToleration(req *v1beta1.AdmissionRequest) ([]patchOperation, error) {
	// This handler should only get called on Pod objects as per the MutatingWebhookConfiguration in the YAML file.
	// However, if (for whatever reason) this gets invoked on an object of a different kind, issue a log message but
	// let the object request pass through otherwise.
	if req.Resource != podResource {
		log.Printf("expect resource to be %s", podResource)
		return nil, nil
	}

	// Parse the Pod object.
	raw := req.Object.Raw
	pod := core.Pod{}
	if _, _, err := universalDeserializer.Decode(raw, nil, &pod); err != nil {
		return nil, fmt.Errorf("could not deserialize pod object: %v", err)
	}


	resources := sets.String{}

	for _,container := range pod.Spec.Containers{
		for resourceName:= range container.Resources.Requests{

			log.Printf("find resource name:%s in container:%s -- is extended resource name:%v", resourceName,container.Name,helper.IsExtendedResourceName(resourceName))
			if helper.IsExtendedResourceName(resourceName){
				resources.Insert(string(resourceName))
			}

		}
	}

	for _,container := range pod.Spec.InitContainers{
		for resourceName:= range container.Resources.Requests{
			log.Printf("find resource name:%s in init container:%s -- is extended resource name:%v", resourceName,container.Name,helper.IsExtendedResourceName(resourceName))
			if helper.IsExtendedResourceName(resourceName){
				resources.Insert(string(resourceName))
			}
		}
	}

	patches:= []patchOperation{}
	// Doing .List() so that we get a stable sorted list.
	// This allows us to test adding tolerations for multiple extended resources.
	tolerations:=[]string{}
	for _, resource := range resources.List() {
		helper.AddOrUpdateTolerationInPod(&pod, &core.Toleration{
			Key:      resource,
			Operator: core.TolerationOpExists,
			Effect:   core.TaintEffectNoSchedule,
		})

		tolerations = append(tolerations,"{\"key\":"+resource+",\"operator\":"+string(core.TolerationOpExists)+",\"effect\":"+string(core.TaintEffectNoSchedule)+"}")

	}

	log.Printf("tolerations:%v",tolerations)
	patches = append(patches, patchOperation{
		Op:    "add",
		Path:  "/spec/tolerations",
		Value: "["+strings.Join(tolerations,"")+"]",
	})
	// Retrieve the `runAsNonRoot` and `runAsUser` values.
	//var runAsNonRoot *bool
	//var runAsUser *int64
	//if pod.Spec.SecurityContext != nil {
	//	runAsNonRoot = pod.Spec.SecurityContext.RunAsNonRoot
	//	runAsUser = pod.Spec.SecurityContext.RunAsUser
	//}

	// Create patch operations to apply sensible defaults, if those options are not set explicitly.

	//if runAsNonRoot == nil {
	//	patches = append(patches, patchOperation{
	//		Op:    "add",
	//		Path:  "/spec/securityContext/runAsNonRoot",
	//		// The value must not be true if runAsUser is set to 0, as otherwise we would create a conflicting
	//		// configuration ourselves.
	//		Value: runAsUser == nil || *runAsUser != 0,
	//	})
	//
	//	if runAsUser == nil {
	//		patches = append(patches, patchOperation{
	//			Op:    "add",
	//			Path:  "/spec/securityContext/runAsUser",
	//			Value: 1234,
	//		})
	//	}
	//} else if *runAsNonRoot == true && (runAsUser != nil && *runAsUser == 0) {
	//	// Make sure that the settings are not contradictory, and fail the object creation if they are.
	//	return nil, errors.New("runAsNonRoot specified, but runAsUser set to 0 (the root user)")
	//}

	return patches, nil
}

func main() {
	certPath := filepath.Join(tlsDir, tlsCertFile)
	keyPath := filepath.Join(tlsDir, tlsKeyFile)

	mux := http.NewServeMux()
	mux.Handle("/mutate", admitFuncHandler(applyToleration))
	server := &http.Server{
		// We listen on port 8443 such that we do not need root privileges or extra capabilities for this server.
		// The Service object will take care of mapping this port to the HTTPS port 443.
		Addr:    ":8443",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}




