package main

import (
	"html/template"
	"os"

	"github.com/Guilherme-Joviniano/go-template-examples-and-usages/must"
	"github.com/Guilherme-Joviniano/go-template-examples-and-usages/templateWithExternalFile"
)

type Course struct {
	Name     string
	Workload int32
}

func BasicExampleOfTemplateUsage() {
	course := Course{"Go", 40}
	courseTemplate := template.New("CourseTemplate")

	courseTemplate, _ = courseTemplate.Parse("Curso: {{.Name}} - Carga Horária: {{.Workload}}")

	err := courseTemplate.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}

func main() {
	BasicExampleOfTemplateUsage()
	must.MustExampleOfUsage()
	templateWithExternalFile.TemplateWithExternalFileExampleOfUsage()
}
