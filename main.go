package main

import "sort"

type flats struct {
	id                                 int64
	name                               string
	stateOfTheHouse                    string
	price                              int64
	currency                           string
	flatRoomCount                      int64
	town                               string
	radiusFromTheLocationOfFinderMetre int64
}

func sortBy(flat []flats, less func(a, b flats) bool) []flats {
	result := make([]flats, len(flat))
	copy(result, flat)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}
func sortByPriceAsc(flat []flats) []flats {
	return sortBy(flat, func(a, b flats) bool {
		return a.price < b.price
	})
}
func sortByPriceDesc(flat []flats) []flats {
	return sortBy(flat, func(a, b flats) bool {
		return a.price > b.price
	})
}

func searchByPrice(flat []flats, limit int64) []flats {
	return filterBy(flat, func(flat flats) bool {
		return flat.price <= limit
	})
}
func searchWithinPrice(flat []flats, minPrice, maxPrise int64) []flats {
	return filterBy(flat, func(flat flats) bool {
		if flat.price < minPrice {
			return false
		}
		if flat.price > maxPrise {
			return false
		}
		return true
	})
}

func findByTown(flat []flats, town string) []flats {
	return filterBy(flat, func(flat flats) bool {
		return flat.town == town
	})
}
func findByTowns(flat []flats, towns []string) []flats {
	result := make([]flats, 0)
	for _, district := range towns  {
		result = append(result, findByTown(flat, district)...)
	}
	return result
}

func sortByRadiusAsc(flat []flats) []flats {
	return sortBy(flat, func(a, b flats) bool {
		return a.radiusFromTheLocationOfFinderMetre < b.radiusFromTheLocationOfFinderMetre
	})
}
func sortByRadiusDesc(flat []flats) []flats {
	return sortBy(flat, func(a, b flats) bool {
		return a.radiusFromTheLocationOfFinderMetre > b.radiusFromTheLocationOfFinderMetre
	})
}

func filterBy(flat []flats, predicate func(flat flats) bool) []flats {
	result := make([]flats, 0)
	for _, flat := range flat {
		if predicate(flat) {
			result = append(result, flat)
		}
	}
	return result
}
func searchFlatsOnRadius2500Metres(flat []flats, radiusMetre int64) []flats {
	return filterBy(flat, func(flat flats) bool {
		return flat.radiusFromTheLocationOfFinderMetre <= radiusMetre
	})
}
func searchFlatsByStateOfTheHouse(flat []flats, stateOfTheHouse string) []flats {
	return filterBy(flat, func(flat flats) bool {
		return flat.stateOfTheHouse == stateOfTheHouse
	})
}
func searchAllFlatsByFlatRoomCount(flat []flats, FlatRoomCount int64) []flats {
	return filterBy(flat, func(flat flats) bool {
		return flat.flatRoomCount == FlatRoomCount
	})
}


func main() {

}