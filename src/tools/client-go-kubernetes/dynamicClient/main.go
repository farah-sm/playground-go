package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)


// This is a cool way to list the custom resources. However, it's a pretty monotonous way to do so, you need to provide EVERY
// SINGLE GVR of each CRD. Not ideal. 

// To intorudce "Discovery" to print the stout the GVR of every CRD and consume via the dynamicClient. 

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/config", "location of your kubeconfig file")
//	msg := flag.String("msg", "", "-msg='Pulling image'")
	flag.Parse()

	// clientcmd library reads the config file and creates a configuration object 'config'
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// handle error
		fmt.Printf("error %s, not able to build config from flags", err.Error())
	}
	// the clientset object is the remote control that allows the program to interact with the API server, using the configuration parsed to it
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, not ble to create clientset", err.Error())
	}

	ctx := context.Background()

	resources, err := dynClient.Resource(schema.GroupVersionResource{
		Group:	"helm.cattle.io",
		Version: "v1",
		Resource: "helmcharts",
	}).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s, not Able to get resource via dynamic client", err.Error())
	}

	for _, b := range resources.Items {
		fmt.Printf("CRD Name: %v in the Namepsace: %v\n", b.GetName(), b.GetNamespace())

	}




}
