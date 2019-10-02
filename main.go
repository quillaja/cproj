package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	var proj Params

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Project name: ")
	input.Scan()
	proj.Project = input.Text()

	fmt.Print("Libraries (eg 'pthread' in '-lpthread'): ")
	input.Scan()
	proj.Libraries = clean(strings.Split(input.Text(), " "))

	fmt.Print("Library search paths: ")
	input.Scan()
	proj.LibDirs = clean(strings.Split(input.Text(), " "))

	fmt.Print("Include search paths: ")
	input.Scan()
	proj.IncDirs = clean(strings.Split(input.Text(), " "))

	// fmt.Printf("%#v\n", proj)

	os.Mkdir("src", 0755)
	os.Mkdir("build", 0755)
	os.Mkdir("lib", 0755)
	os.Mkdir("include", 0755)

	makeT := template.Must(template.New("makefile").Parse(makefileTemplate))
	mainT := template.Must(template.New("main").Parse(maincppTemplate))

	makeF, _ := os.Create("Makefile")
	mainF, _ := os.Create("src/main.cpp")

	makeT.Execute(makeF, proj)
	mainT.Execute(mainF, nil)
}

type Params struct {
	Project   string
	Libraries []string
	LibDirs   []string
	IncDirs   []string
}

func clean(in []string) (out []string) {
	out = make([]string, 0)
	for i := range in {
		if in[i] != "" {
			out = append(out, strings.TrimSpace(in[i]))
		}
	}
	return
}
