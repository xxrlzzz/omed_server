package editors

type Editor interface {
	Cmd(string) error
}


type OpenEditorDetails struct {
	File    string `json:"file"`
	Line    string `json:"line"`
	Ide    string `json:"ide"`
	Branch  string `json:"branch"`
	Project string `json:"project"`
}

