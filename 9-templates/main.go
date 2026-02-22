package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

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
}
