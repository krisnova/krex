package runtime

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespaceLabel = "namespace"
)

type ClusterExplorer struct {
	Items []*MenuItem
}

func (n *ClusterExplorer) Title() string {
	return "Cluster Level Resource Explorer"
}

func (n *ClusterExplorer) List() error {
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

func (n *ClusterExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select cluster resources: ", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *ClusterExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case namespaceLabel:
		namespaceExplorer := &NamespaceExplorer{
			PreviousItem: item,
		}
		return Explore(namespaceExplorer)
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
	return nil
}
