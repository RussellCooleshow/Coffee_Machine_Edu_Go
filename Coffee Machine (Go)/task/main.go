package main

import (
	"fmt"
)

func main() {

	availWater := 400
	availMilk := 540
	availCoffeeBeans := 120
	availCups := 9
	availMoney := 550

	// ...

	for {

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

}

func availableResource(availWater, availMilk, availCoffeeBeans, availCups, availMoney *int) {

	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", *availWater)
	fmt.Printf("%d ml of milk\n", *availMilk)
	fmt.Printf("%d g of coffee beans\n", *availCoffeeBeans)
	fmt.Printf("%d disposable cups\n", *availCups)
	fmt.Printf("$%d of money \n", *availMoney)
	fmt.Println()

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
