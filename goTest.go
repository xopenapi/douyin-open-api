package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "这是一个web服务器")
}

func main() {
	//http.HandleFunc("/douyin/login", sayhelloName)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
	//http.HandleFunc("/douyin/login/redirect1", sayhelloName)
	//http.ListenAndServe(":8889", nil)

	http.HandleFunc("/douyin/login", func(w http.ResponseWriter, r *http.Request) {
		clientKey := "awpamlod9sp7tgqv"
		redirectUrl := "http://proxy-admin.dev.lucfish.com:7898/douyin/login/redirect"
		response_type := "code"
		scope := "user_info"
		//url := fmt.Sprintf("https://open.douyin.com/platform/oauth/connect/?client_key=%?&response_type=%s?scope=%?&redirect_uri=%s",
		url := fmt.Sprintf("https://open.douyin.com/oauth/connect/?client_key=%?&response_type=%s?scope=%?&redirect_uri=%s",
			clientKey,
			response_type,
			scope,
			redirectUrl,
		)
		log.Print(url)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/douyin/login/redirect", func(w http.ResponseWriter, r *http.Request) {
		log.Print("---------------------------------------------redirect")
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Fprintf(w, "这是一个web服务器")
	})
	http.ListenAndServe(":8888", nil)

}
