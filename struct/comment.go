
package main

import "time"

type Comment struct {
	CommentID int
	ArticleID int
	Message string
	CreatedAt time.Time
}