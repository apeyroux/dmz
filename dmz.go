package main

import (
	"github.com/j4/dmz/rsa"
	"html/template"
	//	"io/ioutil"
	"log"
	"net/http"
)

func keyHandler(w http.ResponseWriter, r *http.Request) {

	type Page struct {
		KeyPem string
		PubPem string
		Ssh    string
	}

	log.Printf("Génération de clefs")
	pkPem, pubPem, pubSSHAK, err := rsa.GenerateKey()
	if err != nil {
		return
	}

	// ioutil.WriteFile("id_rsa.pub", pubSSHAK, 0600)
	// ioutil.WriteFile("id_rsa.pem", pubPem, 0600)
	// ioutil.WriteFile("id_rsa", pkPem, 0600)

	p := Page{string(pkPem), string(pubPem), string(pubSSHAK)}
	t, err := template.ParseFiles("tpl/index.html")
	if err != nil {
		w.Write([]byte("err tpl"))
	}
	t.Execute(w, p)

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/key/", keyHandler)
	http.ListenAndServe(":8080", nil)
}
