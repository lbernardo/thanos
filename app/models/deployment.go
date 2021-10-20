package models

import "fmt"

func NewDeployment(name, image string, replicas, port int64) *Deployment {
	return &Deployment{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
		Metadata: DeploymentMetadata{
			Name:      fmt.Sprintf("deploy-%v", name),
			Namespace: "default",
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: DeploymentSpec{
			Replicas: replicas,
			Selector: Selector{
				MatchLabels: Labels{
					App: name,
				},
			},
			Template: Template{
				Metadata: TemplateMetadata{
					Labels: Labels{
						App: name,
					},
				},
				Spec: TemplateSpec{
					Containers: []Container{
						{
							Image:           image,
							ImagePullPolicy: "Always",
							Name:            fmt.Sprintf("container-%v", name),
							Ports: []Port{
								{
									ContainerPort: port,
								},
							},
						},
					},
					ImagePullSecrets: []ImagePullSecret{
						{
							Name: "gitlab",
						},
					},
					RestartPolicy: "Always",
				},
			},
		},
	}
}

type Deployment struct {
	APIVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Metadata   DeploymentMetadata `json:"metadata"`
	Spec       DeploymentSpec     `json:"spec"`
}

type DeploymentMetadata struct {
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
}

type Labels struct {
	App string `json:"app"`
}

type DeploymentSpec struct {
	Replicas int64    `json:"replicas"`
	Selector Selector `json:"selector"`
	Template Template `json:"template"`
}

type Selector struct {
	MatchLabels Labels `json:"matchLabels"`
}

type Template struct {
	Metadata TemplateMetadata `json:"metadata"`
	Spec     TemplateSpec     `json:"spec"`
}

type TemplateMetadata struct {
	Labels Labels `json:"labels"`
}

type TemplateSpec struct {
	Containers       []Container       `json:"containers"`
	ImagePullSecrets []ImagePullSecret `json:"imagePullSecrets"`
	RestartPolicy    string            `json:"restartPolicy"`
}

type Container struct {
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Name            string `json:"name"`
	Ports           []Port `json:"ports"`
}

type Port struct {
	ContainerPort int64 `json:"containerPort"`
}

type ImagePullSecret struct {
	Name string `json:"name"`
}
