package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s", e.When, e.What)
}

func doThrowError() *MyError {
	return &MyError{
		time.Now(),
		"something went wrong",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cant be -ve %v", float64(e))
}

func main() {
	v := doThrowError()
	if v != nil {
		fmt.Printf("error %v", v)

	}

	fmt.Println(ErrNegativeSqrt(-2).Error())

}
