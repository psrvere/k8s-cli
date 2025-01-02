package client

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	ClientSet          *kubernetes.Clientset
	DynamicClient      *dynamic.DynamicClient
	ApiExtensionClient apiextensionsclient.Interface
}

func NewKubeClient() (*KubeClient, error) {
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

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("error in creating dynamic client: %v", err)
	}

	apieClient, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("error in creating api extention client: %v", err)
	}

	return &KubeClient{
		ClientSet:          clientset,
		DynamicClient:      dynamicClient,
		ApiExtensionClient: apieClient,
	}, nil
}

func GetKubeClient() *KubeClient {
	c, err := NewKubeClient()
	if err != nil {
		log.Fatalf("error in getting KubeClient: %v", err)
	}
	return c
}

func GetDeploymentClient() v1.DeploymentInterface {
	c := GetKubeClient()
	return c.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
}
