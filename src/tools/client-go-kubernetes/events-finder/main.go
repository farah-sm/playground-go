package main

import (
	"context"
	"strings"
	//"encoding/json"
	"flag"
	"fmt"

	//"strings"

	//	"strings"

	// "k8s.io/client-go" // exposes all the interfaces used by the client to interact with the API server read more: https://github.com/kubernetes/client-go
	// "k8s.io/api" // all the kubernetes resources are interacted with by the api library
	// "k8s.io/apiMachinery" // Access to utility methods that help developing an API
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")
	msg := flag.String("msg", "", "-msg='Pulling image'")
	flag.Parse()

	// clientcmd library reads the config file and creates a configuration object 'config'
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// handle error
		fmt.Printf("error %s, not able to build config from flags", err.Error())
	}
	// the clientset object is the remote control that allows the program to interact with the API server, using the configuration parsed to it
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, not ble to create clientset", err.Error())
	}

	ctx := context.Background()

	fmt.Printf("------------------------------\n")
	fmt.Printf("Event in Cluster: \n")
	fmt.Printf("------------------------------\n")

	event, err := clientset.CoreV1().Events("").List(ctx, metav1.ListOptions{})

	if err != nil {
		// handle error
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for i, e := range event.Items {
		i=i+1
		if strings.Contains(e.Message, *msg) {
			fmt.Printf("\n%d)Event: %v.\nTime: %v.\nMessage: %v.\n",i, e.Name, e.EventTime, e.Message)
		}
	}
}

