package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/Guilherme-Joviniano/go-template-examples-and-usages/must"
	"github.com/Guilherme-Joviniano/go-template-examples-and-usages/templateWithExternalFile"
)

type Course struct {
	Name     string
	Workload int32
}

type Courses []Course

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlTemplate := template.Must(template.New("template.html").ParseFiles("./templateWithExternalFile/template.html"))
		err := htmlTemplate.Execute(w, Courses{
			{"Go", 40},
			{"Java", 100},
			{"Javascript", 60},
			{"PHP", 80},
		})

		if err != nil {
			panic(err)
		}
	})

	BasicExampleOfTemplateUsage()
	must.MustExampleOfUsage()
	templateWithExternalFile.TemplateWithExternalFileExampleOfUsage()
	
	http.ListenAndServe(":8080", nil)
}
