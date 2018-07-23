package main

import (
	"fmt"

	"github.com/kris-nova/krex/trans"
)

func main() {
	window, _ := trans.GetNewWindow(trans.DefaultHeight, trans.DefaultWidth)
	window.StartScreen("Start message")
	selection := window.Prompt("A great title", []string{"item 1", "item 2", "item 3"})
	fmt.Println(selection)
}
