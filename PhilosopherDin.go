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
			leftCanBeServed := <-leftIn
			if leftCanBeServed {
				rightOut <- philoId
				rightCanBeServed := <-rightIn
				if rightCanBeServed {
					fmt.Println(philoId, "eating for the", eatCount, "th time")
					isEating = true
					eatCount++
					time.Sleep(1 * time.Second)
					leftOut <- philoId
					rightOut <- philoId
				} else {
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

func fork(forkId int, leftOut, rightOut chan bool, sharedIn chan int, leftPhiloID, rightPhiloID int) {
	isAvalible := true
	philoUsingFork := 0

	for {
		philoID := <-sharedIn

		if philoID == leftPhiloID {
			if isAvalible {
				leftOut <- true
				philoUsingFork = philoID
				isAvalible = false
			} else {
				if philoUsingFork == philoID {
					isAvalible = true
					philoUsingFork = 0
				} else {
					leftOut <- false
				}
			}
		}

		if philoID == rightPhiloID {
			if isAvalible {
				rightOut <- true
				philoUsingFork = philoID
				isAvalible = false
			} else {
				if philoUsingFork == philoID {
					isAvalible = true
					philoUsingFork = 0
				} else {
					rightOut <- false
				}
			}

		}

		if philoID != rightPhiloID && philoID != leftPhiloID {
			fmt.Println("ERROR, on fork", forkId, "got id:", philoID, "expected one of:", leftPhiloID, rightPhiloID)
		}
	}
}

func main() {
	var comms = [10]chan bool{}
	for i := 0; i < 10; i++ {
		comms[i] = make(chan (bool), 1)
	}

	var forkIn = [5]chan int{}
	for i := 0; i < 5; i++ {
		forkIn[i] = make(chan (int), 1)
	}

	//forkId int, leftOut, rightOut chan bool, sharedIn chan int, leftPhiloID, rightPhiloID int
	go fork(1, comms[9], comms[0], forkIn[0], 5, 1)
	go fork(2, comms[1], comms[2], forkIn[1], 1, 2)
	go fork(3, comms[3], comms[4], forkIn[2], 2, 3)
	go fork(4, comms[5], comms[6], forkIn[3], 3, 4)
	go fork(5, comms[7], comms[8], forkIn[4], 4, 5)

	//philoId int, leftIn, rightIn chan bool, leftOut, rightOut chan int
	go philo(1, comms[0], comms[1], forkIn[0], forkIn[1])
	go philo(2, comms[2], comms[3], forkIn[1], forkIn[2])
	go philo(3, comms[4], comms[5], forkIn[2], forkIn[3])
	go philo(4, comms[6], comms[7], forkIn[3], forkIn[4])
	go philo(5, comms[8], comms[9], forkIn[4], forkIn[0])

	for {

	}
}
