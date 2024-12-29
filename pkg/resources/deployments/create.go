package deployments

import (
	"context"
	"fmt"
	"k8scli/pkg/client"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CreateDeploymentOptions struct {
	Name          string
	Replicas      int32
	Image         string
	Protocol      corev1.Protocol
	ContainerPort int32

	Client *client.KubeClient
}

func NewDeploymentOptions() {

}

func (o CreateDeploymentOptions) CreateDeployment() {
	deployment := o.createDeployment()
	deployClient := client.GetDeploymentClient()
	deploy, err := deployClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error creating deployment: %v", err)
	}

	fmt.Println("Deployment Successful!")
	fmt.Println(deploy)
}

func (o CreateDeploymentOptions) createDeployment() *appsv1.Deployment {
	labels := map[string]string{"app": o.Name}
	selector := metav1.LabelSelector{MatchLabels: labels}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: appsv1.SchemeGroupVersion.Version, Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:   o.Name,
			Labels: labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &o.Replicas,
			Selector: &selector,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: o.buildPodSpec(),
			},
		},
	}
}

// only supports one image
func (o CreateDeploymentOptions) buildPodSpec() corev1.PodSpec {
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  o.Image, // uses full image name as container name
				Image: o.Image,
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: o.ContainerPort,
						Protocol:      o.Protocol,
					},
				},
			},
		},
	}
}
