package main

import (
	"fmt"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

func main() {

	gocron.Every(1).Second().Do(innerFunc("user"))

	gocron.Clear()

	<-gocron.Start()

}

func innerFunc(name string) string {
	return "Hello, " + name + "!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello, world!")
}
