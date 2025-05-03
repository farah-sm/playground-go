package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)
// Purpose
// Services with no backing pods (silent failure).
// Endpoints pointing to not-ready pods (stale DNS).

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s, not able to build config from flags", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, not ble to create clientset", err.Error())
	}

	ctx := context.Background()

	// --------------------------------------------------- PODS -----------------------------------------------------------------------------------

	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}
	mef := make(map[string]string)
	for _, pod := range pods.Items {
		var b string
		for a, r := range pod.Labels {

			 b = fmt.Sprintf("%s:%s", a, r)
			 mef[pod.Name] += b	+ " "
			}
	}

	// ---------------------------------------------------SERVICES -----------------------------------------------------------------------------------
	services, err := clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s, listing services", err.Error())
	}
	var i string
	for _, s := range services.Items {
		for a, r := range s.Spec.Selector {
			 i = fmt.Sprintf("%s:%s", a,r)
			for n, kv := range mef {
				if strings.Contains(kv, i) {
					// count++
					fmt.Println("----------------------------------")
					fmt.Printf("Service: %v. \nPod: %v\nMatched on: %v\n", s.Name, n, kv)
					fmt.Println("----------------------------------")
				}
			}
		}
	}
	fmt.Println("----------------------------------")
}

