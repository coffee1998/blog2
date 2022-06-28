package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//fmt.Println("blog server start...")
	//time.Sleep(30 * time.Second)
	//fmt.Println("blog server end...")
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":9999", nil))
	fmt.Println("start...")
	getInstanceId()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("end")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func getInstanceId() {
	url := "http://100.100.100.200/latest/meta-data/instance-id"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bs))
		}
	}
}
