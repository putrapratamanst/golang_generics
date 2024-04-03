package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// input : [1,2,3], (n) => n * 2
// output : [2,4,6]

func MapValues[T constraints.Ordered](values []T, mapFunc func (T)T) []T { // mapFunc is called parameter as argument
	var newValues []T // declare variable

	for _, v := range values { // looping
		newValue := mapFunc(v) // set newValue from result mapFunc. This function will calculate first and then assign to new value
		newValues = append(newValues, newValue) // append result of calculate to array
	}
	
	return newValues
}
func main(){
	result := MapValues([]float64{1.1,2.2,3.2}, func(i float64) float64 {
		return i * 2
	})

	fmt.Println(result)
}