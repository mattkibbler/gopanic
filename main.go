package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi. Let's panic...")

	SimplyPanic()
	panicWithReturnValue := PanicWithNamedReturnValue()
	fmt.Printf("Value returned from PanicWithNamdReturnValue: %v", panicWithReturnValue)
	PanicInAGoGroutine()

	// Finally just panic without recovering to see what happens
	PanicWithoutRecovery()

	fmt.Println("This message should not show")
}

func PrintStartMessage(message string) {
	fmt.Println("\n===================")
	fmt.Println(message)
}

func SimplyPanic() {
	PrintStartMessage("SimplyPanic started")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	panic("SimplyPanic paniced")
}

func PanicWithNamedReturnValue() (result string) {
	PrintStartMessage("PanicWithNamedReturnValue started")
	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()
	panic("PanicWithNamedReturnValue paniced")
}

func PanicInAGoGroutine() {
	PrintStartMessage("PanicInAGoGroutine started")
	panicChannel := make(chan string)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				panicChannel <- fmt.Sprintf("Recovered from panic: %v\n", r)
			}
		}()
		// Some asynchronous stuff
		time.Sleep(1 * time.Second)
		panic("PanicInAGoGroutine paniced")
	}()

	// Wait for the panic
	messageFromChannel := <-panicChannel

	fmt.Printf("PanicInAGoGroutine finished, message from goroutine: %v", messageFromChannel)
}

func PanicWithoutRecovery() {
	PrintStartMessage("PanicWithoutRecovery started")
	panic("PanicWithoutRecovery paniced")
}
