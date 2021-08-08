package main

import (
	"fmt"
	"github.com/cowabungal/solutions/myrandom"
	"github.com/huandu/xstrings"
)

/* Переведите проект из предыдущего задания на Go mod (если вы этого не сделали по ходу этого задания).
Уберите использование библиотеки github.com/fatih/color и приведите в
соответствие файлы go.mod и go.sum с помощью команды go mod tidy.
Измените версию библиотеки github.com/huandu/xstrings на 1.2.1.
Соберите проект в режиме vendor, директорию vendor добавьте в .gitignore */

func main()  {
	fmt.Println("Hello world")

	fmt.Println(xstrings.Shuffle(myrandom.City()))
	fmt.Println(myrandom.Digit())

}
