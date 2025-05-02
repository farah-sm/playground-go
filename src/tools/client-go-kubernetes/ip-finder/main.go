package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// kubectl api-resources --verbs=list -o name | paste -d, -s | xargs kubectl get --ignore-not-found --show-kind -oyaml | grep ip-addr

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")
	uname := flag.String("ip", "", "-ldap='username'")
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

	pods, err := clientset.ResourceV1beta1().
}