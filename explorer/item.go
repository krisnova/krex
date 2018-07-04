package runtime

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
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

func (m *MenuItem) GetReadable() string {
	return fmt.Sprintf("%s * %s", m.kind, m.name)
}

func NewMenuItemFromReadable(readable string) *MenuItem {
	spl := strings.Split(readable, "*")
	return &MenuItem{
		kind: spl[0],
		name: spl[1],
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
