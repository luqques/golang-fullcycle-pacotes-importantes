package main

import (
	"net/http"
	"os"
	"strings"
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
	tmpl := template.Must(template.ParseFiles("content.html"))
	err = tmpl.Execute(os.Stdout, cursos)

	//com webserver
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles("content.html"))
		err := t.Execute(w, cursos)
		if err != nil {
			panic(err)
		}
	})

	//com vários arquivos
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	template2 := template.New("content.html")
	template2.Funcs((template.FuncMap{"ToUpper": strings.ToUpper}))
	template2 = template.Must(template2.ParseFiles(templates...))
	err = template2.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
