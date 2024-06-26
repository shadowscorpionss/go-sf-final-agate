// Пакет для работы с моделями приложения
package models

//full detailed posts
type NewsFullDetailed struct {
	Id      int
	Title   string
	PubTime int64  // publication time
	Link    string // reference to original
	Content string // post content
}

//short detailed
type NewsShortDetailed struct {
	Id      int
	Title   string
	PubTime int64  // publication time
	Link    string // reference to original
}

//comments model
type Comment struct {
	Id         int
	PostId     int    //Post that comment is connected to
	ParentId   int    //comment parent
	Content    string //comment itself
	AuthorName string //who send comment
	PubTime    int64
}

type NewComment struct{
	PostId int
	ParentId int
	Content string
	AuthorName string
}
