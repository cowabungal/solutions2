package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

/*
В последнем задании 2-го модуля нужно было обработать ошибку io.EOF теперь
попробуйте обработать ее с помощью пакета errors.
○ Основная логика должна быть вынесена в отдельную функцию, которая
возвращает результат и ошибку. Это нужно для того чтобы в случае ошибки
мы могли проверить ее тип.
Создайте переменную limit укажите ее в произвольное значение и если counter
строк превысит лимит вы должны выбросить ошибку.
Ошибка при превышении лимита, должна быть структурой, которая
имплементирует интерфейс error и содержит необходимые свойства для вывода:
fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
В функции main, обработайте кейс с превышением лимита, при помощи пакета
errors. Кейс не должен создавать панику, лишь печатать сообщение с ошибкой
fmt.Println("string count exceed limit, please read another file =) err: ", err.Error()) в
случае, если получена наша имплементация интерфейса error.
*/

type myError struct {
	message string
	limit int
	lastString string
}

func (e *myError) Error() string {
	return e.message
}

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

	_, err = Reader(file)
	var enf *myError

	if errors.Is(err, io.EOF) {
		log.Printf("конец файла\n")

	} else if errors.As(err, &enf) {
		log.Printf("string count exceed limit, please read another file =) err: %s\n", err.Error())
	}
}

func Reader(file *os.File) (int, error) {
	limitErr := myError{
		message:    "limit line for reading",
		limit:      20,
		lastString: "",
	}

	reader := bufio.NewReader(file)

	for i := 1; ; i++ {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			return i - 1, err
		}

		if i > limitErr.limit {
			limitErr.lastString = string(rune(i - 1))
			return i - 1, &limitErr
		}
	}

}
