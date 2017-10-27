package error

import (
	"fmt"
	"log"
	"os"
)

func first() {
	fmt.Println("first function executed")
}
func second() {
	fmt.Println("Defer 2nd")
}
func FigureOutDefer() {
	defer second()
	first()
	deferFileClose()
}

func deferFileClose() {
	f, _ := os.Open("./error/sample.txt")

	defer func() {
		err := f.Close()

		if err != nil {
			log.Fatal("Error when closing a file", err)
		}
	}()
}
