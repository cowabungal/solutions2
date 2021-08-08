package internal

import "errors"

type Partner struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	discount     bool
}

func NewPartner(name string, age int, balance int, debt int, discount bool) *Partner {
	return &Partner{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

func (c *Partner) CalcDiscount() (int,error) {
	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (c *Partner) WrOffDebt() error {
	c.Debt = 0
	return nil
}
