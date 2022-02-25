package main

// Try,Except?
import (
	"errors"
	"fmt"
	"math"
)

func main() {
	f := 5.55
	var result float64
	var err error
	result, err = Sqrt(f)

	if err == nil {
		fmt.Println("结果是", result)
	} else {
		println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("开根号不能是负数")
	} else {
		return math.Sqrt(f), nil
	}
}
