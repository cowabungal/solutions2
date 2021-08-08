package internal

func CalcPrice(c Discounter, price int) (int, error)  {
	discount, err := c.CalcDiscount()
	if err != nil {
		return 0, err
	}

	return price - discount, nil
}
