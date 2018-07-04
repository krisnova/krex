package explorer

type Explorable interface {
	//Title() string
	List() error
	RunPrompt() (string, error)
	Execute(selection string) error
}
