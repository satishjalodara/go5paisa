package go5paisa

import "fmt"

type InvalidOrderError struct {
	Err error
}

func (e *InvalidOrderError) Error() string {
	return fmt.Sprintf("Invalid order %v", e.Err)
}
