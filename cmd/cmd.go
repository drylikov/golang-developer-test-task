package cmd

import "flag"

type Args struct {
	Env    string
	Source string
}

func (a *Args) Parse() {
	env := flag.String("env", "dev", "env for app")
	filePath := flag.String("source", "", "file source url or local file path")
	flag.Parse()

	a.Env = *env
	a.Source = *filePath
}
