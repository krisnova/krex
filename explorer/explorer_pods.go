package explorer

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	podLabel = "Pod"
)

type PodsExplorer struct {
	Items              []*MenuItem
	PreviousItem       *MenuItem
	NamespaceToExplore string
	Filters            map[string]string
	PreviousExplorer   Explorable
}

func (n *PodsExplorer) List() error {
	n.Items = []*MenuItem{}

	// Pods
	pods, err := k8sclient.CoreV1().Pods(n.NamespaceToExplore).List(v1.ListOptions{
		LabelSelector: labels.SelectorFromSet(n.Filters).String(),
	})
	if err != nil {
		return err
	}
	for _, item := range pods.Items {
		m := &MenuItem{}
		m.SetName(item.Name)
		m.SetKind(podLabel)
		n.Items = append(n.Items, m)
	}
	n.Items = AddGoBack(n.Items)
	return nil
}

func (n *PodsExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select pod resources: ", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *PodsExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case podLabel:
		podExplorer := &PodExplorer{
			PreviousItem:       item,
			PodToExplore:       item.GetName(),
			NamespaceToExplore: n.NamespaceToExplore,
			PreviousExplorer:   n,
		}
		return Explore(podExplorer)
	case actionLabel:
		if strings.Contains(item.GetName(), "../") {
			return Explore(n.PreviousExplorer)
		}
		return fmt.Errorf("unknown action selection: %s", selection)
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
}

func (n *PodsExplorer) Kind() string {
	return "pods"
}
