package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Ознакомьтесь с пакетом os и bufio.
// Скачайте файл in.txt и положите его в директорию. В файле main.go напишите код, который считывает данные из файла in.txt и
// построчно записывает их в файл out.txt, нумеруя каждую строку. Если файла out.txt
// нет, то он должен создаваться.
// С помощью отложенных вызовов закройте файловые дескрипторы. При закрытии
// файла out.txt программа должна вывести в консоль, сколько строк и байт было записано в файл.
// Реализуйте функцию logTime(), которая не принимает на вход параметров,
// определяет и выводит на экран время выполнения вашей программы. Вывод
// времени должен происходить перед завершением работы функции main().

var t = time.Now()

func main() {
	defer logTime()

	rfile, err := os.Open("06_task_in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rfile.Close()

	data := make([]byte, 10000)

	for {
		n, err := rfile.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data = data[:n]
	}

	wfile, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer wfile.Close()

	sData := string(data[:])
	byteCounter := 0
	s := strings.Split(sData, "\n")

	for i := 0; i < len(s); i++ {
		// Костыль для предотвращения последней пустой строки
		if i == len(s) - 1 {
			byteCount, err := fmt.Fprintf(wfile, "%d. %s", i+1, s[i])
			if err != nil {
				fmt.Println("unable to write to file")
				return
			}
			byteCounter += byteCount
		} else {
			byteCount, err := fmt.Fprintf(wfile, "%d. %s\n", i+1, s[i])
			if err != nil {
				fmt.Println("unable to write to file")
				return
			}
			byteCounter += byteCount
		}
	}

	fmt.Printf("\n%d byte was written\n", byteCounter)
}

func logTime() {
	fmt.Println(time.Since(t))
}
