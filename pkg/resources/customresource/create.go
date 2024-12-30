package customresource

import (
	"context"
	"fmt"
	"log"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

const customPod string = "CustomPod"

type CreateCustomResourceOptions struct {
	Name string
	Data *map[string]interface{}

	namespace string
	group     string // API group i.e. your domain/organization
	version   string
	kind      string // camel case
	labels    map[string]string
}

func NewCreateCustomResouceOptions(appName, kind, containerName, image string, port int32) CreateCustomResourceOptions {
	opts := &CreateCustomResourceOptions{
		Name:      appName,
		namespace: "default",
		group:     "k8scli.io",
		version:   "v1",
		kind:      customPod,
		labels:    map[string]string{"app": appName},
	}
	opts.addData(kind, containerName, image, port)
	return *opts
}

func (o *CreateCustomResourceOptions) addData(kind, name, image string, port int32) {
	switch kind {
	case customPod:
		o.createDataCustomPod(name, image, port)
	default:
		log.Fatal("unknown custom resource type")
	}
}

func (o *CreateCustomResourceOptions) createDataCustomPod(name, image string, port int32) {
	o.Data = &map[string]interface{}{
		"containers": []map[string]interface{}{
			{
				"name":  name,
				"image": image,
				"ports": []map[string]interface{}{
					{
						"containerPort": port,
						"protocol":      "TCP",
					},
				},
			},
		},
	}
}

func (o CreateCustomResourceOptions) CreateCustomResource(client dynamic.Interface) error {
	gvr := schema.GroupVersionResource{
		Group:    o.group,
		Version:  o.version,
		Resource: strings.ToLower(o.kind) + "s",
	}

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": o.group + "/" + o.version,
			"kind":       o.kind,
			"metadata": map[string]interface{}{
				"name":      o.Name,
				"namespace": o.namespace,
				"labels":    o.labels,
			},
			"spec": o.Data,
		},
	}

	_, err := client.Resource(gvr).Namespace(o.namespace).Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create customer resource: %w", err)
	}

	fmt.Println("Created custom resource successfully!")
	return nil
}
