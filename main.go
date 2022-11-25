package main

import (
	"github.com/bluesentinelsec/OffensiveGoLang/pkg/windows/execution"
)

func main() {

	execution.RunPowerShell("notepad.exe")
}
