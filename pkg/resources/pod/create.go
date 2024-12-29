package pod

import (
	"context"
	"fmt"
	"k8scli/pkg/client"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CreatePodOptions struct {
	Name          string
	Image         string
	ContainerPort int32
}

func NewCreatePodOptions(name, image string, port int32) CreatePodOptions {
	return CreatePodOptions{
		Name:          name,
		Image:         image,
		ContainerPort: port,
	}
}

func (o CreatePodOptions) CreatePod() error {
	client := client.GetKubeClient()
	podSpec := o.getPodSpecs()
	_, err := client.ClientSet.CoreV1().Pods("default").Create(context.TODO(), podSpec, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("error creating pod: %v", err)
	}
	fmt.Printf("%v Pod creation Successful!\n", o.Name)
	return nil
}

func (o CreatePodOptions) getPodSpecs() *v1.Pod {
	labels := map[string]string{"app": o.Name}
	return &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   o.Name,
			Labels: labels,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyAlways,
			Containers: []v1.Container{
				{
					Name:  o.Name,
					Image: o.Image,
					Ports: []v1.ContainerPort{
						{
							ContainerPort: o.ContainerPort,
							Protocol:      v1.ProtocolTCP,
						},
					},
				},
			},
		},
	}
}
