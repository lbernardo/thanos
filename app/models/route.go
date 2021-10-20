package models

import "fmt"

func NewRoute(name, host string) *Route {
	return &Route{
		Kind:       "IngressRoute",
		APIVersion: "traefik.containo.us/v1alpha1",
		Metadata: RouteMetadata{
			Name: fmt.Sprintf("ingress-route-%v", name),
		},
		Spec: RouteSpec{
			Entrypoints: []string{
				"web",
			},
			Routes: []Routes{
				{
					Match: fmt.Sprintf("Host(`%v`) && PathPrefix(`/`)", host),
					Kind:  "Rule",
					Services: []ServiceRoutes{
						{
							Name: fmt.Sprintf("service-%v", name),
							Port: "80",
						},
					},
				},
			},
		},
	}
}

type Route struct {
	Kind       string        `json:"kind"`
	APIVersion string        `json:"apiVersion"`
	Metadata   RouteMetadata `json:"metadata"`
	Spec       RouteSpec     `json:"spec"`
}

type RouteMetadata struct {
	Name string `json:"name"`
}

type RouteSpec struct {
	Entrypoints []string `json:"entrypoints"`
	Routes      []Routes `json:"routes"`
}

type Routes struct {
	Match    string          `json:"match"`
	Kind     string          `json:"kind"`
	Services []ServiceRoutes `json:"services"`
}

type ServiceRoutes struct {
	Name string `json:"name"`
	Port string `json:"port"`
}
