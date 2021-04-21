package editors

import (
	"fmt"
	"os/exec"
	"strings"
)

var validEditor = [3]string{"goland", "idea", "clion"}

type JetBrain struct {
	Detail OpenEditorDetails
	IDEName string
}

func (j *JetBrain) validIDE() bool {
	ide := strings.ToLower(j.IDEName)
	for _, s := range validEditor {
		if ide == s {
			return true
		}
	}
	return false
}

func (j *JetBrain) Cmd(codebase string) error {
	if !j.validIDE() {
		return fmt.Errorf("invalid ide name %s",j.IDEName)
	}

	//fmt.Printf("cmd: %s %s:%s in %s\n", j.IDEName, j.Detail.File, j.Detail.Line, codebase)

	cmd := exec.Command(j.IDEName, fmt.Sprintf("%s:%s",j.Detail.File, j.Detail.Line))
	cmd.Dir = codebase
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}