package explorer

import (
	"fmt"
	"strings"

	"os"

	"github.com/manifoldco/promptui"
)

const (
	editLabel = "Edit"
)

type MenuItem struct {
	kind string
	name string
}

func (m *MenuItem) SetKind(t string) {
	m.kind = t
}

func (m *MenuItem) SetName(t string) {
	m.name = t
}

const rowSize = 20

func (m *MenuItem) GetReadable() string {
	kindLen := len(m.GetKind()) + 2
	whitespace := " "
	for kindLen < rowSize {
		whitespace = fmt.Sprintf("%s ", whitespace)
		kindLen++
	}
	return fmt.Sprintf("[%s]%s%s", m.GetKind(), whitespace, m.GetName())
}

func NewMenuItemFromReadable(readable string) *MenuItem {
	readable = strings.Replace(readable, "[", "", 1)
	spl := strings.Split(readable, "]")
	return &MenuItem{
		kind: strings.TrimSpace(spl[0]),
		name: strings.TrimSpace(spl[1]),
	}
}

func (m *MenuItem) GetKind() string {
	return m.kind
}
func (m *MenuItem) GetName() string {
	return m.name
}

func NewPromptFromMenuItems(title string, items []*MenuItem) promptui.Select {
	var strs []string
	for _, item := range items {
		strs = append(strs, item.GetReadable())
	}
	p := promptui.Select{
		Label: title,
		Items: strs,
		Size:  len(strs),
	}
	return p
}

func AddGoBack(items []*MenuItem) []*MenuItem {
	m := &MenuItem{}
	m.SetKind(actionLabel)
	m.SetName("Go back ../")
	return append(items, m)
}

func AddEdit(prev *MenuItem, items []*MenuItem) []*MenuItem {
	m := &MenuItem{}
	m.SetKind(editLabel)
	m.SetName(fmt.Sprintf("%s %s (%s)", prev.GetKind(), prev.GetName(), os.Getenv("EDITOR")))
	return append(items, m)
}
