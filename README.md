# Get Fibonachi Slice (REST-API)
The REST API service on Golang, which has one method-which returns a slice from the Fibonacci sequence from the indices `x` to `y`. 
Code have tests.

## Input parameters
`x`, `y` - the integer ordinal indices of the Fibonacci sequence

## Output
The output should list all the numbers, the Fibonacci sequences with ordinal numbers from `x` to `y`    
`{"x":X_PARAM,"y":Y_PARAM,"answer":[*The slice of Fibonacci sequence*]}`

## Limitations
Conditions under which the service will give the correct answer is:    
`y >= x` and `x, y > 0` and `x, y <= 92`    
The restriction of 92 is due to the overflow of a variable of type int

## Examples
For example, for `x=0` and `y=5` service will respond `{"x":0,"y":5,"answer":[0,1,1,2,3,5]}`    
URL for this example is http://127.0.0.1:8080/fibonachi?x=0&y=5

## Quick start
+ Install Go (if you haven't already): https://golang.org/doc/tutorial/getting-started
+ In project directory run in terminal: `go run main.go`
+ Ready! Try to send URL in your browser http://127.0.0.1:8080/fibonachi?x=X_PARAM&y=Y_PARAM, where `X_PARAM` and `Y_PARAM` correspond to the "Limitations" section
+ (Additionally): run tests in project directory with command `go test -v` in terminal

