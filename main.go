package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Bio struct {
	ID        int
	Nama      string
	Email	  string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	http.HandleFunc("/", templates)
	http.HandleFunc("/login", loginweb)
	http.HandleFunc("/logout", logoutweb)

	fmt.Println("Your server will be serve at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func templates(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func loginweb(w http.ResponseWriter, r*http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		var email = r.FormValue("email")

		emails := []string{"Stefanus@gmail.com", "christanto@yahoo.com", "diantono@ymail.com"}
		output := generatebiodata(emails)

		var argint int

		for i, x := range emails {
			if email == x {
				argint = i
			}
		}

		for i, x := range output {
			if argint == i {
				var data = map[string]string{"email": email, "message":"Welcome "+ x.Nama, "Alamat":x.Alamat, "Alasan":x.Alasan}
				if err := tmpl.Execute(w, data); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}

		return
}

http.Error(w, "", http.StatusBadRequest)
}

func logoutweb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("login.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
}

http.Error(w, "", http.StatusBadRequest)
}

func generatebiodata(bio []string)[]Bio{
	nama := []string{"Stefanus", "Christanto", "Diantono"}
	alamat := []string{"jalan", "road", "street"}
	alasan := []string{"reason", "cincong", "ngadi-ngadi"}

	generate := make([]Bio, 0)

	var p Bio
	for key, value := range bio {
		p.Nama = nama[key]
		p.Email = value
		p.Alamat = alamat[key]
		p.Alasan = alasan[key]
		generate = append(generate, p)
	}
	return generate
}
