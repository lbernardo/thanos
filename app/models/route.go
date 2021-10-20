package models

import "fmt"

func NewRoute(name, host string, port int64) *Route {
	return &Route{
		Kind:       "IngressRoute",
		APIVersion: "traefik.containo.us/v1alpha1",
		Metadata: RouteMetadata{
			Name:      fmt.Sprintf("ingress-route-%v", name),
			Namespace: "default",
		},
		Spec: RouteSpec{
			EntryPoints: []string{
				"web",
			},
			Routes: []RouteElement{
				{
					Match: fmt.Sprintf("Host(`%v`) && PathPrefix(`/`)", host),
					Kind:  "Rule",
					Services: []RouteElementService{
						{
							Name: fmt.Sprintf("service-%v", name),
							Port: port,
						},
					},
				},
			},
		},
	}
}

type Route struct {
	APIVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Metadata   RouteMetadata `json:"metadata"`
	Spec       RouteSpec     `json:"spec"`
}

type RouteMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type RouteSpec struct {
	EntryPoints []string       `json:"entryPoints"`
	Routes      []RouteElement `json:"routes"`
}

type RouteElement struct {
	Kind     string                `json:"kind"`
	Match    string                `json:"match"`
	Services []RouteElementService `json:"services"`
}

type RouteElementService struct {
	Name string `json:"name"`
	Port int64  `json:"port"`
}
