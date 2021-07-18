package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("hello world")

	// var s, sep string
	// start := time.Now()

	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// 	//fmt.Println(i, os.Args[i])
	// }

	// fmt.Println(s, time.Since(start))

	argSlc := []string{"arg1", "arg2", "arg3"}

	alternative(argSlc)
	alternative1(argSlc)
	alternative2(argSlc)

}

func alternative(slc []string) {
	var s, sep string
	start := time.Now()

	for i := 1; i < len(slc); i++ {
		s += sep + slc[i]
		sep = " "
	}

	fmt.Println("alterative:", s, time.Since(start))

}

func alternative1(slc []string) {
	s, sep := "", ""
	// var s, sep string
	// var s, sep = "", ""
	// var s, sep string = "", ""

	start1 := time.Now()

	for _, arg := range slc {
		s += sep + arg
		sep = " "
	}
	fmt.Println("alterative1:", s, time.Since(start1))
}

func alternative2(slc []string) {
	start2 := time.Now()
	fmt.Println("alterative2:", strings.Join(slc, " "), time.Since(start2))
}
