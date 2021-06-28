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
