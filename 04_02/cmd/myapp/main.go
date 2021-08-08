package main

import (
	"fmt"
	"log"
	"rebrain_go_basic/solutions2/04_02/internal"
)

/*
Предположим, что теперь нам необходимо не позволить внешнему коду
предоставлять нам реализацию свойства CalcDiscount, а жестко задать реализацию в привязке к структуре Customer, для этого сделайте CalcDiscount
методом, а не свойством структуры Customer:
○ Логику оставить, как была в функции CalcDiscount.
○ Константу DEFAULT_DISCOUNT перенести в пакет internal.
○ Свойство Discount сделать нередактируемым вне пакета internal.

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
