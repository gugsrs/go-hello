package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
var templates = template.Must(template.ParseFiles("edit.html", "view.html", "search.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|delete|search)/([a-zA-Z0-9]+)$")

type Page struct {
	Title    string
	Body     []byte
	Contacts []string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {
	filename := "contatos.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	text := string(p.Body)

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
	return err
}

func loadPage(title string) (*Page, error) {
	filename := "contatos.txt"
	body, _ := ioutil.ReadFile(filename)
	contacts := strings.Split(string(body), "\n")
	return &Page{Title: title, Body: body, Contacts: contacts}, nil
}

/*
 *	Data delegate
 */

func deleteEntry(position int, p *Page) error {
	filename := "contatos.txt"

	var newText = ""
	for index, contacts := range p.Contacts {
		if index != position {
			newText = newText + contacts + "\n"
		}
	}
	return ioutil.WriteFile(filename, []byte(newText), 0600)
}

func filterPageData(text string, p *Page) {
	var filteredContacts string
	for _, contacts := range p.Contacts {
		if strings.Contains(strings.ToLower(contacts), strings.ToLower(text)) {
			filteredContacts = filteredContacts + contacts + "\n"
		}
	}
	p.Contacts = strings.Split(filteredContacts, "\n")
}

/*
 * Handlers
 */
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	name := r.FormValue("name")
	number := r.FormValue("number")
	email := r.FormValue("email")
	if name == "" || number == "" || email == "" {

	} else {
		p := &Page{Title: title, Body: []byte(name + " - " + number + " - " + email + "\n")}
		err := p.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/view/"+"Lista", http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		return
	}
	strings.Replace(title, "?", "", 1)

	i, _ := strconv.ParseInt(title, 0, 64)
	deleteEntry(int(i), p)
	http.Redirect(w, r, "/view/"+"Lista", http.StatusFound)
}

func searchHandler(w http.ResponseWriter, r *http.Request, title string) {
	searchText := r.FormValue("contactName")
	p, err := loadPage(title)
	if err != nil {
		return
	}
	filterPageData(searchText, p)
	renderTemplate(w, "search", p)
}

func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.HandleFunc("/search/", makeHandler(searchHandler))

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}
	http.ListenAndServe(":8080", nil)
}
