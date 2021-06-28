package cmd

import (
	"github.com/sanster23/manifest-gen/pkg/helm"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(newHelmCommand())
}

const helmLongDesc = `Long description`

func newHelmCommand() *cobra.Command {
	o := &helm.HelmOptions{}

	var cmd = &cobra.Command{
		Use:   "helm",
		Short: "Generate helm chart manifests",
		Long:  helmLongDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	f := cmd.Flags()
	f.StringVar(&o.Name, "name", "hello-world", "name of the helm chart")
	f.StringVar(&o.Image, "image", "", "docker image")
	f.StringVar(&o.ImageTag, "image-tag", "latest", "docker image tag")
	f.StringVar(&o.ImagePullPolicy, "image-pull-policy", "Always", "image pull policy for containers")
	f.StringVar(&o.ServiceType, "service-type", "ClusterIP", "service type [ClusterIP, NodePort, LoadBalancer] ")
	f.IntVar(&o.ServicePort, "service-port", 80, "service port")

	return cmd

}
