package main

import "fmt"

func philo(philoId int) {
	leftHand := make(chan int)
	rightHand := make(chan int)
	philoId = philoId
	isEating := false

	for {
		if !isEating {
			<-rightHand
			<-leftHand
			fmt.Println("eating")
			isEating = true
		} else {
			isEating = false
			fmt.Println("thinking")
			leftHand <- 33
			rightHand <- 33
		}
	}
}

func fork(forkId int) {
	isUsed := false

	available := make(chan int)
	available <- 33
	for {

	}
}

func main() {
	go philo(1)
	go philo(2)
	go philo(3)
	go philo(4)
	go philo(5)

	go fork(1)
	go fork(2)
	go fork(3)
	go fork(4)
	go fork(5)

	for {

	}

}

///////////////////////////////////////////////////////////////////7
