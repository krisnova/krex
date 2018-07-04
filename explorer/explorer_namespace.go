package explorer

import (
	"fmt"

	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	statefulSetLabel = "StatefulSet"
	deploymentLabel  = "Deployment"
	daemonSetLabel   = "DaemonSet"
	debugLabel       = "Debug"
	actionLabel      = "Action"
)

type NamespaceExplorer struct {
	Items              []*MenuItem
	PreviousItem       *MenuItem
	NamespaceToExplore string
	PreviousExplorer   Explorable
}

func (n *NamespaceExplorer) List() error {
	n.Items = []*MenuItem{}
	n.Items = AddEdit(n.PreviousItem, n.Items)

	// StatefulSet
	ss, err := k8sclient.AppsV1().StatefulSets(n.NamespaceToExplore).List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for _, item := range ss.Items {
		m := &MenuItem{}
		m.SetName(item.Name)
		m.SetKind(statefulSetLabel)
		n.Items = append(n.Items, m)
	}

	// Deployment
	ds, err := k8sclient.AppsV1().Deployments(n.NamespaceToExplore).List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for _, item := range ds.Items {
		m := &MenuItem{}
		m.SetName(item.Name)
		m.SetKind(deploymentLabel)
		n.Items = append(n.Items, m)
	}

	// DaemonSet
	dss, err := k8sclient.AppsV1().DaemonSets(n.NamespaceToExplore).List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for _, item := range dss.Items {
		m := &MenuItem{}
		m.SetName(item.Name)
		m.SetKind(daemonSetLabel)
		n.Items = append(n.Items, m)
	}

	m := &MenuItem{}
	m.SetKind(debugLabel)
	m.SetName(fmt.Sprintf("Run a debugging pod in the Namespace and shell exec [%s]", options.ShellExecImage))
	n.Items = append(n.Items, m)
	n.Items = AddGoBack(n.Items)
	return nil
}

func (n *NamespaceExplorer) RunPrompt() (string, error) {
	prompt := NewPromptFromMenuItems("Select application resources", n.Items)
	_, selection, err := prompt.Run()
	return selection, err
}

func (n *NamespaceExplorer) Execute(selection string) error {
	item := NewMenuItemFromReadable(selection)
	switch item.GetKind() {
	case daemonSetLabel:
		daemonSetExplorer := &DaemonSetExplorer{
			PreviousItem:       item,
			DaemonSetToExplore: item.name,
			NamespaceToExplore: n.NamespaceToExplore,
			PreviousExplorer:   n,
		}
		return Explore(daemonSetExplorer)
	case deploymentLabel:
		deploymentExplorer := &DeploymentExplorer{
			PreviousItem:        item,
			DeploymentToExplore: item.name,
			NamespaceToExplore:  n.NamespaceToExplore,
			PreviousExplorer:    n,
		}
		return Explore(deploymentExplorer)
	case statefulSetLabel:
		statefulSetExplorer := &StatefulSetExplorer{
			PreviousItem:         item,
			StatefulSetToExplore: item.name,
			NamespaceToExplore:   n.NamespaceToExplore,
			PreviousExplorer:     n,
		}
		return Explore(statefulSetExplorer)
	case actionLabel:
		if strings.Contains(item.GetName(), "../") {
			return Explore(n.PreviousExplorer)
		}
		return fmt.Errorf("unknown action selection: %s", selection)
	case debugLabel:
		// Deploy a pod and exec into it
		Exec("kubectl", []string{"run", "-n", n.NamespaceToExplore, "-i", "--tty", "krex-debug-pod", "--image", options.ShellExecImage, "--", "sh"})
		return Explore(n)
	case editLabel:
		Exec("kubectl", []string{"edit", "namespace", n.NamespaceToExplore})
		return Explore(n)
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
	return nil
}
