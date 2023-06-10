package main

import (
	"fmt"
)

/* var (
	availWater       := 400
	availMilk        := 540
	availCoffeeBeans := 120
	availCups        := 9
	availMoney       := 550
) */

func main() {

	// write your code here
	//Can this block be done better?
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
	availWater := 400
	availMilk := 540
	availCoffeeBeans := 120
	availCups := 9
	availMoney := 550

	// ...

	for {
		//availableResource(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)

		var action string
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		switch action {
		case "buy":
			buyCoffee(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)
		case "fill":
			fillResources(&availWater, &availMilk, &availCoffeeBeans, &availCups)
		case "take":
			takeMoney(&availMoney)
		case "remaining":
			availableResource(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)
		case "exit":
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}

	}

	/*availableResource(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)
	requestAction(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)
	availableResource(&availWater, &availMilk, &availCoffeeBeans, &availCups, &availMoney)
	*/
}

func availableResource(availWater, availMilk, availCoffeeBeans, availCups, availMoney *int) {

	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", *availWater)
	fmt.Printf("%d ml of milk\n", *availMilk)
	fmt.Printf("%d g of coffee beans\n", *availCoffeeBeans)
	fmt.Printf("%d disposable cups\n", *availCups)
	fmt.Printf("$%d of money \n", *availMoney)
	fmt.Println()

	//Can I put the  block below into separate function?
	/*	fmt.Println("Write how many ml of water the coffee machine has:")
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
	*/
}

func buyCoffee(availWater, availMilk, availCoffeeBeans, availCups, availMoney *int) {
	var response string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&response)

	if response == "1" {
		makeCoffee(250, 0, 16, 4, availWater, availMilk, availCoffeeBeans, availCups, availMoney)
	} else if response == "2" {
		makeCoffee(350, 75, 20, 7, availWater, availMilk, availCoffeeBeans, availCups, availMoney)
	} else if response == "3" {
		makeCoffee(200, 100, 12, 6, availWater, availMilk, availCoffeeBeans, availCups, availMoney)
	} else if response == "back" {
		return
	} else {
		fmt.Println("Invalid option. Please try again.")
	}
}

func makeCoffee(water, milk, coffeeBeans, cost int, availWater, availMilk, availCoffeeBeans, availCups, availMoney *int) {
	if *availWater < water {
		fmt.Println("Sorry, not enough water!")
		return
	}
	if *availMilk < milk {
		fmt.Println("Sorry, not enough milk!")
		return
	}
	if *availCoffeeBeans < coffeeBeans {
		fmt.Println("Sorry, not enough coffee beans!")
		return
	}
	if *availCups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		return
	}

	fmt.Println("I have enough resources, making you a coffee!")
	*availWater -= water
	*availMilk -= milk
	*availCoffeeBeans -= coffeeBeans
	*availMoney += cost
	*availCups -= 1
}

func fillResources(availWater, availMilk, availCoffeeBeans, availCups *int) {
	var addWater, addMilk, addCoffeeBeans, addCups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addCoffeeBeans)
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&addCups)

	*availWater += addWater
	*availMilk += addMilk
	*availCoffeeBeans += addCoffeeBeans
	*availCups += addCups
}

func takeMoney(availMoney *int) {
	fmt.Printf("I gave you $%d\n", *availMoney)
	*availMoney = 0
}

/*func requestAction(availWater, availMilk, availCoffeeBeans, availCups, availMoney *int) {
	var action string
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&action)

	if action == "buy" {
		var response string
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
		fmt.Scan(&response)
		if response == "1" {
			*availWater -= 250
			*availCoffeeBeans -= 16
			*availMoney += 4
			*availCups -= 1
		} else if response == "2" {
			*availWater -= 350
			*availMilk -= 75
			*availCoffeeBeans -= 20
			*availMoney += 7
			*availCups -= 1

		} else if response == "3" {
			*availWater -= 200
			*availMilk -= 100
			*availCoffeeBeans -= 12
			*availMoney += 6
			*availCups -= 1

		}

	}

	if action == "fill" {
		var addWater, addMilk, addCoffeeBeans, addCups int
		fmt.Println("Write how many ml of water you want to add:")
		fmt.Scan(&addWater)
		fmt.Println("Write how many ml of milk you want to add:")
		fmt.Scan(&addMilk)
		fmt.Println("Write how many grams of coffee beans you want to add:")
		fmt.Scan(&addCoffeeBeans)
		fmt.Println("Write how many disposable cups you want to add:")
		fmt.Scan(&addCups)

		*availWater += addWater
		*availMilk += addMilk
		*availCoffeeBeans += addCoffeeBeans
		*availCups += addCups
	}

	if action == "take" {
		fmt.Printf("I gave you $%d\n", *availMoney)
		*availMoney = 0

	}

}*/

/*func waterCalculation(cupsAmount int) int {
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
*/
