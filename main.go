package main

import (
	"fmt"
)

type GoldStore struct {
	Name         string
	GoldRate     float64 // Per gram in INR
	MakingCharge float64 // Per gram in INR
	Discount     float64 // In percentage
}

func calculateTotalCost(store GoldStore, weight float64) float64 {
	totalPrice := (store.GoldRate * weight) + (store.MakingCharge * weight)
	discountAmount := totalPrice * (store.Discount / 100)
	return totalPrice - discountAmount
}

func main() {
	goldStores := []GoldStore{
		{"Tanishq", 6200, 750, 5},
		{"Kalyan Jewellers", 6150, 700, 6},
		{"Malabar Gold & Diamonds", 6180, 680, 7},
		{"PC Jeweller", 6130, 720, 4},
		{"Local Jewelry Store", 6100, 650, 8},
	}

	ringWeight := 5.0     // grams
	braceletWeight := 9.0 // grams
	bangleWeight := 18.0  // grams

	fmt.Println("Gold Store Comparison:")
	bestStore := ""
	bestPrice := 1e9 // Initialize with a high value

	for _, store := range goldStores {
		totalCost := calculateTotalCost(store, ringWeight) + calculateTotalCost(store, braceletWeight) + calculateTotalCost(store, bangleWeight)
		fmt.Printf("%s - Total Cost: ₹%.2f\n", store.Name, totalCost)

		if totalCost < bestPrice {
			bestPrice = totalCost
			bestStore = store.Name
		}
	}

	fmt.Printf("\nBest store for purchase: %s with total cost ₹%.2f\n", bestStore, bestPrice)
}
