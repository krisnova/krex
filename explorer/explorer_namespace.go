package explorer

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
//namespaceLabel = "namespace"
)

type NamespaceExplorer struct {
	Items        []*MenuItem
	PreviousItem *MenuItem
}

func (n *NamespaceExplorer) Title() string {
	return "Cluster Level Resource Explorer"
}

func (n *NamespaceExplorer) List() error {
	ns, err := runtimeInstance.clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for _, item := range ns.Items {
		m := &MenuItem{}
		m.SetKind(namespaceLabel)
		m.SetName(item.Name)
		n.Items = append(n.Items, m)
	}
	// TODO add CRDs
	return nil
}

func (n *NamespaceExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select cluster resources: ", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *NamespaceExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case namespaceLabel:

	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
	return nil
}
