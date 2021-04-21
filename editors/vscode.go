package editors

import (
	"fmt"
	"os/exec"
)

type Vscode struct {
	Detail OpenEditorDetails
}

func (v *Vscode) Cmd(codebase string) error {
	//fmt.Printf("cmd: code -g %s:%s in %s\n", v.Detail.File, v.Detail.Line, codebase)

	cmd := exec.Command("code", "-g", fmt.Sprintf("%s:%s",v.Detail.File, v.Detail.Line))
	cmd.Dir = codebase
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
