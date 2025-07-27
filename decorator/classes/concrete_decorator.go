package classes


type TommatoTopping struct{
	Pizza IPizza
}

func(t *TommatoTopping)GetPrice()int{
	return t.Pizza.GetPrice()+7
}