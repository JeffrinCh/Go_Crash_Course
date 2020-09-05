package main

import (
	"fmt"
	"math"

	"github.com/cjeffrin/go_crash_course/03_packages/strutil"
)

func greetings(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
	fmt.Println(greetings("Jolly"))

}
