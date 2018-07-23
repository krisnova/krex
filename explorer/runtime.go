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

package explorer

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// Import to initialize client auth plugins.
	"fmt"

	"github.com/kris-nova/krex/trans"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
	k8sclient *kubernetes.Clientset
	options   *RuntimeOptions
	krexLogo  = `  _
 | |
 | | ___ __ _____  __
 | |/ / '__/ _ \ \/ /
 |   <| | |  __/>  <
 |_|\_\_|  \___/_/\_\

`

	// transXY (aside from being the best variable name in the universe)
	// is where we map the transport system used to navigate the terminal's
	// X and Y buffer. We wrap ncurses.h and draw our own terminal buffer
	// from scratch. This window needs to be handled like a delicate flower
	// as if it does not exit cleanly we can skew the user's terminal buffer
	// and cause really nasty side effects.
	transXY *trans.TransWindow
)

type RuntimeOptions struct {
	KubeconfigPath string
	ShellExecImage string
}

func Init(opt *RuntimeOptions) error {
	var err error
	transXY, err = trans.GetNewWindow(trans.DefaultHeight, trans.DefaultWidth)
	defer transXY.End()
	if err != nil {
		return fmt.Errorf("unable to initialize trans system: %v", err)
	}
	transXY.StartScreen(krexLogo)
	config, err := clientcmd.BuildConfigFromFlags("", opt.KubeconfigPath)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	k8sclient = clientset
	options = opt
	return nil
}
