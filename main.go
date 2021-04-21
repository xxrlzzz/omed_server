package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"omed_server/editors"
	"omed_server/util"
	"strconv"
)

var appConf map[string]string
var codebaseConf map[string]string

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func init() {
	confFile, err := ioutil.ReadFile("omed-conf.yaml")
	if err != nil {
		log.Fatalf("fail to load conf.yaml with error %#v", err)
	}
	configMap := make(map[string]map[string]string)
	err = yaml.Unmarshal(confFile, configMap)
	if err != nil {
		log.Fatalf("fail to parse conf.yaml with error %#v", err)
	}
	appConf = configMap["app"]
	codebaseConf = configMap["codebase"]

	filePath := appConf["log-path"]
	logFile := util.TryOpen(filePath)
	if logFile != nil {
		log.SetOutput(logFile)
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	var detail editors.OpenEditorDetails

	detail.File = r.FormValue("file")
	detail.Project = r.FormValue("project")
	detail.Line = r.FormValue("line")
	detail.Branch = r.FormValue("branch")
	detail.Ide = r.FormValue("ide")
	if _, err := strconv.Atoi(detail.Line); err != nil {
		detail.Line = "1"
	}
	if detail.File == "" || detail.Project == "" {
		log.Printf("empty file/project")
		return
	}
	//fmt.Printf("%#v\n", detail)

	var editor editors.Editor
	 codebase := codebaseConf[detail.Project]
	switch detail.Ide {
	case "vscode":
		{
			editor = &editors.Vscode{Detail: detail}
			break
		}
	case "jetbrains":
		{
			ideName := r.FormValue("ide")
			editor = &editors.JetBrain{Detail: detail, IDEName: ideName}
			break
		}
	}

	go func() {
		if editor == nil {
			return
		}

		if codebase == "" {
			log.Printf("fail to loacte codebase %s", detail.Project)
			return
		}
		err := editor.Cmd(codebase)
		if err != nil {
			log.Printf("open vscode error:%#v", err)
		}
	}()
	_, _ = fmt.Fprintf(w, "ok")
}

func main() {

	http.HandleFunc("/file", Logging(FileHandler))

	log.Printf("omed server stary at %s", appConf["port"])
	_ = http.ListenAndServe(fmt.Sprintf(":%s", appConf["port"]), nil)
}
