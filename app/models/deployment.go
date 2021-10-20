package models

func NewDeployment(name, image string, replicas, port int64) *Deployment {
	return &Deployment{
		Kind:       "Deployment",
		APIVersion: "v1",
		Metadata: DeploymentMetadata{
			Name: name,
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
						Container{
							Name:  name,
							Image: image,
							Ports: []Port{
								Port{
									ContainerPort: port,
								},
							},
							ImagePullPolicy: "Always",
						},
					},
					ImagePullSecrets: []ImagePullSecret{
						ImagePullSecret{
							Name: "gitlab",
						},
					},
				},
			},
		},
	}
}

type Deployment struct {
	Kind       string             `json:"kind"`
	APIVersion string             `json:"apiVersion"`
	Metadata   DeploymentMetadata `json:"metadata"`
	Spec       DeploymentSpec     `json:"spec"`
}

type DeploymentMetadata struct {
	Name string `json:"name"`
}

type DeploymentSpec struct {
	Replicas int64    `json:"replicas"`
	Selector Selector `json:"selector"`
	Template Template `json:"template"`
}

type Selector struct {
	MatchLabels Labels `json:"matchLabels"`
}

type Labels struct {
	App string `json:"app"`
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
}

type Container struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	Ports           []Port `json:"ports"`
	ImagePullPolicy string `json:"imagePullPolicy"`
}

type Port struct {
	ContainerPort int64 `json:"containerPort"`
}

type ImagePullSecret struct {
	Name string `json:"name"`
}
