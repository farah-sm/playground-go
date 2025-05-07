package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")
	//uname := flag.String("ldap", "", "-ldap='username'")
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

	fmt.Printf("------------------------------\n")
	get_pods(clientset, ctx,"Pods")
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

// where is this IP address, 