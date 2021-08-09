package main

import (
	"fmt"
	"time"
)


/*
Используя функции spinner и fib, напишите программу, которая параллельно
считает 44ое и 45ое числа Фибоначчи, выводя при этом спиннер в консоль, и
выводит оба числа в консоль.
*/

func main() {
	go spinner(100 * time.Millisecond)

	go func(){
		fibNN := fib(45)
		fmt.Printf("\rFibonacci(%d) = %d\n", 45, fibNN)
	}()

	fibN := fib(44)
	fmt.Printf("\rFibonacci(%d) = %d\n", 44, fibN)

	time.Sleep(5 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
