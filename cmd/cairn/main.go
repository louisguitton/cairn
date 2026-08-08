package main

import "github.com/louisguitton/cairn/cmd"

var version = "dev"

func main() {
	cmd.SetVersion(cmd.ResolveVersion(version))
	cmd.Execute()
}
