# Get Fibonachi Slice
The REST API service on Golang, which has one method-which returns a slice from the Fibonacci sequence from the indices `x` to `y`

## Input parameters
`x`, `y` - the integer ordinal numbers of the fibonacci sequence

## Output
The output should list all the numbers, the Fibonacci sequences with ordinal numbers from `x` to `y`.
```json
{"x":X_NUM,"y":Y_NUM,"answer":[*The slice of Fibonacci sequence*]}
```

## Limitations
Conditions under which the service will give the correct answer is:    
`y >= x` and `x, y > 0` and `x, y <= 92`

## Examples
For example, for `x=0` and `y=5` service will respond `{"x":0,"y":5,"answer":[0,1,1,2,3,5]}`

## Quick start
