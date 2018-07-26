package explorer

import (
	"fmt"
	"strings"
)

const (
	podsLabel = "Pods"
)

type DaemonSetExplorer struct {
	Items              []*MenuItem
	PreviousItem       *MenuItem
	NamespaceToExplore string
	DaemonSetToExplore string
	PreviousExplorer   Explorable
}

func (n *DaemonSetExplorer) List() error {
	n.Items = []*MenuItem{}
	n.Items = AddEdit(n.PreviousItem, n.Items)
	m := &MenuItem{}
	m.SetKind(podsLabel)
	m.SetName("Get Pods")
	n.Items = append(n.Items, m)
	n.Items = AddGoBack(n.Items)
	n.Items = AddExit(n.Items)
	return nil
}

func (n *DaemonSetExplorer) RunPrompt() (string, error) {
	var strs []string
	for _, item := range n.Items {
		strs = append(strs, item.GetReadable())
	}
	selection := transXY.Prompt("Select DaemonSet resources", strs)
	checkExitItem(selection)
	return selection, nil
}

func (n *DaemonSetExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case podsLabel:
		podsExplorer := &PodsExplorer{
			PreviousItem: item,
			Filters: map[string]string{
				"k8s-app": n.DaemonSetToExplore,
			},
			NamespaceToExplore:   n.NamespaceToExplore,
			PreviousExplorer:     n,
			PreviousResourceName: n.DaemonSetToExplore,
		}
		return Explore(podsExplorer)
	case editLabel:
		Exec("kubectl", []string{"edit", "daemonset", n.DaemonSetToExplore, "-n", n.NamespaceToExplore})
		return Explore(n)
	case actionLabel:
		if strings.Contains(item.GetName(), "../") {
			return Explore(n.PreviousExplorer)
		}
		return fmt.Errorf("unknown action selection: %s", selection)
	case exitLabel:
		return Exit()
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
}

func (n *DaemonSetExplorer) Kind() string {
	return "daemonset"
}
