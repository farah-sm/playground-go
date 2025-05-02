package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

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
	uname := flag.String("ldap", "", "-ldap='username'")
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



	// ------------------------------------------------------------------
	
	rb, err := clientset.RbacV1().RoleBindings("").List(ctx, metav1.ListOptions{})
	if err != nil {
		// handle error
		fmt.Printf("error %s, listing roles from namespace", err.Error())
	}
	roler := make(map[string]int64)
	
	jsonData, err := json.MarshalIndent(rb, "", "  ") // "  " for 2-space indentation
	if err != nil {
		panic(fmt.Errorf("failed to marshal RoleBinding to JSON: %w", err))
	}

	var una string
	//fmt.Print(string(jsonData))
	for _, rba := range rb.Items {
		// newRoler[rba.Subjects[0].Name]
		if strings.Contains(string(jsonData), *uname) {
			// fmt.Printf("In Question: %v\n\n", *uname)

			una = *uname
			roler[rba.RoleRef.Name]++
			//fmt.Printf("Service Account: %v has the following roles: \n%v \n",*uname, rba.RoleRef.Name)
		//	fmt.Printf("------------------------------\n")
			// try a struct instead of a map, also looping through all roles is long, try to do a check like
			// service account: has x roles...
		}
	}
	fmt.Printf("\nFor the SA: %v\n", una)
	fmt.Println("Has the Following roles")
	h:=0
	for r, _ := range roler {
		h++
		fmt.Printf("%v) %v\n",h, r)

	} 


}
