package main

import (
	"log"
	"os"
	"text/template"

	"github.com/occamshub-dev/grok"
)

func main() {
	g, _ := grok.New()
	addError := g.AddPatternsFromPath("./patterns")
	if addError != nil {
		log.Fatalln(addError)
	}
	tpl, _ := template.New("patterns").Parse("package grok\n\nvar patterns = map[string]string{\n" +
		"{{ range $key, $value := . }}" +
		"\t" + `"` + "{{ $key }}" + `": ` + "`" + "{{ $value }}" + "`,\n" +
		"{{ end }}" +
		"\n}")

	f, _ := os.Create("patterns.go")
	exeError := tpl.Execute(f, g.GetRawPatterns())
	if exeError != nil {
		log.Fatalln(exeError)
	}
}
