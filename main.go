package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    helloHandler := func(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "Hello, world\n")
    }
    
    postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "Posting Article...\n")
    }

    listArticleHandler := func(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "Article List\n")
    }
    
    articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
        articleID := 1
        resString := fmt.Sprintf("Article No.%d\n", articleID)
        io.WriteString(w, resString)
    }

    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/article", postArticleHandler)
    http.HandleFunc("/article/list", listArticleHandler)
    http.HandleFunc("/article/1", articleDetailHandler)

    log.Println("server start at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
