package main

import (
	"context"
	"flag"
	"fmt"

	// "k8s.io/client-go" // exposes all the interfaces used by the client to interact with the API server read more: https://github.com/kubernetes/client-go
	// "k8s.io/api" // all the kubernetes resources are interacted with by the api library
	// "k8s.io/apiMachinery" // Access to utility methods that help developing an API
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/usr/local/google/home/saedf/.kube/config", "location of your kubeconfig file")

	// clientcmd library reads the config file and creates a configuration object 'config'
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		// handle error
		fmt.Printf("error %s, not able to build condif feom flags", err.Error())
	}
	// the clientset object is the remote control that allows the program to interact with the API server, using the configuration parsed to it
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		// handle error
		fmt.Printf("error %s, not ble to create clientset", err.Error())
	}

	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})

	if err != nil {
		// handle error
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}
	

	fmt.Printf("Pods in default Namespace:")
	
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}


	// fmt.Println("\nDeployments are\n")

	// deployment, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	
	// if err != nil {
	// 	// handle error
	// 	fmt.Printf("error %s, issues listing deployments", err.Error())
	// }	
	// for _, d := range deployment.Items {
	// 	fmt.Printf("%s\n", d.Name)
	// }

	// fmt.Println("\nSecrets Are")

	// secrets, err := clientset.CoreV1().Secrets("default").List(ctx, metav1.ListOptions{})

	// if err != nil {
	// 	// handle error
	// 	fmt.Printf("error %s, listing secrets", err.Error())
	// }

	// for _, s := range secrets.Items {
	// 	fmt.Printf("%s\n", s.Name)
	// }


}





// Intent

// Hardware security module (HSM) passcodes: Without passcodes you cannot decommission and reuse the StorageGRID and ONTAP appliances. The passcode unlocks access to the disks. See Hardware security module (HSM) bootstrap: Save credentials.
// Hardware security module(HSM) lockcodes: Provided to Thales if you must renew licenses. See Hardware security module (HSM) bootstrap: Apply license for instructions on how to get lock code.
// Hardware security module (HSM) cluster license: Provided by Thales after providing lock codes to them.
// cellconfig: Tracks all IPs allocated at allocation time.
// Root Admin kubeconfig - needed to manage Google Distributed Cloud (GDC) air-gapped root cluster if Active Directory Federation Services breaks. The root kubeconfig can also be used to get the org-admin kubeconfig and system cluster kubeconfigs.
// Licenses to the appliances - needed for reset.
