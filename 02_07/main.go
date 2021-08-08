package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Напишите программу, которая считывает данные из файла, проверяет, что все
поля заполнены и записывает считанные данные в файл data/data_out.txt в
формате Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n.
Если какое-то поле не заполнено, то программа должна вызвать панику, передав
строку вида parse error: empty field on string %d.
Объявите необходимые отложенные вызовы.
Обработайте панику таким образом, чтобы после обработки на экран вывелось
содержимое файла data_out.txt, которое было записано до возникновения паники.
*/

func main() {
	slc := make([][]string, 0)
	slcOut := make([]string, 0)

	defer func() {
		if err := recover(); err != nil {
			switch err {
			case fmt.Sprintf("parse error: empty field on string %d", len(slcOut)):
				for _, v := range slcOut {
					fmt.Print(v)
				}

			default:
				panic("fatal error")
			}
		}
	}()

	file, err := os.Open("07_task_in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		// Преобразование считываемых строк в нужный вид
		line = strings.ReplaceAll(line, "\n", "")
		slc = append(slc, strings.Split(line, "|"))
	}

	file.Close()
	os.Mkdir("data", 0777)

	file, err = os.Create("./data/data_out.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i, v := range slc {
		if v[0] == "" || v[1] == "" || v[2] == "" {
			panic(fmt.Sprintf("parse error: empty field on string %d", i))
		}

		slcOut = append(slcOut, fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, v[0], v[1], v[2]))

		writer.WriteString(slcOut[i])
		writer.Flush()
	}
}
