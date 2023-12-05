package main

import (
	"os"

	pygmy "github.com/WilkinsonK/pygmy/lib"
)

func EvalStatus(status pygmy.PyStatus) {
	if status.IsException() {
		if status.IsExit() {
			os.Exit(status.ExitCode())
		}
		status.ExitStatusException()
	}
}

func main() {
	inter := pygmy.PygmyNewPython()
	defer inter.FinalizeEx()

	EvalStatus(inter.PreInitializeFromBytesArgs(os.Args))
	EvalStatus(inter.SetArgv(os.Args))
	EvalStatus(inter.InitializeFromConfig())

	inter.RunMain()
	inter.FinalizeEx()
}
