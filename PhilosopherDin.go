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
	for {}
}

func main() {
	var channelArray = [5]chan int{}
	for i := 0; i<5; i++{
		channelArray[i] = make(chan int)
	}
	
	for i := 0; i<5; i++{
		go philo(i+1, channelArray[i], channelArray[(i+1)%5])	
	}
	
	for i := 0; i<5; i++{
		go fork(i+1, channelArray[i])
	}

	for {

	}
}