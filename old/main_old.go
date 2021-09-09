// Version of code, where URL looks like localhost:8080/fibonachi/x/y - not like localhost:8080/fibonachi?x=0&y=5
package main_old

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

// Проверяет параметры запроса и выводит срез
func getFibonachi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	fmt.Println(r)

	// получение параметров и их конвертация в целое число (int)
	if params["x"] == "" {
		http.Error(w, "missing value x", http.StatusBadRequest)
		return
	}

	if params["y"] == "" {
		http.Error(w, "missing value y", http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(params["x"])
	if err != nil {
		http.Error(w, "can't convert x into int", http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(params["y"])
	if err != nil {
		http.Error(w, "can't convert y into int", http.StatusBadRequest)
		return
	}

	slice := getFibonachiSlice(x, y)
	//fmt.Println(x, y, slice)

	task := Task{x, y, slice}
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func myCustom404Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>To get Fibonacci slice (JSON)</h1> from x to y (integer numbers) go to URL <a href='http://127.0.0.1:8000/fibonachi/0/8'>http://127.0.0.1:8000/fibonachi/x/y</a><br> (for example - the URL show result at x=0, y=8). <br>Conditions under which the service will give the correct answer: y >= x and x, y > 0 and x, y <= 92")
}

func main_old() {
	router := mux.NewRouter()
	http.Handle("/", router)
	// обработка 404 ошибки
	router.NotFoundHandler = http.HandlerFunc(myCustom404Handler)
	router.HandleFunc("/fibonachi/{x:[0-9]+}/{y:[0-9]+}", getFibonachi).Methods("GET")
	fmt.Println("Server started. Try to get Fibonachi slice (JSON) from x to y (integer numbers) in URL http://127.0.0.1:8000/fibonachi/x/y. Conditions under which the service will give the correct answer: y >= x and x, y > 0 and x, y <= 92")
	log.Fatal(http.ListenAndServe(":8000", router))
}
