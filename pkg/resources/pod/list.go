package pod

import (
	"context"
	"fmt"
	"k8scli/pkg/client"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListPodOptions struct {
}

func NewListPodOption(name string) *ListPodOptions {
	return &ListPodOptions{}
}

func (o *ListPodOptions) ListPods() (*v1.PodList, error) {
	client := client.GetKubeClient()
	pods, err := client.ClientSet.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing pod: %v", err)
	}
	return pods, nil
}
