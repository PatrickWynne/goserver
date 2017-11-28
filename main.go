package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
)

type Page struct {
    Title string
    Body  []byte
}
type JsonObject struct {
    Id    int
    Name string
  }
  
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/html/", htmlHandler)
    http.HandleFunc("/jsonResult/", jsonResultHandler)
    http.ListenAndServe(":9090", nil)
}
//a very simple handler returns a string
func viewHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    urlString := strings.Replace(r.URL.Path[1:], "/", "", 2)
    fmt.Fprintf(w, "Hi there, I love %s! \n", urlString)
    fmt.Fprint(w, "<a href='/html/gopage'>Click here to see a new page</a>")
}
//handler returns an html page composed of the url path parameter and html content from file
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
//returns an html page if available using the url path parameter provided
func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
//return json
func jsonResultHandler(w http.ResponseWriter, r *http.Request) {
    jsonObject := JsonObject{1, "Programmer"}  
    js, err := json.Marshal(jsonObject)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }  
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
  }
