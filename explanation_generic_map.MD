```
package main

import "fmt"

// input : [1,2,3], (n) => n * 2
// output : [2,4,6]

func MapValues(values []int, mapFunc func (int)int) []int { // mapFunc is called parameter as argument
	var newValues []int // declare variable

	for _, v := range values { // looping
		newValue := mapFunc(v) // set newValue from result mapFunc. This function will calculate first and then assign to new value
		newValues = append(newValues, newValue) // append result of calculate to array
	}

	return newValues
}
func main(){
	result := MapValues([]int{1,2,3}, func(i int) int {
		return i * 2
	})

	fmt.Println(result)
}
```

The problem is we can only pass the array of integer. To solve this problem, we can use library constraints. Package constraints defines a set of useful constraints to be used with type parameters.

`"golang.org/x/exp/constraints"`

install this package and then use in function

```
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
```
