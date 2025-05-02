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

	fmt.Printf("------------------------------\n")

	get_roles(clientset, ctx, "Roles", *uname)
	fmt.Printf("------------------------------\n")
	// get_pods(clientset, ctx, "pods")
}

func get_services_external_ip(clientset *kubernetes.Clientset, ctx context.Context, name string) {
	fmt.Printf("%s in Namespace: \n", name)
	services, err := clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing services", err.Error())
	}


	for _, s := range services.Items {
		fmt.Printf("%v\n", s.Spec.ClusterIP)
	}

}

func get_pods(clientset *kubernetes.Clientset, ctx context.Context, name string) {
	fmt.Printf("%s in Namespace: \n", name)
	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod: %v. The IP is: %v. The corresponding Node IP: %v\n", pod.Name, pod.Status.PodIP, pod.Status.HostIP)
	}


}

func get_events(clientset *kubernetes.Clientset, ctx context.Context, name string) {
	fmt.Printf("%s in Namespace: \n", name)
	event, err := clientset.CoreV1().Events("").List(ctx, metav1.ListOptions{})

	if err != nil {
		// handle error
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for i, e := range event.Items {
		i=i+1
		fmt.Printf("\n%d)Event: %v.\nTime: %v.\nMessage: %v.\n",i, e.Name, e.EventTime, e.Message)
	}
}


func get_roles(clientset *kubernetes.Clientset, ctx context.Context, name string, username *string) {
	username = flag.String("ldap", "", "-ldap='username'")
	flag.Parse()
	fmt.Printf("%s in Namespace: \n\n", name)
	
	rb, err := clientset.RbacV1().RoleBindings("").List(ctx, metav1.ListOptions{})
	if err != nil {
		// handle error
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}
	roler := make(map[string]int64)
	
	jsonData, err := json.MarshalIndent(rb, "", "  ") // "  " for 2-space indentation
	if err != nil {
		panic(fmt.Errorf("failed to marshal RoleBinding to JSON: %w", err))
	}
	//newRoler := make(map[string]any)

	for _, rba := range rb.Items {
		// newRoler[rba.Subjects[0].Name]
		if strings.Contains(string(jsonData), *username) {
			// rba.Subjects[0].Name
			roler[rba.RoleRef.Name]++

			fmt.Printf("Service Account: %v has the following roles: \n%v \n",*username, rba.RoleRef.Name)
			fmt.Printf("------------------------------\n")
			// try a struct instead of a map, also looping through all roles is long, try to do a check like
			// service account: has x roles...
		}

	}
}

// where is this IP address, 