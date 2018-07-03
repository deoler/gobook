// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Кол-во аргументов: ", len(os.Args)-1)
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i,": ",os.Args[i])
	}
}

//!-
