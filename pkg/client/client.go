package client

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	ClientSet *kubernetes.Clientset
}

func NewKubeClient(kubeconfig string) (*KubeClient, error) {
	configpath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", configpath)
	if err != nil {
		log.Fatalf("error in building cofig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error in creating new client: %v", err)
	}
	return &KubeClient{
		ClientSet: clientset,
	}, nil
}
