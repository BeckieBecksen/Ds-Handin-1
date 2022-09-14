package main

import (
	"fmt"
	"time"
)

func philo(philoId int, leftHand, rightHand, requestLeftHand, requestRightHand chan int, eatingTime int) {
	isEating := false
	eatCount := 1
	for {
		if !isEating {
			fmt.Println(philoId, "eating")
			requestLeftHand <- 22
			requestRightHand <- 22
			<-rightHand
			<-leftHand
			fmt.Println(philoId, "eating for the", eatCount, "th time" )
			isEating = true
			eatCount++
			time.Sleep(time.Duration(eatingTime/2) * time.Second)
		} else {
			isEating = false
			fmt.Println(philoId, "thinking")
			leftHand <- 33
			rightHand <- 33
		}
	}
}

func fork(forkId int, available, request chan int) {
	available <- 33
	inUse := false
	for {
		if inUse{
			<- available
			inUse = false
		}else{
			<- request
			inUse = true
			available <- 33
		}
	}
}

func main() {
	var channelArray = [5]chan int{}
	for i := 0; i < 5; i++ {
		channelArray[i] = make(chan int)
	}
	
	var eatingTimes = [5]int{2,3,5,7,11}
	
	var requestChannels = [5]chan int{}
	for i := 0; i < 5; i++ {
		requestChannels[i] = make(chan int)
	}

	for i := 0; i < 5; i++ {
		go fork(i+1, channelArray[i], requestChannels[i])
	}

	for i := 0; i < 5; i++ {
		go philo(i+1, channelArray[i], channelArray[(i+1)%5], requestChannels[i], requestChannels[(i+1)%5], eatingTimes[i])
	}
	

	for {

	}
}