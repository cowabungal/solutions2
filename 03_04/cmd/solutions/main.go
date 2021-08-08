package main

import (
	"fmt"
	"github.com/cowabungal/solutions/Internal/myrandom"
	"github.com/huandu/xstrings"
)

/*
1. Ознакомьтесь со структурой проекта project-layout.
2. Создайте новый проект(модуль go mod) c названием gopackages-layout.
3. Код из второго задания модуля перенесите в этот проект, в качестве структуры
используйте project-layout.
4. Используйте только нужные директории.
5. Запустите его и проверьте, что все работает корректно.
*/

func main()  {
	fmt.Println("Hello world")

	fmt.Println(xstrings.Shuffle(myrandom.City()))
	fmt.Println(myrandom.Digit())

}
