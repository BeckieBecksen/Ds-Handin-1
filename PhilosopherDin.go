package main

import "fmt"

func philo(philoId int, leftHand chan int, rightHand chan int) {

	isEating := false

	for {
		if !isEating {
			fmt.Println(philoId, "eating")
			<-rightHand
			<-leftHand
			isEating = true
		} else {
			isEating = false
			fmt.Println(philoId, "thinking")
			leftHand <- 33
			rightHand <- 33
		}

	}
}

func fork(forkId int, available chan int) {

	available <- 33
	for {

	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go philo(1, ch1, ch2)
	go philo(2, ch2, ch3)
	go philo(3, ch3, ch4)
	go philo(4, ch4, ch5)
	go philo(5, ch5, ch1)

	go fork(1, ch1)
	go fork(2, ch2)
	go fork(3, ch3)
	go fork(4, ch4)
	go fork(5, ch5)

	for {

	}

}

///////////////////////////////////////////////////////////////////7
