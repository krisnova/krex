package trans

// Configuration options

var Cfg = &ConfigOptions{
	QuitChar: 'q',
}

type ConfigOptions struct {
	QuitChar rune
}
