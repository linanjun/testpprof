package world

import "fmt"

func Add(value string) string {
	result := fmt.Sprint("a+", value)
	return result
}
