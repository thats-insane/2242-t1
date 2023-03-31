package main

import (
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func Greeting(w http.ResponseWriter, r *http.Request) {
	time = time.Now()
	day := time.Weekday().string()
	greeting := "Right now it is" + time + ". Enjoy the rest of your" + day + "."
	w.Write([]byte(greeting))
}

type Random struct {
	quote string
}

func main() {

}
