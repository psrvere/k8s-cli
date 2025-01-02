package crd

import (
	"context"
	"fmt"
	"k8scli/pkg/client"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CustomResourceDefinition struct {
	Definition *apiextensionsv1.CustomResourceDefinition
}

func NewCustomResourceDefinition() CustomResourceDefinition {
	return CustomResourceDefinition{
		Definition: &apiextensionsv1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       "CustomResourceDefinition",
				APIVersion: apiextensionsv1.SchemeGroupVersion.Version,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "custompods.k8scli.io",
			},
			Spec: apiextensionsv1.CustomResourceDefinitionSpec{
				Group: "k8scli.io",
				Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
					{
						Name:    "v1",
						Served:  true,
						Storage: true,
						Schema: &apiextensionsv1.CustomResourceValidation{
							OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensionsv1.JSONSchemaProps{
									"spec": {
										Type: "object",
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"container": {
												Type: "array",
												Items: &apiextensionsv1.JSONSchemaPropsOrArray{
													Schema: &apiextensionsv1.JSONSchemaProps{
														Type: "object",
														Properties: map[string]apiextensionsv1.JSONSchemaProps{
															"name": {
																Type: "string",
															},
															"image": {
																Type: "string",
															},
															"ports": {
																Type: "array",
																Items: &apiextensionsv1.JSONSchemaPropsOrArray{
																	Schema: &apiextensionsv1.JSONSchemaProps{
																		Type: "object",
																		Properties: map[string]apiextensionsv1.JSONSchemaProps{
																			"containerPort": {
																				Type: "integer",
																			},
																			"protocol": {
																				Type: "string",
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Scope: apiextensionsv1.NamespaceScoped,
				Names: apiextensionsv1.CustomResourceDefinitionNames{
					Plural:     "custompods",
					Singular:   "custompod",
					Kind:       "CustomPod",
					ShortNames: []string{"cp"},
				},
			},
		},
	}
}

func (crd CustomResourceDefinition) CreateCustomSourceDefinition() error {
	client := client.GetKubeClient()
	_, err := client.ApiExtensionClient.ApiextensionsV1().CustomResourceDefinitions().Create(context.Background(), crd.Definition, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create custom resource definition: %v", err)
	}
	fmt.Println("Created custom resource deinition successfully!")
	return nil
}
