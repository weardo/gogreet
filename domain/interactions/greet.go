package interactions

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Curse(name string) string {
	return fmt.Sprintf("Go to hell, %s!", name)
}