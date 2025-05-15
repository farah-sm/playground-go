package main

import (
	"context"
	"flag"
	"fmt"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"os"
	"strings"
	"k8s.io/client-go/tools/clientcmd"
)

// This is a cool way to list the custom resources. 
// However, it's a pretty monotonous way to do so, you need to provide EVERY
// SINGLE GVR of each CRD. Not ideal.

// To introduce "Discovery" to print the stout the GVR of every CRD and consume via the dynamicClient.

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")
	flag.Parse()

	// clientcmd library reads the config file and creates a configuration object 'config'
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s, not able to build config from flags", err.Error())
	}

	ctx :=context.Background()

	// the clientset object is the control that allows the program to interact with the API server, using the configuration parsed to it
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		fmt.Printf("error %s, not ble to create clientset", err.Error())
	}

	// Create an API extensions client, for getting the resource name
	apiextensionsClient, err := apiextensionsv1.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating API extensions client: %v\n", err)
		os.Exit(1)
	}

	// An effort to pull in the Groups and Versions
	_, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		fmt.Printf("error %s, not Able to get resource via dynamic client", err.Error())
	}

	// In an effort to pull in the resource names
	crds, err := apiextensionsClient.CustomResourceDefinitions().List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing CRD: %v\n", err)
		os.Exit(1)
	}

	// map to store the groups and versions of each CRD
	gv := make(map[string]string)

	// to store the crd resource names
	crdz := []string{}

	// looping through each CRD resource name and appending it into the slice
	for _, crd := range crds.Items {
		plural := crd.Spec.Names.Plural
		crdz = append(crdz, plural)
		fmt.Printf("%v, ", crd.Spec.Names.Plural)

	}

	for _, b := range apiResourceLists {
		version := make([]string, 2)
		if strings.Contains(b.GroupVersion, "/") {
			version = strings.Split(b.GroupVersion, "/")

			// fmt.Printf("G:%v V:%v\n", version[0], version[1])
			a := version[0]
			b := version[1]
			gv[a] += b

			// a := version[0]
			// b := version[1]
			// gv[a] += b
		}
		fmt.Printf("\n%v\n", gv)

		i := 0
		for g, v := range gv {

			//if strings.Contains(g, crdz[i]) {

			fmt.Printf("G: %v, V: %v, R: %v", g, v, crdz[i])
			//}

		}

	}

}
