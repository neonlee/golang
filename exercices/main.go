package main

import (
	"exercices/counter"
	"exercices/routines"
)

func main() {
	counter.CounterAtTen()
	//routines.RunTimeRountinesWithWaitGroup()
	routines.RoutinesWithMutexAndRutime()
	//calculator.Calculator()
}
