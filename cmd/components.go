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
	componentsName         string
	componentsOutputFormat string
)

var ComponentsCmd = &cobra.Command{
	Use:   "components",
	Short: "List all Dapr components. Supported platforms: Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		if kubernetesMode {
			if allNamespaces {
				resourceNamespace = meta_v1.NamespaceAll
			}
			err := kubernetes.PrintComponents(componentsName, resourceNamespace, componentsOutputFormat)
			if err != nil {
				print.FailureStatusEvent(os.Stdout, err.Error())
				os.Exit(1)
			}
		}
	},
	Example: `
# List Kubernetes Dapr components
dapr components -k

# List define namespace Dapr components in Kubernetes mode
dapr components -k -n default

# List all namespaces Dapr components in Kubernetes mode
dapr components -k --all-namespaces
`,
}

func init() {
	ComponentsCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "If true, list all Dapr pods in all namespaces")
	ComponentsCmd.Flags().StringVarP(&componentsName, "name", "", "", "The components name to be printed (optional)")
	ComponentsCmd.Flags().StringVarP(&resourceNamespace, "namespace", "n", "default", "List Define namespace pod in a Kubernetes cluster")
	ComponentsCmd.Flags().StringVarP(&componentsOutputFormat, "output", "o", "list", "Output format (options: json or yaml or list)")
	ComponentsCmd.Flags().BoolVarP(&kubernetesMode, "kubernetes", "k", false, "List all Dapr components in a Kubernetes cluster")
	ComponentsCmd.Flags().BoolP("help", "h", false, "Print this help message")
	ComponentsCmd.MarkFlagRequired("kubernetes")
	RootCmd.AddCommand(ComponentsCmd)
}
