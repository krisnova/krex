package explorer

import (
	"fmt"
	"strings"

	"os"
)

const (
	editLabel = "Edit"
	exitLabel = "Exit"
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

func AddExit(items []*MenuItem) []*MenuItem {
	m := &MenuItem{}
	m.SetKind(exitLabel)
	m.SetName("Exit")
	return append(items, m)
}

func checkExitItem(item string) {
	if item == "" {
		Exit()
	}
}
