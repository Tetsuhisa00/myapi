package handlers

import (
	"io"
	"net/http"
	"strconv"
    "log"

	"encoding/json"

	"github.com/Tetsuhisa00/myapi/models"
	"github.com/gorilla/mux"
)

// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {

    io.WriteString(w, "Hello, world!\n")
}


// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

    var reqArticle models.Article
    
    if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
        http.Error(w, "fail to decode json\n", http.StatusBadRequest)
    }
    
    article := reqArticle
    json.NewEncoder(w).Encode(article)
}


// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
    queryMap := req.URL.Query()
    var page int
    if p, ok := queryMap["page"]; ok && len(p) > 0 {
        var err error
        page, err = strconv.Atoi(p[0])
        if err != nil {
            http.Error(w, "Invalid query parameter", http.StatusBadRequest)
            return 
        }
    } else {
        page = 1
    }
    log.Println(page)

    articleList := []models.Article{models.Article1, models.Article2}
    json.NewEncoder(w).Encode(articleList)
}


// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
    articleID , err := strconv.Atoi(mux.Vars(req)["id"])
    if err != nil {
        http.Error(w, "Invalid query parameter", http.StatusBadRequest)
        return 
    }
    log.Println(articleID)
    article := models.Article1
    json.NewEncoder(w).Encode(article)
}



// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
    article := models.Article1
    json.NewEncoder(w).Encode(article)
}


// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
    comment := models.Comment1
    json.NewEncoder(w).Encode(comment)
}
