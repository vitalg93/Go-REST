package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	X      int   `json:"x"`
	Y      int   `json:"y"`
	Answer []int `json:"answer"`
}

var tasks []Task

func getFibonachi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// получение параметров и их конвертация в целое число (int)
	x, err := strconv.Atoi(params["x"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(params["y"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slice := getFibonachiSlice(x, y)
	fmt.Println(x, y, slice)

	task := Task{x, y, slice}
	//fmt.Println(task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Вычисление числа из последовательности Фибоначчи,
func fibonachi() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

// Возвращает срез последовательности чисел из ряда Фибоначчи от x до y
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

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/fibonachi/{x:[0-9]+}/{y:[0-9]+}", getFibonachi).Methods("GET")
	fmt.Println("Server started. Try to get Fibonachi slice (JSON) from x to y by URL http://127.0.0.1:8000/fibonachi/x/y. Valid answer by x >= y and x, y > 0 and x, y <= 92")
	log.Fatal(http.ListenAndServe(":8000", router))
}
