
package handlers



import (
   "fmt"
   "io"
   "net/http"
   "strconv"
   "github.com/gorilla/mux"
)


// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {

    io.WriteString(w, "Hello, world!\n")
}


// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
    
    io.WriteString(w, "Posting Article...\n")
}


// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

    io.WriteString(w, "Article List\n")
}


// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
    articleID , err := strconv.Atoi(mux.Vars(req)["id"])
    if err != nil {
        http.Error(w, "Invalid query parameter", http.StatusBadRequest)
        return 
    }
    resString := fmt.Sprintf("Article No.%d\n", articleID)
    io.WriteString(w, resString)
}



// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {

    io.WriteString(w, "Posting Nice...\n")
}


// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
    
    io.WriteString(w, "Posting Comment...\n")
}
