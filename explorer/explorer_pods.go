package explorer

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	podLabel = "Pod"
)

type PodsExplorer struct {
	Items                []*MenuItem
	PreviousItem         *MenuItem
	NamespaceToExplore   string
	Filters              map[string]string
	PreviousExplorer     Explorable
	PreviousResourceName string
}

func (n *PodsExplorer) List() error {
	n.Items = []*MenuItem{}

	// List Pods
	pods, err := k8sclient.Core().Pods(n.NamespaceToExplore).List(v1.ListOptions{})
	if err != nil {
		return err
	}

	var podlist []apiv1.Pod
	switch n.PreviousExplorer.Kind() {
	case "deployment":
		// get replicasets
		replicaSetMap := make(map[string]bool)
		replicaSets, err := k8sclient.Extensions().ReplicaSets(n.NamespaceToExplore).List(v1.ListOptions{})
		if err != nil {
			return err
		}

		for _, replicaset := range replicaSets.Items {
			for _, owner := range replicaset.GetOwnerReferences() {
				if strings.ToLower(owner.Name) == n.PreviousResourceName {
					replicaSetMap[replicaset.GetName()] = true
					break
				}
			}
		}

		// get pods
		podlist = getOwnerPods(pods, replicaSetMap, "replicaset")

	case "statefulset":
		podlist = getOwnerPods(pods, map[string]bool{n.PreviousResourceName: true}, "statefulset")

	case "daemonset":
		podlist = getOwnerPods(pods, map[string]bool{n.PreviousResourceName: true}, "daemonset")
	}

	for _, item := range podlist {
		m := &MenuItem{}
		m.SetName(item.Name)
		m.SetKind(podLabel)
		n.Items = append(n.Items, m)
	}
	n.Items = AddGoBack(n.Items)
	n.Items = AddExit(n.Items)
	return nil
}

// getOwnerPods is a functio that gets pods from map of owners
func getOwnerPods(pods *apiv1.PodList, owners map[string]bool, resourceKindToMatch string) []apiv1.Pod {
	var podlist []apiv1.Pod

	for _, pod := range pods.Items {
		for _, owner := range pod.GetOwnerReferences() {
			if _, ok := owners[owner.Name]; ok && strings.ToLower(owner.Kind) == resourceKindToMatch {
				podlist = append(podlist, pod)
				break
			}
		}
	}

	return podlist
}

func (n *PodsExplorer) RunPrompt() (string, error) {
	var strs []string
	for _, item := range n.Items {
		strs = append(strs, item.GetReadable())
	}
	selection := transXY.Prompt("Select Pod resource", strs)
	return selection, nil
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
	case exitLabel:
		return Exit()
	default:
		return fmt.Errorf("unable to parse selection: %s", selection)
	}
}

func (n *PodsExplorer) Kind() string {
	return "pods"
}
