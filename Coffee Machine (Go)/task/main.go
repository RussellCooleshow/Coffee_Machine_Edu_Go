package main

import (
	"fmt"
)

func main() {
	// write your code here
	/*	processStart()
		grinding()
		boilingWater()
		mixingWaterAndBeans()
		pouringCoffee()
		pouringMilk()
		readyMessage()
		fmt.Println("Write how many cups of coffee you will need:")
		var cupsAmount int
		fmt.Scan(&cupsAmount)
		water := waterCalculation(cupsAmount)
		milk := milkCalculation(cupsAmount)
		coffeeBeans := beansCalculation(cupsAmount)

		fmt.Printf("For %d cups of coffee you will need:\n", cupsAmount)
		fmt.Printf("%d ml of water\n", water)
		fmt.Printf("%d ml of milk\n", milk)
		fmt.Printf("%d g of coffee beans\n", coffeeBeans) */

	availableResource()
}

func availableResource() {
	var availWater int
	var availMilk int
	var availCoffeeBeans int
	var cupsNeeded int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&availWater)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&availMilk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&availCoffeeBeans)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cupsNeeded)

	canMake := true

	maxCupsWater := availWater / 200
	maxCupsMilk := availMilk / 50
	maxCupsCoffeeBeans := availCoffeeBeans / 15

	maxCups := maxCupsWater
	if maxCupsMilk < maxCups {
		maxCups = maxCupsMilk
	}
	if maxCupsCoffeeBeans < maxCups {
		maxCups = maxCupsCoffeeBeans
	}

	if maxCups < cupsNeeded {
		canMake = false
	}

	if canMake {
		if maxCups > cupsNeeded {
			fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", maxCups-cupsNeeded)
		} else {
			fmt.Println("Yes, I can make that amount of coffee")
		}
	} else {
		fmt.Printf("No, I can make only %d cups of coffee\n", maxCups)
	}

}
func waterCalculation(cupsAmount int) int {
	return cupsAmount * 200
}

func milkCalculation(cupsAmount int) int {
	return cupsAmount * 50
}
func beansCalculation(cupsAmount int) int {
	return cupsAmount * 15

}

func processStart() {
	fmt.Println("Starting to make a coffee")
}
func grinding() {
	fmt.Println("Grinding coffee beans")
}
func boilingWater() {
	fmt.Println("Boiling water")
}
func mixingWaterAndBeans() {
	fmt.Println("Mixing boiled water with crushed coffee beans")
}
func pouringCoffee() {
	fmt.Println("Pouring coffee into the cup")
}
func pouringMilk() {
	fmt.Println("Pouring some milk into the cup")
}
func readyMessage() {
	fmt.Println("Coffee is ready!")
}
