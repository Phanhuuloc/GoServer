package main

import (
	"fmt"
	"net/http"

	"github.com/russross/blackfriday"
	"github.com/julienschmidt/httprouter"

	//"os"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//http.HandleFunc("/markdown", GenerateMarkdown)
	//http.Handle("/", http.FileServer(http.Dir("public")))
	//http.ListenAndServe(":"+port, nil)

	r := httprouter.New()
	//r.GET("/", HomeHandler)
	r.Handler("GET", "/", http.FileServer(http.Dir("public")))
	r.HandlerFunc("POST", "/markdown", GenerateMarkdown)
	// Posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":" + port, r)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}