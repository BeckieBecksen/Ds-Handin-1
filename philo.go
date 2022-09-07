package main

type philosopher struct {
	isEating  bool
	leftHand  int
	rightHand int
	philoId   int
}

func buildP(Id int) philosopher {
	var p1 philosopher
	p1.philoId = Id
	return p1
}

type fork struct {
	isUsed bool
	forkId int
}

func buildF(Id int) fork {
	var f1 fork
	f1.forkId = Id
	return f1
}

func main() {

}
