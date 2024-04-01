```
package main
import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	result := Add(1, 2)
	fmt.Println(result)
}
```

This code is working. But, what if I want to input the argument in Add as float or int?

```
package main
import "fmt"

func AddInt(a int, b int) int {
	return a + b
}
func AddFloat(a float64, b float64) float64 {
	return a + b
}
func main() {
	result := AddFloat(1.1, 2)
	fmt.Println(result)
}
```

---

Usually, I declare another variable to handle this case.

But we can make it simple and general. We will use generic (go version > v1.18)

```
package main
import "fmt"

func Add[T int | float64](a T, b T) T{
	return a + b
}

func main(){
	result := Add(10.1, 2)
	fmt.Println(result)
}
```



---

This was so cool. But the problem is there is a lot of data type.

The example: `func Add[T int | int8 | int16 | int32 | float64 | float32]`

We will use an interface to handle this argument.

```
package main

import "fmt"

type Num interface{
	int | int8 | int16 | int32 | float64 | float32
}
func Add[T Num](a T, b T) T {
	return a + b
}

func main(){
	result := Add(1,3.3)
	fmt.Println(result)
}
```
Cool!

It's very common in golang to use own type aliases.
Example.

```
package main

import "fmt"

type UserID int

type Num interface{
	int | int8 | int16 | int32 | float64 | float32
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
```
The result will be an error like this. ` UserID does not satisfy Num (possibly missing ~ for int in Num)`

So basically, this tilda `~` allow you to use any type that is an alias for underlying type

```
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
```