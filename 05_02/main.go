package main

import (
	"fmt"
	"runtime"
	"time"
)


/*
Используя код из предыдущего задания, с помощью GOMAXPROCS добейтесь
поочередного выполнения функций подсчета 32го и 33го чисел Фибоначчи,
запущенных в отдельных горутинах, и вывода их в консоль, отображая при этом
спиннер. Устанавливайте GOMAXPROCS путем добавления строки
runtime.GOMAXPROCS(1) в начало функции main. Учтите, что для вытеснения
основного потока нужно вызвать какую-то из блокирующих операций (например,
time.Sleep()).
*/

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		spinner(100 * time.Millisecond)
	}()

	go func(){
		fibNN := fib(32)
		fmt.Printf("\rFibonacci(%d) = %d\n", 32, fibNN)
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		fibN := fib(33)
		fmt.Printf("\rFibonacci(%d) = %d\n", 33, fibN)
		time.Sleep(100 * time.Millisecond)
	}()

	time.Sleep(1 * time.Second)
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
