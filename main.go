package main

import (
	"log"
	"net/http"
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
		tmpl := template.Must(template.ParseFiles("layout/index.html", "layout/children-list.html"))
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		// log.Print(r.Header.Get("HX-Request"))
		sortingType := r.PostFormValue("sorting-type")
		log.Printf("sorting type = %v", sortingType)
		order := 1
		if sortingType == "fy" {
			order *= -1
		}
		slices.SortFunc(data.ChildrenList, func(c1 turma.Child, c2 turma.Child) int {
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
