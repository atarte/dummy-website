package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var (
	AppName string = "dummy-website"
	Version string = "0.0.1"
	// Build   string = "abcd"
	Address string = "127.0.0.1"
	Port    string = "8888"
)

func main() {
	r := mux.NewRouter()

	staticDir := "/assets/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	r.HandleFunc("/", MainHandler)

	srv_adr := Address + ":" + Port

	srv := &http.Server{
		Handler: r,
		Addr:    srv_adr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Lauching server on : " + srv_adr)

	log.Fatal(srv.ListenAndServe())

}

type MainArgs struct {
	AppName string
	Address string
	Port    string
	Version string
}

func CombineAddress(address, port string) string {
	return fmt.Sprintf("%s:%s", address, port)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Handler")

	h := MainArgs{
		AppName: AppName,
		Address: Address,
		Port:    Port,
		Version: Version,
	}

	f := template.FuncMap{
		"CombineAddress": CombineAddress,
	}

	t := template.New("Main template").Funcs(f)

	t = template.Must(t.ParseFiles(
		"assets/templates/layout.tmpl",
		"assets/templates/main.tmpl",
	))

	err := t.ExecuteTemplate(w, "layout", h)
	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
