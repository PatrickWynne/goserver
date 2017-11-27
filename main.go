package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/html/", htmlHandler)
    http.ListenAndServe(":9090", nil)
}
//a very simple handler returns a string
func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
//handler returns an html page composed of the url path parameter and html content from file
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
//compiles the actual page to be returned
func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
