package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Howdy. My name is Cahlil and I like playing Yu-Gi-Oh! and I dislike women. My e-mail address is 2021154337@ub.edu.bz if you ever want to contact me."))
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Right now it is" + time.Now().Format(time.Kitchen) + ". Enjoy the rest of your" + time.Now().Weekday().String() + "."))
}

func Random(w http.ResponseWriter, r *http.Request) {
	quotes := []string{
		"I was so broke I was eating oxygen sandwiches. - Zachary Fox",
		"I don't care. - Barrington Hendricks",
		"What? - Jordan Carter",
		"Watch yo jet. - Casey Lawrence",
		"You don't exist. - Noah Smith",
		"I hate when GIRLS die. - Jeffrey Williams",
		"Damn, boy. - Anthony Fantano",
	}
	rand.Seed(time.Now().UnixNano())

	w.Write([]byte(quotes[rand.Intn(len(quotes))]))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/greeting", Greeting)
	mux.HandleFunc("/random", Random)
	log.Fatal(http.ListenAndServe(":4000", mux))
}
