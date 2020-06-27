package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Content string
}

func main() {
	txtfile := flag.String("file", "first-post.txt", "The name of a text file in the current directory")
	flag.Parse()
	renderTemplate("template.tmpl", *txtfile)
}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(templateFile, txtfile string) {
	tmplData := data{Content: readFile(txtfile)}

	t := template.Must(template.New(templateFile).ParseFiles(templateFile))
	newTemplate := strings.Split(txtfile, ".")[0] + ".html"
	file, _ := os.Create(newTemplate)

	err := t.Execute(file, tmplData)

	if err != nil {
		panic(err)
	}
}
