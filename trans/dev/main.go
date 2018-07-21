package main

import "github.com/kris-nova/krex/trans"

func main() {
	window, _ := trans.GetNewWindow(trans.DefaultHeight, trans.DefaultWidth)
	window.StartScreen("Start message")
	window.Prompt([]string{"item 1", "item 2", "item 3"})
}
