package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/ChimeraCoder/anaconda"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world from %s", runtime.Version())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", html.EscapeString(r.URL.Path)[1:])
}

func bindListenServe() {
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	anaconda.SetConsumerKey(CONSUMERKEY)
	anaconda.SetConsumerSecret(CONSUMERSECRET)
	api := anaconda.NewTwitterApi(ACCESSTOKEN, ACCESSTOKENSECRET)

	searchText := "Boston"
	searchResult, err := api.GetSearch(searchText, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Recent tweets that contain #%s:\n", searchText)
	for _, tweet := range searchResult.Statuses {
		if tweet.Text != "" {
			fmt.Fprint(w, tweet.Text)
		}
	}
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/", hello)
	//http.HandleFunc("/", echo)
	//http.HandleFunc("/tweet", getTweet)
	bindListenServe()
}
