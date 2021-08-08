package main

import (
	"fmt"
	newcolor "github.com/cowabungal/solutions/color"
	"github.com/cowabungal/solutions/myrandom"
	"github.com/fatih/color"
	"github.com/huandu/xstrings"
)

/* Создайте в проекте module03 пакет (с произвольным названием), который будет
состоять из двух go файлов и содержать 2 функции, по одной в каждом файле:
City() - возвращает случайный город,
Digit() - возвращает случайное число строчного типа (one, two, three и т.д.).
Обе функции должны формировать результат с помощью функции Random из
пакета wordz. При этом не внося никаких изменений в пакет wordz.
Выведите через fmt.Println результат функции City и Digit в файле main.go
С помощью утилиты go get, установите пакет для работы со строками xstrings.
В файле main.go примените функцию Shuffle из этого пакета к результату функции
City(). И угадайте, какое название города вывелось :) */

func main()  {
	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(xstrings.Shuffle(myrandom.City()))
	fmt.Println(myrandom.Digit())

}
