package must

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int32
}

func MustExampleOfUsage() {
	course := Course{"Go", 40}
	courseTemplate := template.Must(template.New("CourseTemplate").Parse("Curso: {{.Name}} - Carga Horária: {{.Workload}}"))
	err := courseTemplate.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
