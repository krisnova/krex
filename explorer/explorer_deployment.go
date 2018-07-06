package explorer

import (
	"fmt"
	"strings"
)

const ()

type DeploymentExplorer struct {
	Items               []*MenuItem
	PreviousItem        *MenuItem
	NamespaceToExplore  string
	DeploymentToExplore string
	PreviousExplorer    Explorable
}

func (n *DeploymentExplorer) List() error {
	n.Items = []*MenuItem{}
	n.Items = AddEdit(n.PreviousItem, n.Items)
	m := &MenuItem{}
	m.SetKind(podsLabel)
	m.SetName("Get Pods")
	n.Items = append(n.Items, m)
	n.Items = AddGoBack(n.Items)
	return nil
}

func (n *DeploymentExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select deployment resources: ", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *DeploymentExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case podsLabel:
		podsExplorer := &PodsExplorer{
			PreviousItem: item,
			Filters: map[string]string{
				"k8s-app": n.DeploymentToExplore,
			},
			NamespaceToExplore: n.NamespaceToExplore,
			PreviousExplorer:   n,
		}
		return Explore(podsExplorer)
	case editLabel:
		Exec("kubectl", []string{"edit", "deployment", n.DeploymentToExplore, "-n", n.NamespaceToExplore})
		return Explore(n)
	case actionLabel:
		if strings.Contains(item.GetName(), "../") {
			return Explore(n.PreviousExplorer)
		}
		return fmt.Errorf("unknown action selection: %s", selection)
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
}

func (n *DeploymentExplorer) Kind() string {
	return "deployment"
}
