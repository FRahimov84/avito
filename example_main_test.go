package main

import	"fmt"

var data = []House{
	{id:1, price:25_000, info:"1-комнатная квартира с ремонтом", distanceFromCenter:10, region:"Сино1"},
	{id:2, price:10_000, info:"1-комнатная квартира", distanceFromCenter:20, region:"Сино2"},
	{id:3, price:15_000, info:"1-комнатная квартира", distanceFromCenter:7, region:"Сино1"},
	{id:4, price:5_000, info:"1-комнатная квартира БУ", distanceFromCenter:15, region:"Сино2"},
	{id:5, price:55_000, info:"3-комнатная квартира", distanceFromCenter:5, region:"Сино3"},
	{id:6, price:125_000, info:"4-комнатная квартира с ремонтом", distanceFromCenter:6, region:"Сино1"},
}

func ExampleSortByPriceAsc()  {
	fmt.Println(sortByPrice(data, "ASC"))
	//	Output: [{4 5000 1-комнатная квартира БУ 15 Сино2} {2 10000 1-комнатная квартира 20 Сино2} {3 15000 1-комнатная квартира 7 Сино1} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {5 55000 3-комнатная квартира 5 Сино3} {6 125000 4-комнатная квартира с ремонтом 6 Сино1}]
}
func ExampleSortByPriceDesc() {
	fmt.Println(sortByPrice(data,"DESC"))
	//	Output: [{6 125000 4-комнатная квартира с ремонтом 6 Сино1} {5 55000 3-комнатная квартира 5 Сино3} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {3 15000 1-комнатная квартира 7 Сино1} {2 10000 1-комнатная квартира 20 Сино2} {4 5000 1-комнатная квартира БУ 15 Сино2}]
}

func ExampleSortByDistanceAsc() {
	fmt.Println(sortByDistance(data, "ASC"))
	//	Output: [{5 55000 3-комнатная квартира 5 Сино3} {6 125000 4-комнатная квартира с ремонтом 6 Сино1} {3 15000 1-комнатная квартира 7 Сино1} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {4 5000 1-комнатная квартира БУ 15 Сино2} {2 10000 1-комнатная квартира 20 Сино2}]
}
func ExampleSortByDistanceDesc() {
	fmt.Println(sortByPrice(data, "DESC"))
	//	Output: [{6 125000 4-комнатная квартира с ремонтом 6 Сино1} {5 55000 3-комнатная квартира 5 Сино3} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {3 15000 1-комнатная квартира 7 Сино1} {2 10000 1-комнатная квартира 20 Сино2} {4 5000 1-комнатная квартира БУ 15 Сино2}]
}

func ExampleSearchByMaxPrice_HasManyResult() {
	fmt.Println(searchByMaxPrice(data, 25_000))
	//	Output: [{1 25000 1-комнатная квартира с ремонтом 10 Сино1} {3 15000 1-комнатная квартира 7 Сино1} {2 10000 1-комнатная квартира 20 Сино2} {4 5000 1-комнатная квартира БУ 15 Сино2}]
}
func ExampleSearchByMaxPrice_HasNotAnyResult() {
	fmt.Println(searchByMaxPrice(data, 1_000))
	//	Output: []
}

func ExampleSearchPriceRange_HasNotAnyResult() {
	fmt.Println(searchByPriceRange(data, 175_000, 300_000))
	//	Output: []
}
func ExampleSearchPriceRange_HasManyResult() {
	fmt.Println(searchByPriceRange(data, 15_000, 100_000))
	//	Output: [{5 55000 3-комнатная квартира 5 Сино3} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {3 15000 1-комнатная квартира 7 Сино1}]
}

func ExampleSearchByRegion_HasManyResult() {
	fmt.Println(searchByRegion(data, []string{"Сино1", "Сино3"}))
	//	Output: [{6 125000 4-комнатная квартира с ремонтом 6 Сино1} {5 55000 3-комнатная квартира 5 Сино3} {1 25000 1-комнатная квартира с ремонтом 10 Сино1} {3 15000 1-комнатная квартира 7 Сино1}]
}
func ExampleSearchByRegion_HasNotAnyResult() {
	fmt.Println(searchByRegion(data, []string{"Сино10", "Сино5"}))
	//	Output: []
}