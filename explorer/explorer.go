package explorer

import "os"

func Explore(explorable Explorable) error {
	err := explorable.List()
	if err != nil {
		return err
		//return fmt.Errorf("error calling list: %v", err)
	}
	selection, err := explorable.RunPrompt()
	if err != nil {
		return err
		//return fmt.Errorf("error calling runprompt: %v", err)
	}
	err = explorable.Execute(selection)
	if err != nil {
		return err
		//return fmt.Errorf("error calling execute: %v", err)
	}
	return nil
}

func Exit() error {
	os.Exit(0)
	return nil
}
