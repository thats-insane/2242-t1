package main

import (
	"log"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Howdy. My name is Cahlil and I like playing Yu-Gi-Oh! and I hate women. My e-mail address is 2021154337@ub.edu.bz if you ever want to contact me."))
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	greeting := "Right now it is" + time.Now().Format(time.Kitchen) + ". Enjoy the rest of your" + time.Now().Weekday().String() + "."

	w.Write([]byte(greeting))
}

func Random(w http.ResponseWriter, r *http.Request) {
}

func main() {
	log.Fatal(http.ListenAndServe(":4000", nil))
}
