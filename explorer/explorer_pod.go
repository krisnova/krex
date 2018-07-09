package explorer

import (
	"fmt"
	"os"
	"strings"
)

const (
	logsLabel           = "Logs"
	execLabel           = "Exec"
	describeLabel       = "Describe"
	logAndDescribeLabel = "Logs and Describe"
)

type PodExplorer struct {
	Items              []*MenuItem
	PreviousItem       *MenuItem
	NamespaceToExplore string
	PodToExplore       string
	PreviousExplorer   Explorable
}

func (n *PodExplorer) List() error {
	n.Items = []*MenuItem{}

	n.Items = AddEdit(n.PreviousItem, n.Items)

	n.Items = append(n.Items, &MenuItem{
		kind: logsLabel,
		name: "Tail logs",
	})

	n.Items = append(n.Items, &MenuItem{
		kind: execLabel,
		name: "Shell exec (sh) into Pod",
	})

	n.Items = append(n.Items, &MenuItem{
		kind: describeLabel,
		name: "Describe the Pod",
	})

	n.Items = append(n.Items, &MenuItem{
		kind: logAndDescribeLabel,
		name: "Describe the Pod and then tail the logs",
	})

	// Logs
	// Exec
	// Describe
	// Log and Describe

	n.Items = AddGoBack(n.Items)
	n.Items = AddExit(n.Items)
	return nil
}

func (n *PodExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select pod resources: ", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *PodExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case logAndDescribeLabel:
		Exec("kubectl", []string{"describe", "pod", "--namespace", n.NamespaceToExplore, n.PodToExplore})
		Exec("kubectl", []string{"logs", n.PodToExplore, "-n", n.NamespaceToExplore, "-f"})
		return Explore(n)
	case describeLabel:
		Exec("kubectl", []string{"describe", "pod", "--namespace", n.NamespaceToExplore, n.PodToExplore})
		return Explore(n)
	case execLabel:
		Exec("kubectl", []string{"exec", "-it", "--namespace", n.NamespaceToExplore, n.PodToExplore, "sh"})
		return Explore(n)
	case logsLabel:
		Exec("kubectl", []string{"logs", n.PodToExplore, "-n", n.NamespaceToExplore, "-f"})
		return Explore(n)
	case editLabel:
		Exec("kubectl", []string{"edit", "pods", n.PodToExplore, "-n", n.NamespaceToExplore})
		return Explore(n)
	case actionLabel:
		if strings.Contains(item.GetName(), "../") {
			return Explore(n.PreviousExplorer)
		}
		return fmt.Errorf("unknown action selection: %s", selection)
	case exitLabel:
		os.Exit(0)
		return nil
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
}

func (n *PodExplorer) Kind() string {
	return "pod"
}
