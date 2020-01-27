package main

import (
	"sort"
)

type House struct {
	id                 int
	price              int
	info               string
	distanceFromCenter int
	region             string
}

func sortByPriceDesс(a, b House) bool {
	return a.price > b.price
}
func sortByPriceAsc(a, b House) bool {
	return a.price < b.price
}

func sortByPrice(data []House, howToSort string) []House {
	if howToSort == "DESC" {
		sort.Slice(data, func(i, j int) bool {
			return sortByPriceDesс(data[i], data[j])
		})
	} else if howToSort == "ASC" {
		sort.Slice(data, func(i, j int) bool {
			return sortByPriceAsc(data[i], data[j])
		})
	}
	return data
}

func sortByDistanceDesс(a, b House) bool {
	return a.distanceFromCenter > b.distanceFromCenter
}
func sortByDistanceAsc(a, b House) bool {
	return a.distanceFromCenter < b.distanceFromCenter
}

func sortByDistance(data []House, howToSort string) []House {
	if howToSort == "DESC" {
		sort.Slice(data, func(i, j int) bool {
			return sortByDistanceDesс(data[i], data[j])
		})
	} else if howToSort == "ASC" {
		sort.Slice(data, func(i, j int) bool {
			return sortByDistanceAsc(data[i], data[j])
		})
	}
	return data
}

func searchByMaxPrice(data []House, maxPrice int) []House {
	result := make([]House, 0)
	for _, value := range data {
		if value.price <= maxPrice {
			result = append(result, value)
		}
	}
	return result
}

func searchByPriceRange(data []House, minPrice, maxPrice int) []House {
	result := make([]House, 0)
	for _, value := range data {
		if value.price <= maxPrice && value.price >= minPrice {
			result = append(result, value)
		}
	}
	return result
}

func searchByRegion(data []House, regions []string) []House {
	result := make([]House, 0)
	for _, house := range data {
		for _, region := range regions {
			if house.region == region {
				result = append(result, house)
			}
		}
	}
	return result
}

func main() {}
