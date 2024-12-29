package pod

import (
	"context"
	"fmt"
	"k8scli/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeletePodOptions struct {
	Name string
}

func NewDeletePodOption(name string) *DeletePodOptions {
	return &DeletePodOptions{
		Name: name,
	}
}

func (o *DeletePodOptions) DeletePod() error {
	client := client.GetKubeClient()
	err := client.ClientSet.CoreV1().Pods("default").Delete(context.TODO(), o.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("error deleting pod: %v", err)
	}
	fmt.Printf("%v pod deleted successfully!\n", o.Name)
	return nil
}
