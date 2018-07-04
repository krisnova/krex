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

package runtime

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type RuntimeOptions struct {
	KubeconfigPath string
	ShellExecImage string
}

type Runtime struct {
	options   *RuntimeOptions
	root      *Vertex
	clientset *kubernetes.Clientset
}

func NewRuntime(opt *RuntimeOptions, root *Vertex) *Runtime {
	return &Runtime{
		options: opt,
		root:    root,
	}
}

func (r *Runtime) Init() error {
	config, err := clientcmd.BuildConfigFromFlags("", r.options.KubeconfigPath)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	r.clientset = clientset
	runtimeInstance = r
	return nil
}

func (r *Runtime) Walk() error {
	return r.root.RecursiveSelect()
}

var runtimeInstance *Runtime
