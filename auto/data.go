package auto

import "go-blog-api/models"

var users = []models.User{
	{Nickname: "jhon Doe", Email: "jhon@email.com", Password: "123456789"},
}

var posts = []models.Post{
	{
		Title:   "Title",
		Content: "Hello World",
	},
}
