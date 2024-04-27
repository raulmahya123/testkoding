package main

import "fmt"

type IFruit struct {
	fruitId   int
	fruitName string
	fruitType string
	stock     int
}

func main() {
	// Deklarasi data buah
	fruits := []IFruit{
		{
			fruitId:   1,
			fruitName: "Apel",
			fruitType: "IMPORT",
			stock:     10,
		},
		{
			fruitId:   2,
			fruitName: "Kurma",
			fruitType: "IMPORT",
			stock:     20,
		},
		{
			fruitId:   3,
			fruitName: "Apel",
			fruitType: "IMPORT",
			stock:     50,
		},
		{
			fruitId:   4,
			fruitName: "Manggis",
			fruitType: "LOCAL",
			stock:     100,
		},
		{
			fruitId:   5,
			fruitName: "Jeruk Bali",
			fruitType: "LOCAL",
			stock:     10,
		},
		{
			fruitId:   6,
			fruitName: "Kurma",
			fruitType: "IMPORT",
			stock:     20,
		},
		{
			fruitId:   7,
			fruitName: "Salak",
			fruitType: "LOCAL",
			stock:     150,
		},
	}

	//  1.Buah apa saja yang dimiliki Andi? (fruitName)
	fmt.Println("\nJawaban:No. 1")
	fmt.Println("Buah yang dimiliki Andi:")
	for _, fruit := range fruits {
		fmt.Println(fruit.fruitName)
	}

	// 2.Andi memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah
	// (fruitType). Berapa jumlah wadah yang dibutuhkan? Dan ada buah apa saja di
	// masing-masing wadah?
	fmt.Println("\nJawaban:No. 2")
	importFruits := make(map[string][]IFruit)
	localFruits := make(map[string][]IFruit)

	for _, fruit := range fruits {
		if fruit.fruitType == "IMPORT" {
			importFruits[fruit.fruitName] = append(importFruits[fruit.fruitName], fruit)
		} else {
			localFruits[fruit.fruitName] = append(localFruits[fruit.fruitName], fruit)
		}
	}

	// 3.Berapa total stock buah yang ada di masing-masing wadah?
	fmt.Println("\nJumlah wadah yang dibutuhkan dan buah di masing-masing wadah:")
	fmt.Printf("Wadah IMPORT: %d\n", len(importFruits))
	for fruitName, fruits := range importFruits {
		fmt.Printf("- %s: %d buah\n", fruitName, len(fruits))
	}

	fmt.Printf("Wadah LOCAL: %d\n", len(localFruits))
	for fruitName, fruits := range localFruits {
		fmt.Printf("- %s: %d buah\n", fruitName, len(fruits))
	}
	fmt.Println("\nJawaban:No. 3")

	fmt.Println("\nTotal stock buah di masing-masing wadah:")
	totalStockImport := 0
	for _, fruits := range importFruits {
		for _, fruit := range fruits {
			totalStockImport += fruit.stock
		}
	}
	fmt.Printf("Total stock di wadah IMPORT: %d\n", totalStockImport)

	totalStockLocal := 0
	for _, fruits := range localFruits {
		for _, fruit := range fruits {
			totalStockLocal += fruit.stock
		}
	}
	fmt.Printf("Total stock di wadah LOCAL: %d\n", totalStockLocal)
}
