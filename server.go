package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintln(w, "Hello Sasha!")
	fmt.Fprintln(w, "Try new features: Current Date && Delay!")
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/date", CurrDate)
	http.HandleFunc("/wait", Delay)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Delay(w http.ResponseWriter, r *http.Request){
	value := r.URL.Query().Get("delay")
	if value != ""{
		SleepTime, err := strconv.Atoi(value)
		if err ==nil {
			time.Sleep(time.Duration(SleepTime) * time.Millisecond)
			fmt.Fprintln(w, "Sleep for " +value+ "ms completed")
		}else{
			fmt.Fprint(w,"Error. Write another duration of delay.")
		}
	}
}


func CurrDate(w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Current date: %s", time.Now())
}