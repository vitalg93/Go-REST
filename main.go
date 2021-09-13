package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	X      int   `json:"x"`
	Y      int   `json:"y"`
	Answer []int `json:"answer"`
}

// TODO: Save resolved tasks to cache
//var tasks []Task

// Calculating a number from the Fibonacci sequence
func fibonachi() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

// Returns a slice of a sequence of numbers from the Fibonacci series from x to y
func getFibonachiSlice(x, y int) []int {
	f := fibonachi()
	var result []int
	for i := 0; i <= y; i++ {
		value := f()
		if i >= x && i <= y {
			result = append(result, value)
		}
	}
	return result
}

// Checks the request parameters and outputs a slice
func getFibonachi(w http.ResponseWriter, r *http.Request) {
	textX := r.FormValue("x")
	textY := r.FormValue("y")

	if textX == "" {
		http.Error(w, "missing value x", http.StatusBadRequest)
		return
	}

	if textY == "" {
		http.Error(w, "missing value y", http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(textX)
	if err != nil {
		http.Error(w, "can't convert x into int", http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(textY)
	if err != nil {
		http.Error(w, "can't convert y into int", http.StatusBadRequest)
		return
	}

	slice := getFibonachiSlice(x, y)
	task := Task{x, y, slice}

	// TODO: for saving in cache
	//tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Route handler
func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/fibonachi", getFibonachi)
	return r
}

func main() {
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}
