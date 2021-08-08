package main

import (
	"fmt"
	"log"
	"myapp/internal"
)

/*
Добавьте в пакет internal интерфейс Discounter, который требует следующей
сигнатуры метода:
CalcDiscount() (int, error)
Измените функцию CalcPrice так, чтобы на вход она теперь принимала объект,
реализующий интерфейс Discounter.

*/

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	fmt.Printf("%+v\n", cust)

	cashe, err := internal.CalcPrice(cust, 2000)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(cashe)
}
