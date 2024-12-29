package deployments

import (
	"context"
	"fmt"
	"k8scli/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteDeploymentOptions struct {
	Name string
}

func NewDeleteDeploymentOptions(name string) DeleteDeploymentOptions {
	return DeleteDeploymentOptions{
		Name: name,
	}
}

func (o DeleteDeploymentOptions) DeleteDeployment() error {
	fmt.Println("In DeleteDeployment")
	deployClient := client.GetDeploymentClient()
	err := deployClient.Delete(context.TODO(), o.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Deletion of %v is Succesful!\n", o.Name)
	return nil
}
