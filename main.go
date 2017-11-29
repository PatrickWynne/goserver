package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "log"
	/*"net/url"*/
)

type Page struct {
    Title string
    Body  []byte
}
type JsonObject struct {
    Id    int
    Name string
}
type User struct{
    UserId int
    Title string
    Body string
}
  
func main() {
    http.HandleFunc("/text/", textHandler)
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/html/", htmlHandler)
    http.HandleFunc("/jsonResult/", jsonResultHandler)
    http.HandleFunc("/apiConsumer/", apiConsumerHandler)
    http.HandleFunc("/post", postHandler)
    http.ListenAndServe(":9090", nil)
}
func textHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}
//return render html + string
func viewHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    urlString := strings.Replace(r.URL.Path[1:], "/", "", 2)
    fmt.Fprintf(w, "Hi there, I love %s! \n", urlString)
    fmt.Fprint(w, "<a href='/html/gopage'>Click here to see a new page</a>")
}
//return render html
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
//return render [page].html
func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
//return json result
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
  //return external RESTful API json result [GET]
  func apiConsumerHandler(w http.ResponseWriter, r *http.Request){
    url := "https://jsonplaceholder.typicode.com/posts/3"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		log.Fatal("NewRequest: ", err)
		return
    }
    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
    defer resp.Body.Close()
    var record User
    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
    }
    w.Header().Set("Content-Type", "application/json")
    
    js, err := json.Marshal(record)
    w.Write(js)
  }
func postHandler(w http.ResponseWriter, r *http.Request){
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "Hello %s", name)
}