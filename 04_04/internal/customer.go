package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	discount     bool
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

func (c *Customer) CalcDiscount() (int,error) {
	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
			return 0, nil
		}
		return result, nil
}
