package main

import (
	"io"

	"github.com/spf13/cobra"
)

var createHelmDesc = `
This command creates helm chart with a predefined schema and given input.

The directory structure and starter manifests would be fixed and dynamic.
`

type helmOptions struct {
	name            string
	image           string
	imageTag        string
	imagePullPolicy string
	serviceType     string
	servicePort     int
}

func newCreateHelmCmd(out io.Writer) *cobra.Command {

	o := &helmOptions{}

	cmd := &cobra.Command{
		Use:   "create helm manifests",
		Short: "create a new helm chart",
		Long:  createHelmDesc,
	}

	f := cmd.Flags()
	f.StringVar(&o.name, "name", "hello-world", "name of the helm chart")
	f.StringVar(&o.image, "image", "", "docker image")
	f.StringVar(&o.imageTag, "image-tag", "latest", "docker image tag")
	f.StringVar(&o.imagePullPolicy, "image-pull-policy", "Always", "image pull policy for containers")
	f.StringVar(&o.serviceType, "service-type", "ClusterIP", "service type [ClusterIP, NodePort, LoadBalancer] ")
	f.IntVar(&o.servicePort, "service-port", 80, "service port")

	return cmd
}
