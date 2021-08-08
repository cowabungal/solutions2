package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
Напишите программу, которая построчно считывает файл.
Выведите в консоль количество строк в формате Total strings: %d.
Корректно обработайте в отложенном вызове ошибки закрытия файловых
дескрипторов.
Корректно обработайте ошибку окончания файла EOF.
*/

func main() {
	file, err := os.Open("08_task_in.txt")
	if err != nil {
		log.Printf("ошибка при открытии файла: %s\n", file.Name())
		os.Exit(1)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("ошибка при закрытии файла: %s: %s\n", file.Name(), err.Error())
		}
	}()

	reader := bufio.NewReader(file)

	for i := 1; ; i++ {
		_, err := reader.ReadString('\n')
		if err == io.EOF { // если конец файла
			log.Printf("конец файла на строке: %d\n", i)
			fmt.Printf("Total strings: %d\n", i-1)
			break
		}
	}
}
