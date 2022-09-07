package main

import "fmt"

type philosopher struct {
	leftHand  chan int
	rightHand chan int
	philoId   int
	isEating  bool
}

func buildP(Id int, forkarr []fork) philosopher {
	var p1 philosopher
	p1.philoId = Id
	if Id == 4 {
		p1.leftHand = forkarr[0].ch
	} else {
		p1.leftHand = forkarr[Id+1].ch
	}
	p1.rightHand = forkarr[Id].ch

	return p1

}

type fork struct {
	isUsed bool
	forkId int
	ch     chan int
}

func buildF(Id int) fork {
	var f1 fork
	f1.forkId = Id
	f1.ch = make(chan int)

	return f1
}

func main() {
	var forks = [5]fork{}
	var philosophers = [5]philosopher{}

	//channel := make(chan int)
	for i := 0; i < 5; i++ {
		forks[i] = buildF(i)
	}
	for i := 0; i < 5; i++ {
		philosophers[i] = buildP(i, forks[:])
	}

}

func (p philosopher) changeState() {

	if p.isEating {
		fmt.Println(p.philoId, " is eating")
		p.isEating = false
	} else {
		fmt.Println(p.philoId, " is thinking")
		p.isEating = true
		p.leftHand <- 33
		p.rightHand <- 22

	}

}
