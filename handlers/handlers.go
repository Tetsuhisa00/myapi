
package handlers



import (
   "fmt"
   "io"
   "net/http"
)


// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodGet {
	     io.WriteString(w, "Hello, world!\n")
     } else {
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
     }
}


// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodPost {
	io.WriteString(w, "Posting Article...\n")
    } else {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}


// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodGet {
        io.WriteString(w, "Article List\n")
    } else {
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}


// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
    articleID := 1
    resString := fmt.Sprintf("Article No.%d\n", articleID)
    if req.Method == http.MethodGet {
	io.WriteString(w, resString)
    } else {
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}



// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodPost {
	io.WriteString(w, "Posting Nice...\n")
    } else {
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}


// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodPost {
        io.WriteString(w, "Posting Comment...\n")
    } else {
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
   }
}
