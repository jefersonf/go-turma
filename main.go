package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"turminha/turma"

	"golang.org/x/exp/slices"
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
		tmpl := template.Must(template.ParseFiles("layout/index.html", "layout/children-list.html"))
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		// log.Print(r.Header.Get("HX-Request"))
		sortingType := r.PostFormValue("sorting-type")
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
		tmpl := template.Must(template.ParseFiles("layout/children-list.html"))
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
		s := strings.Split(strings.Trim(turmaScanner.Text(), "\"'"), ",")

		fullname := strings.Trim(s[0], "\"'")
		name, _ := turma.GetFirstAndLastName(s[0])

		b, err := time.Parse("02/01/2006", strings.Trim(s[1], "\"'"))
		if err != nil {
			panic(err)
		}
		childAgeInMinutes := time.Since(b).Minutes()
		childAge := childAgeInMinutes / (60 * 24 * 365)
		birthDate := b.Format("Jan 02, 2006")

		var genre string
		s[3] = strings.Trim(s[3], "\"'")
		if strings.Compare(s[3], "M") == 0 {
			genre = "Boy"
		} else if strings.Compare(s[3], "F") == 0 {
			genre = "Girl"
		}

		t.AddChild(turma.Child{
			FullName:     fullname,
			Name:         name,
			BirthDate:    birthDate,
			Age:          int(childAge),
			AgeInMinutes: childAgeInMinutes,
			Genre:        genre,
		})
	}

	return t
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
