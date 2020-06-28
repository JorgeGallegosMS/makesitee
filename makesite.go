package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Content string
}

func main() {
	txtfile := flag.String("file", "", "The name of a text file in the current directory")
	directory := flag.String("dir", "", "Finds all the text files in the current directory")
	flag.Parse()

	if *txtfile != "" {
		generateHTMLFile("template.tmpl", *txtfile)
	}

	if *directory != "" {
		files := findTxtFiles(*directory)

		for _, file := range files {
			if strings.Contains(file.Name(), ".txt") {
				generateHTMLFile("template.tmpl", file.Name())
			}
		}
	}
}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func generateHTMLFile(templateFile, txtfile string) {
	tmplData := data{Content: readFile(txtfile)}

	t := template.Must(template.New(templateFile).ParseFiles(templateFile))
	newTemplate := strings.Split(txtfile, ".")[0] + ".html"
	file, _ := os.Create(newTemplate)

	err := t.Execute(file, tmplData)

	if err != nil {
		panic(err)
	}
}

func findTxtFiles(directory string) []os.FileInfo {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
