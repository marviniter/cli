// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation and Dapr Contributors.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

import (
	"os"

	"github.com/dapr/cli/pkg/kubernetes"
	"github.com/dapr/cli/pkg/print"
	"github.com/spf13/cobra"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	configurationName         string
	configurationOutputFormat string
)

var ConfigurationsCmd = &cobra.Command{
	Use:   "configurations",
	Short: "List all Dapr configurations. Supported platforms: Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		if kubernetesMode {
			if allNamespaces {
				resourceNamespace = meta_v1.NamespaceAll
			}
			err := kubernetes.PrintConfigurations(configurationName, resourceNamespace, configurationOutputFormat)
			if err != nil {
				print.FailureStatusEvent(os.Stdout, err.Error())
				os.Exit(1)
			}
		}
	},
	Example: `
# List Kubernetes Dapr configurations
dapr configurations -k

# List define namespace Dapr configurations in Kubernetes mode
dapr configurations -k -n default

# List all namespaces Dapr configurations in Kubernetes mode
dapr configurations -k --all-namespaces
`,
}

func init() {
	ConfigurationsCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "If true, list all Dapr pods in all namespaces")
	ConfigurationsCmd.Flags().StringVarP(&configurationName, "name", "", "", "The configuration name to be printed (optional)")
	ConfigurationsCmd.Flags().StringVarP(&resourceNamespace, "namespace", "n", "default", "List Define namespace pod in a Kubernetes cluster")
	ConfigurationsCmd.Flags().StringVarP(&configurationOutputFormat, "output", "o", "list", "Output format (options: json or yaml or list)")
	ConfigurationsCmd.Flags().BoolVarP(&kubernetesMode, "kubernetes", "k", false, "List all Dapr configurations in a Kubernetes cluster")
	ConfigurationsCmd.Flags().BoolP("help", "h", false, "Print this help message")
	ConfigurationsCmd.MarkFlagRequired("kubernetes")
	RootCmd.AddCommand(ConfigurationsCmd)
}
