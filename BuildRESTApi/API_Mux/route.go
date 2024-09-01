package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = append(posts, Post{ID:1, Title: "Title 1", Text: "Text 1"})
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	post.ID = len(posts) + 1
	posts = append(posts, post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}