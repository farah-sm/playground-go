package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	startPath := flag.String("p", ".", "Where shall we scan, default is root")
	pod_in_q := flag.String("pn", "", "-pn='saeds'")
	resource := flag.String("r", "", "-r='pod'")
	flag.Parse()

	if *startPath == "" {
		fmt.Println("Usage: go run file.go -path /path/to/dir")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Scanning for config files in %s...\n", *startPath)

	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.Contains(info.Name(), "config") {

			clientSet(path, *pod_in_q, *resource)

		} 

		return nil
})
    if err != nil {
        println("Error", err)
    }

}


func clientSet(path string, p string, r string) {

		config, err := clientcmd.BuildConfigFromFlags("", path)
		if err != nil {
			fmt.Printf("error %s, not able to build config from flags", err.Error())
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("error %s, not ble to create clientset", err.Error())
		}
	
		ctx := context.Background()
		
		if strings.Contains(r, "pod") {
			get_pods(clientset, ctx, "Pods", p)
		}
		if strings.Contains(r, "svc") {
			get_services(clientset, ctx, "Services", p)

		} 
		if r == "secrets" {
			get_secrets(clientset, ctx, "Secrets", p)

		} 
		if r == "ns" {
			get_ns(clientset, ctx, "Namespaces", p)

		}
		if r == "cm" {
			get_cm(clientset, ctx, "ConfigMaps", p)

		}
		if r == "nodes" {
			get_nodes(clientset, ctx, "Nodes", p)
		}
		if r == "" {
			fmt.Println("Sorry be sure to specify the resource you're looking for.")
			flag.PrintDefaults()
		}




	
}

func get_pods(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nPod Name: %v\nPOD IP: %v\nNamespace: %v\n", pod.Name, pod.Status.PodIP, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}

func get_secrets(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing secrets from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nSecret Name: %v\nLabels: %v\nNamespace: %v\n", pod.Name, pod.Labels, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}

func get_ns(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nNS Name: %v\nNS Label: %v\nNamespace: %v\n", pod.Name, pod.Labels, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}

func get_services(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nService Name: %v\nCluster IP: %v\nNamespace: %v\n", pod.Name, pod.Spec.ClusterIP, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}

func get_cm(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nCM Name: %v\nCM Label IP: %v\nNamespace: %v\n", pod.Name, pod.Labels, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}

func get_nodes(clientset *kubernetes.Clientset, ctx context.Context, name string, p string) {
	fmt.Printf("\n%s in with '%v' in name: \n", name, p)
	pods, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error %s, listing pods from namespace", err.Error())
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, p) {
			fmt.Printf("-----------------\n-----------------\nNode Name: %v\nPOD CIDRs: %v\nNamespace: %v\n", pod.Name, pod.Spec.PodCIDRs, pod.Namespace)
		}
	}
	fmt.Println("-----------------\n-----------------\n")


}