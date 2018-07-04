// Copyright © 2018 Kris Nova <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/kris-nova/krex/explorer"
	"github.com/kubicorn/kubicorn/pkg/local"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "krex",
	Short: "Kubernetes Resource Explorer",
	Long:  `Explore Kubernetes resources like a boss.`,
	Run: func(cmd *cobra.Command, args []string) {

		//TODO start with cluster
		err := explorer.Init(opt)
		if err != nil {
			logger.Critical("error during init: %v", err)
			os.Exit(1)
		}
		clusterExplorer := &explorer.ClusterExplorer{}
		err = explorer.Explore(clusterExplorer)
		if err != nil {
			logger.Critical("error during explore: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

var opt = &explorer.RuntimeOptions{}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&opt.KubeconfigPath, "kubeconfig", "k", local.Expand("~/.kube/config"), "The path to a kube config file on the local filesystem")
	RootCmd.PersistentFlags().IntVarP(&logger.Level, "verbosity", "v", 4, "Verbosity 0-4")
	RootCmd.Flags().StringVarP(&opt.ShellExecImage, "image", "i", "ubuntu:latest", "The container image to use for a shell debug")

}
