package templateWithExternalFile

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int32
}

type Courses []Course

func TemplateWithExternalFileExampleOfUsage() {
	htmlTemplate := template.Must(template.New("template.html").ParseFiles("./templateWithExternalFile/template.html"))
	err := htmlTemplate.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 100},
		{"Javascript", 60},
		{"PHP", 80},
	})

	if err != nil {
		panic(err)
	}
}
