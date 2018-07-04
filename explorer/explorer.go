package runtime

import "fmt"

func Explore(explorable Explorable) error {
	fmt.Println(explorable.Title())
	err := explorable.List()
	if err != nil {
		return fmt.Errorf("error calling list: %v", err)
	}
	selection, err := explorable.RunPrompt()
	if err != nil {
		return fmt.Errorf("error calling runprompt: %v", err)
	}
	err = explorable.Execute(selection)
	if err != nil {
		return fmt.Errorf("error calling execute: %v", err)
	}
	return nil
}
