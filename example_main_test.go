package main

import "fmt"

var flat = []flats{
	{
		id:                                 1,
		name:                               "Сдам",
		stateOfTheHouse:                    "Люкс",
		price:                              150_000,
		currency:                           "TJS",
		flatRoomCount:                      3,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 9700,
	},
	{
		id:                                 2,
		name:                               "Продам",
		stateOfTheHouse:                    "Люкс",
		price:                              230_000,
		currency:                           "TJS",
		flatRoomCount:                      2,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 2400,
	},
	{
		id:                                 3,
		name:                               "Продам",
		stateOfTheHouse:                    "Недавно ремонтировано",
		price:                              552_900,
		currency:                           "TJS",
		flatRoomCount:                      2,
		town:                               "Челябинск",
		radiusFromTheLocationOfFinderMetre: 6560,
	},
	{
		id:                                 4,
		name:                               "Продам",
		stateOfTheHouse:                    "Есть мебели, стиральная машина, отремонтирован",
		price:                              1_550_000,
		currency:                           "TJS",
		flatRoomCount:                      5,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 3500,
	},
	{
		id:                                 5,
		name:                               "Сдам",
		stateOfTheHouse:                    "Люкс",
		price:                              23_000,
		currency:                           "TJS",
		flatRoomCount:                      1,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 2_500,
	},
	{
		id:                                 1,
		name:                               "Продам",
		stateOfTheHouse:                    "Люкс",
		price:                              892_000,
		currency:                           "TJS",
		flatRoomCount:                      4,
		town:                               "Худжанд",
		radiusFromTheLocationOfFinderMetre: 520_000,
	},
}
//
func ExampleSortByPriceAsc() {
	ascSort := sortByPriceAsc(flat)
	fmt.Println(ascSort)
	// Output: [{5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {1 Продам Люкс 892000 TJS 4 Худжанд 520000} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500}]
}
//
func ExampleSortByPriceDesc() {
	descSort := sortByPriceDesc(flat)
	fmt.Println(descSort)
	// Output: [{4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSortByRadiusAsc() {
	ascSorted := sortByRadiusAsc(flat)
	fmt.Println(ascSorted)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}
func ExampleSortByRadiusDesc() {
	descSorted := sortByRadiusDesc(flat)
	fmt.Println(descSorted)
	// Output: [{1 Продам Люкс 892000 TJS 4 Худжанд 520000} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {2 Продам Люкс 230000 TJS 2 Душанбе 2400}]
}

func ExampleSearchByPrice_NoResults() {
	noResult := searchByPrice(flat, 20_000)
	fmt.Println(noResult)
	// Output: []
}
func ExampleSearchByPrice_OneResult() {
	noResult := searchByPrice(flat, 23_000)
	fmt.Println(noResult)
	// Output: [{5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}
func ExampleSearchByPrice_MoreResult() {
	moreResult := searchByPrice(flat, 600_000)
	fmt.Println(moreResult)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchWithinPrice_NoResults() {
	result := searchWithinPrice(flat, 24_000, 120_000)
	fmt.Println(result)
	// Output: []
}
func ExampleSearchWithinPrice_OneResults() {
	result := searchWithinPrice(flat, 140_000, 151_000)
	fmt.Println(result)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700}]
}
func ExampleSearchWithinPrice_MoreResults() {
	result := searchWithinPrice(flat, 230_000, 2_500_000)
	fmt.Println(result)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}

func ExampleSearchByTown_NoResult() {
	result := findByTown(flat, "Казань")
	fmt.Println(result)
	// Output: []
}
func ExampleSearchByTown_OneResult() {
	result := findByTown(flat, "Челябинск")
	fmt.Println(result)
	// Output: [{3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560}]
}
func ExampleSearchByTown_MoreResults() {
	result := findByTown(flat, "Душанбе")
	fmt.Println(result)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}
func ExampleSearchByTowns_NoResult() {
	result := findByTowns(flat, []string{"Казань", "Хорог"})
	fmt.Println(result)
	// Output: []
}
func ExampleSearchByTowns_OneResult() {
	result := findByTowns(flat, []string{"Питер", "Худжанд"})
	fmt.Println(result)
	// Output: [{1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}
func ExampleSearchByTowns_MoreResults() {
	result := findByTowns(flat, []string{"Худжанд", "Душанбе"})
	fmt.Println(result)
	// Output: [{1 Продам Люкс 892000 TJS 4 Худжанд 520000} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchByRadiusAsc() {
	ascRadius := sortByRadiusAsc(flat)
	fmt.Println(ascRadius)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}
func ExampleSearchByRadiusDesc() {
	descRadius := sortByRadiusAsc(flat)
	fmt.Println(descRadius)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}

func ExampleSearchFlatsOnRadius2500Metres_NoResult() {
	radiusMetre := searchFlatsOnRadius2500Metres(flat, 2_000)
	fmt.Println(radiusMetre)
	// Output: []
}
func ExampleSearchFlatsOnRadius2500Metres_OneResult() {
	radiusMetre := searchFlatsOnRadius2500Metres(flat, 2_450)
	fmt.Println(radiusMetre)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400}]
}
func ExampleSearchFlatsOnRadius2500Metres_MoreResult() {
	radiusMetre := searchFlatsOnRadius2500Metres(flat, 3_500)
	fmt.Println(radiusMetre)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchFlatsByStateOfTheHouse_NoResult() {
	stateOfFlat := searchFlatsByStateOfTheHouse(flat, "Не в хорошем состоянии")
	fmt.Println(stateOfFlat)
	// Output: []
}
func ExampleSearchFlatsByStateOfTheHouse_OneResult() {
	stateOfFlat := searchFlatsByStateOfTheHouse(flat, "Недавно ремонтировано")
	fmt.Println(stateOfFlat)
	// Output: [{3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560}]
}
func ExampleSearchFlatsByStateOfTheHouse_MoreResult() {
	stateOfFlat := searchFlatsByStateOfTheHouse(flat, "Люкс")
	fmt.Println(stateOfFlat)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}

func ExampleSearchAllFlatsByFlatRoomCount_NoResult() {
	FlatRoomCount := searchAllFlatsByFlatRoomCount(flat, 6)
	fmt.Println(FlatRoomCount)
	// Output: []
}
func ExampleSearchAllFlatsByFlatRoomCount_OneResult() {
	FlatRoomCount := searchAllFlatsByFlatRoomCount(flat, 5)
	fmt.Println(FlatRoomCount)
	// Output: [{4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500}]
}
func ExampleSearchAllFlatsByFlatRoomCount_MoreResult() {
	FlatRoomCount := searchAllFlatsByFlatRoomCount(flat, 2)
	fmt.Println(FlatRoomCount)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Челябинск 6560}]
}

