package classes

type CheeseToppingPizza struct{
	Pizza IPizza
}

func(c *CheeseToppingPizza)GetPrice()int{
	return c.Pizza.GetPrice()+12;
}