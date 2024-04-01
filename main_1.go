package main

import "fmt"

type UserID int

type Num interface{
	~int | int8 | int16 | int32 | float64 | float32
}
func Add[T Num](a T, b T) T {
	return a + b
}

func main(){
	a := UserID(10)
	b := UserID(200)
	result := Add(a,b)
	fmt.Println(result)
}