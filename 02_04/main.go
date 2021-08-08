package main

import "fmt"

// Создайте map, в которой необходимо хранить информацию о выданных читателю печатных изданиях: книгах и периодических изданиях.
// Тип ключей отображения является строкой, тип значений — отображением с
// ключами типа строка и значениями с типом слайс-строк.
// Добавьте несколько произвольных значений, моделирующих наличие изданий на
// руках у читателей.
// Выведите на экран количество читателей с изданиями на руках.

func main()  {
	var Readers map[string]map[string][]string = map[string]map[string][]string{
		"Robin" : map[string][]string {
			"Harry Potter" : []string{
				"first edition",
				"second edition",
				"special edition",
			},
		},

		"Maks" : map[string][]string {
			"Think and rich" : []string{
				"exclusive edition",
				"special edition",
			},
			"Go programming for advanced" : []string{
				"first edition",
				"second edition",
			},
		},

		"Andrey" : nil,
		"Vasiliy": nil,
		"Darya": nil,
		"Alex": nil,
	}

	book_counter := 0
	for i, v := range Readers {
		if v != nil {
			book_counter++

			edition_conter := 0
			for _, slc := range v {
				edition_conter += len(slc)
			}
			fmt.Printf("У %s %d изданий\n", i, edition_conter)
		} else {
			fmt.Printf("У %s %d изданий\n", i, 0)
		}
	}

	fmt.Printf("Общее количество читателей с изданиями на руках: %d\n", book_counter)
}
