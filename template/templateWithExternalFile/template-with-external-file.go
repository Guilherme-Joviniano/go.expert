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
	templates := []string{
		"./templateWithExternalFile/header.html",
		"./templateWithExternalFile/content.html",
		"./templateWithExternalFile/footer.html",
	}

	htmlTemplate := template.Must(template.New("content.html").ParseFiles(templates...))
	
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
