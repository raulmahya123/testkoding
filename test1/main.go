package main

import (
	"encoding/json"
	"fmt"
)

type FruitType string

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type Fruit struct {
	FruitId   int       `json:"fruitId"`
	FruitName string    `json:"fruitName"`
	FruitType FruitType `json:"fruitType"`
	Stock     int       `json:"stock"`
}

func main() {
	// Data JSON yang diberikan
	jsonData := `[ 
		{
			"fruitId": 1,
			"fruitName": "Apel",
			"fruitType": "IMPORT",
			"stock": 10
		},
		{
			"fruitId": 2,
			"fruitName": "Kurma",
			"fruitType": "IMPORT",
			"stock": 20
		},
		{
			"fruitId": 3,
			"fruitName": "Apel",
			"fruitType": "IMPORT",
			"stock": 50
		},
		{
			"fruitId": 4,
			"fruitName": "Manggis",
			"fruitType": "LOCAL",
			"stock": 100
		},
		{
			"fruitId": 5,
			"fruitName": "Jeruk Bali",
			"fruitType": "LOCAL",
			"stock": 10
		},
		{
			"fruitId": 5,
			"fruitName": "KURMA",
			"fruitType": "IMPORT",
			"stock": 20
		},
		{
			"fruitId": 5,
			"fruitName": "Salak",
			"fruitType": "LOCAL",
			"stock": 150
		}
	]`

	// Parsing JSON ke dalam slice dari struct Fruit
	var fruits []Fruit
	err := json.Unmarshal([]byte(jsonData), &fruits)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 1. Buah apa saja yang dimiliki Andi? (fruitName)
	fmt.Println("1. Buah apa saja yang dimiliki Andi?")
	printFruitNames(fruits)

	// 2. Andi memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah (fruitType).Berapa jumlah wadah yang dibutuhkan? Dan ada buah apa saja di masing-masing wadah?
	fmt.Println("\n2. Andi memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah (fruitType).")
	fmt.Println("Berapa jumlah wadah yang dibutuhkan? Dan ada buah apa saja di masing-masing wadah?")
	printFruitWadah(fruits)

	// 3. Berapa total stock buah yang ada di masing-masing wadah?
	fmt.Println("\n3. Berapa total stock buah yang ada di masing-masing wadah?")
	printTotalStock(fruits)
}

// Fungsi untuk mencetak buah apa saja yang dimiliki Andi
func printFruitNames(fruits []Fruit) {
	for _, fruit := range fruits {
		fmt.Println(fruit.FruitName)
	}
}

// Fungsi untuk memisahkan buah ke dalam wadah berdasarkan tipe buah dan mencetak jumlah wadah yang dibutuhkan serta buah apa saja di masing-masing wadah
func printFruitWadah(fruits []Fruit) {
	wadah := make(map[FruitType][]string)

	for _, fruit := range fruits {
		wadah[fruit.FruitType] = append(wadah[fruit.FruitType], fruit.FruitName)
	}

	fmt.Println("Jumlah wadah yang dibutuhkan:", len(wadah))

	for tipe, buah := range wadah {
		fmt.Printf("%s: %v\n", tipe, buah)
	}
}

// Fungsi untuk mencetak total stock buah yang ada di masing-masing wadah
func printTotalStock(fruits []Fruit) {
	wadah := make(map[FruitType]int)

	for _, fruit := range fruits {
		wadah[fruit.FruitType] += fruit.Stock
	}

	for tipe, stock := range wadah {
		fmt.Printf("%s: %d\n", tipe, stock)
	}
}
