//Accessing the K8s cluster outside of the cluster itself to know more about various objects in the K8s cluster.

//Pre-requisites:
//1	Run a K8s cluster using miinikube
//2	Learning more about modules in golang

package main

import (
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)
	api := clientset.CoreV1()
	pods, _ := api.Pods("").List(v1.ListOptions{})

	//Listing out all the pods in the cluster
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
