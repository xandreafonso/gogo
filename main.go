package main

import (
	"fmt"
	"time"
)

// func processando() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	// go processando()
	// go processando()
	// processando()

	canal := make(chan int)

	go func() {
		// canal <- 1 // preencher o canal

		for i := 0; i < 10; i++ {
			canal <- i
			fmt.Println("A jogou i no canal", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			canal <- i
			fmt.Println("B jogou i no canal", i)
		}
	}()

	// go fmt.Println(<-canal) // esvaziar o canal
	// time.Sleep(time.Second * 2)

	// for x := range canal { // loop infinito que vai ficar lendo o canal
	// 	fmt.Println(x)
	// 	fmt.Println("Recebeu i do canal", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(canal, 1)
	worker(canal, 2)

}

func worker(canal chan int, workerId int) {
	for {
		fmt.Println("Recebeu do canal", <-canal, " no worker ", workerId)
		time.Sleep(time.Second)
	}
}
