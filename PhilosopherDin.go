package main

import (
	"fmt"
	"time"
)

func philo(philoId int, leftIn, rightIn chan bool, leftOut, rightOut chan int) {
	isEating := false
	eatCount := 1
	for {
		if !isEating {
			leftOut <- philoId
			leftCanBeServed := <- leftIn
			if leftCanBeServed{
				rightOut <- philoId
				rightCanBeServed := <- rightIn
				if rightCanBeServed{
					fmt.Println(philoId, "eating for the", eatCount, "th time" )
					isEating = true
					eatCount++
					time.Sleep(0* time.Second)
					leftOut <- philoId
					rightOut <- philoId
				}else{
					// tell the other fork that it is free again
					leftOut <- philoId
				}
			}
		} else {
			isEating = false
			fmt.Println(philoId, "thinking")
		}
	}
}

func fork(forkId int, leftOut, rightOut chan bool, sharedIn chan int) {
	philoIDKey := 0
	for {
		philoIDKey = <- sharedIn
		if philoIDKey == forkId{
			rightOut <- true
		}else{
			leftOut <- true
		}
		
		for {
			passID := <- sharedIn
			if passID == philoIDKey{
				philoIDKey = 0
				break
			}else{
				if passID == forkId{
					leftOut <- false
				}else{
					rightOut <- false
				}
			}
		}	
	}
}

func main() {
	var comms = [10]chan bool{}
	for i := 0; i < 10; i++ {
		comms[i] = make(chan bool)
	}
	
	var forkIn = [5]chan int{}
	for i := 0; i < 5; i++ {
		forkIn[i] = make(chan int)
	}
	
	index := 0
	for i := 0; i < 5; i++{
		go philo(i+1, comms[index], comms[index+1],forkIn[(index-1+5)%5], forkIn[(index-2+5)%5] )
		go fork(i+1, comms[(index-1+10)%10], comms[index], forkIn[i])
		index++
		index++
	}

	for {

	}
}