package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

	//com must
	curso2 := Curso{"Python", 30}
	tmp2 := template.Must(template.New("CursoTemplate2").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err = tmp2.Execute(os.Stdout, curso2)

	//com arquivos
	cursos := Cursos{
		{"Go", 40},
		{"Python", 30},
		{"Java", 50},
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	err = tmpl.Execute(os.Stdout, cursos)
}
