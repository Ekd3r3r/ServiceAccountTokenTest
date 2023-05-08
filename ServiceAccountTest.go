package main

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		err := listPods(clientset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing pods: %v\n", err)
		}
		time.Sleep(5 * time.Second)
	}
}

func listPods(clientset *kubernetes.Clientset) error {
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Listing Pods:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s (%s)\n", pod.Name, pod.Status.Phase)
	}

	return nil
}
