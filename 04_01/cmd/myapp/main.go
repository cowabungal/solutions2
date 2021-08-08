package main

import (
	"errors"
	"fmt"
	"log"
	"rebrain_go_basic/solutions2/04_01/internal"
)

/*
1. Форкните репозиторий module04 с кодом данного задания - в группу с вашими
репозиториями - golang_users_repos/<your_gitlab_id>.
2. Создайте у себя в проекте module04 из ветки master ветку module04_01.
3. Реализуйте функцию CalcPrice, которая:
○ принимает на вход указатель на структуру Customer и цену какой-то абстрактной покупки
(int);
○ возвращает в качестве первого аргумента итоговую цену с учетом скидки, а
в качестве второго аргумента - ошибку, в случае если ее вернет реализация
CalcDiscount;
○ в случае возврата из CalcDiscount ошибки первый аргумент установить в 0.
4. Функция должна быть реализована в пакете internal.

*/

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust)

	cashe, err := internal.CalcPrice(cust, 2000)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(cashe)
}
