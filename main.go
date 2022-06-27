package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//fmt.Println("blog server start...")
	//time.Sleep(30 * time.Second)
	//fmt.Println("blog server end...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
	//fmt.Println("start...")
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//	time.Sleep(time.Second * 1)
	//}
	//fmt.Println("end")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
