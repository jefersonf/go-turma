package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"turminha/turma"
	"turminha/types"
)

var (
	listenAddr = "localhost:3000"
	assetsPath = "static/"
)

func main() {
	data := parseData()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(assetsPath))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("layout/index.html"))
		tmpl.Execute(w, data)
	})

	log.Println("server running at", listenAddr)
	panic(http.ListenAndServe(listenAddr, nil))
}

func parseData() *turma.Turma {
	turmaFile, err := os.Open("data/turma.csv")
	checkErr(err)

	turmaScanner := bufio.NewScanner(turmaFile)
	turmaScanner.Split(bufio.ScanLines)

	// Skip first since it is a header
	turmaScanner.Scan()

	t := turma.New()

	for turmaScanner.Scan() {
		s := strings.Split(strings.Trim(turmaScanner.Text(), "\""), ",")
		t.AddChild(types.Child{
			FullName:  s[0],
			BirthDate: s[1],
			Genre:     s[3],
		})
	}

	return t
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
