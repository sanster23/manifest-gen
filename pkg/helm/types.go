package helm

// HelmOptions defines all the options that helm chart can take
type HelmOptions struct {
	Name            string
	Image           string
	ImageTag        string
	ImagePullPolicy string
	ServiceType     string
	ServicePort     int
	Starter         string // --starter
	StarterDir      string
}

// SVC PORTS:
// * -	Multiple
// * -	Port Name/target port
//           Image:
//           image Version:  latest|default
//           Replicas:
//           ENV VARIABLES: PER ENV/ GLOBAL
//           Labels:  DEFAULT/APP NAME, TEAM NAME
//           Resource requests/limits – DEFAULTS?
//          Namespace:
//           LIVENESS/READINESS ( probe timeout

// Deployment defines the sub-config struct for manifest-generator
type Deployment struct {
	Image       string            `json:"image,omitempty"`
	ImageTag    string            `json:"imageTag,omitempty"`
	Replicas    int               `json:"replicas,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	// Resources
}
type Service struct {
}

type Config struct {
	Deployment []*Deployment
	Service    []*Service
}
