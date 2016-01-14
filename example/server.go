package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func idxHdlr(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/a.html")
	ce(err)
	t.ExecuteTemplate(w, "idx", nil)
}

func extHdlr(w http.ResponseWriter, r *http.Request) {
	os.Exit(0);
}

func getHdlr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := r.FormValue("code")
	dateout, err := exec.Command("/bin/date", "+%H-%m-%S").Output()
	ce(err)
	f, err := os.Create(strings.TrimSpace(string(dateout))+".lua")
	ce(err)
	defer f.Close()
	f.Write([]byte(strings.Replace(strings.TrimSpace(code), "\r\n", "\n", -1)))
	codeout, _ := exec.Command("/usr/local/belka/belka", strings.TrimSpace(string(dateout))+".lua").Output()
	os.Remove(strings.TrimSpace(string(dateout))+".lua")
	fmt.Fprintf(w, string(codeout))
}

func main() {
	http.HandleFunc("/", idxHdlr)
	http.HandleFunc("/exit", extHdlr)
	http.HandleFunc("/get", getHdlr)
	http.Handle("/data/", http.FileServer(http.Dir("./")))
	fmt.Println("BelkaServer: start on 8080...")
	err:=http.ListenAndServe(":8080", nil);
	ce(err)
}

func ce(e error) {
	if e != nil {
		log.Fatalf("BelkaServer:  %s", e.Error());
	}
}
