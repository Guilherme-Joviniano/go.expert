package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int32
}

func main() {
	course := Course{"Go", 40}
	courseTemplate := template.New("CourseTemplate")
	
	courseTemplate , _ = courseTemplate.Parse("Curso: {{.Name}} - Carga Horária: {{.Workload}}")
	
	err := courseTemplate.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
