package main

import (
	"errors"
	"fmt"
	"log"
	"myapp/internal"
)

/*
В пакете main, используя логику из функции func startTransaction(debtor
internal.Debtor) error (функция из описания предыдущего задания), разместите
функцию startTransactionDynamic, которая не выдаст ошибку на этапе компиляции, если ей передать неверный аргумент, а вернет ее в рантайме из функции
вызывающей стороне (errors.New("Incorrect type")). Логику не меняйте.
*/

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	fmt.Printf("%+v\n", cust)

	cashe, err := internal.CalcPrice(cust, 2000)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(cashe)

	partner := internal.NewPartner("Alex", 18, 100000, 1000, true)
	startTransactionDynamic(partner)

	fmt.Println(partner)
}

func startTransactionDynamic(p interface{}) error {
	partner, ok := p.(*internal.Partner)
	if !ok {
		return errors.New("incorrect type")
	}
	return partner.WrOffDebt()
}
