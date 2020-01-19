package main

type Price struct {
	currency string
	value int64
}

type Address struct {
	distanceFromCenter int64
}

type House struct{
	id int64
	price Price
	address Address
}

func main() {

}

func sortByPrice(data []House, less func(a,b House)bool){
	
}

