package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	/*
		gocron.Every(1).Second().Do(innerFunc("user"))
		gocron.Clear()
		<-gocron.Start()
	*/

	var numbers = randomArray(10, 10)
	fmt.Println(len(numbers))

	var x = string(6)
	var y json.RawMessage
	y, _ = json.Marshal(x)
	fmt.Println(string(y))

	serialize(1, numbers)

}

func innerFunc(name string) string {
	return "Hello, " + name + "!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello, world!")
}

func randomNums(num int) (int, int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num), rand.Intn(num)
}

func randomArray(size int, top int) []int {
	rand.Seed(time.Now().Unix())
	var result = make([]int, size)
	for i := 0; i < 10; i++ {
		result[i] = rand.Intn(top)
	}
	return result
}

func serialize(number int, numbers []int) {
	type Data struct {
		id   int
		nums []int
	}

	entity := Data{id: number, nums: numbers}

	b, error := json.Marshal(entity)
	if error != nil {
		fmt.Println("Error: ", error)
	}
	fmt.Println(b)
}
