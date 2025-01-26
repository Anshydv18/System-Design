package main

//implement the interface{} int smaller segmented interface such that , client don't have to implement unnecessary function



type WaiterInterface interface{
	ServeDish()  string
	TakeOrder() string
}

type HotelEmployee interface{
	WaiterInterface
	Washdishes() string
}

//break the interface in segmented , so that the client do not import unnnecssary function
//that is not useful for them