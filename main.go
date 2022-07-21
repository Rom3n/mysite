package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is %s. He is %d years old with %d bucks", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {

	bob := User{"Bob", 25, 50, 4.2, 0.8, []string{"eblya", "greblya", "rybalka"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
	//bob := User{"Bob", 25, 50, 4.2, 0.8}
	//bob.setNewName("Not Bob")
	//fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleRequest()
}
