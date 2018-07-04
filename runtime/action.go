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

import "github.com/kubicorn/kubicorn/pkg/logger"

type Action func(params *ActionParametes) error

type ActionParametes struct {
	PodName string
}

func ActionEmpty(params *ActionParametes) error {
	logger.Info("Calling [EMPTY] %s", params.PodName)
	return nil
}

func ActionEdit(params *ActionParametes) error {
	logger.Info("Calling [EDIT] %s", params.PodName)
	// TODO Implement this
	return nil
}

func ActionDescribe(params *ActionParametes) error {
	logger.Info("Calling [DESCRIBE] %s", params.PodName)
	// TODO Implement this
	return nil
}

func ActionLogs(params *ActionParametes) error {
	logger.Info("Calling [LOGS] %s", params.PodName)
	// TODO Implement this
	return nil
}

func ActionContainers(params *ActionParametes) error {
	logger.Info("Calling [CONTAINERS] %s", params.PodName)
	// TODO Implement this
	return nil
}

func ActionShellDebug(params *ActionParametes) error {
	logger.Info("Calling [SHELL DEBUG] %s", params.PodName)
	// TODO Implement this
	return nil
}
