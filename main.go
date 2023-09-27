package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"turminha/turma"
	"turminha/util"

	"golang.org/x/exp/slices"
)

var (
	listenAddr = "localhost:3000"
	assetsPath = "static/"
)

func main() {
	data := util.ParseData()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(assetsPath))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		orderBy := r.PostFormValue("sortingtype")
		if len(orderBy) == 0 {
			orderBy = "name"
		}
		slices.SortFunc(data.ChildrenList, func(c1 turma.Child, c2 turma.Child) int {
			if orderBy == "name" {
				return strings.Compare(c1.Name, c2.Name)
			}
			order := 1
			if orderBy != "youngestFirst" {
				order *= -1
			}
			if int(c1.AgeInMinutes) < int(c2.AgeInMinutes) {
				return -order
			} else if int(c1.AgeInMinutes) > int(c2.AgeInMinutes) {
				return order
			}
			return 0
		})

		tmpl := template.Must(template.ParseFiles("layout/index.html", "layout/children-list.html"))
		tmpl.Execute(w, data)
	})

	log.Println("server running at", listenAddr)
	panic(http.ListenAndServe(listenAddr, nil))
}
