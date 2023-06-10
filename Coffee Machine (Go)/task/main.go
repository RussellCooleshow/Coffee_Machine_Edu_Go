package main

import (
	"fmt"
)

type CoffeeMachine struct {
	availWater       int
	availMilk        int
	availCoffeeBeans int
	availCups        int
	availMoney       int
}

func main() {

	coffeeMachine := CoffeeMachine{

		availWater:       400,
		availMilk:        540,
		availCoffeeBeans: 120,
		availCups:        9,
		availMoney:       550,
	}
	// ...

	for {

		var action string
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		switch action {
		case "buy":
			coffeeMachine.buyCoffee()
		case "fill":
			coffeeMachine.fillResources()
		case "take":
			coffeeMachine.takeMoney()
		case "remaining":
			coffeeMachine.availResources()
		case "exit":
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}

}

func (cm *CoffeeMachine) availResources() {

	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", cm.availWater)
	fmt.Printf("%d ml of milk\n", cm.availMilk)
	fmt.Printf("%d g of coffee beans\n", cm.availCoffeeBeans)
	fmt.Printf("%d disposable cups\n", cm.availCups)
	fmt.Printf("$%d of money\n", cm.availMoney)
}

func (cm *CoffeeMachine) buyCoffee() {

	var response string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&response)

	switch response {
	case "1":
		cm.makeCoffee(250, 0, 16, 4)
	case "2":
		cm.makeCoffee(350, 75, 20, 7)
	case "3":
		cm.makeCoffee(200, 100, 12, 6)
	case "back":
		return
	default:
		fmt.Println("Invalid choice! Try again.")

	}

}

func (cm *CoffeeMachine) makeCoffee(water, milk, coffeeBeans, cost int) {
	if cm.hasEnoughResources(water, milk, coffeeBeans, 1) {
		fmt.Println("I have enough resources, making you a coffee!")
		cm.availWater -= water
		cm.availMilk -= milk
		cm.availCoffeeBeans -= coffeeBeans
		cm.availMoney += cost
		cm.availCups--
	} else {
		fmt.Println("Sorry, not enough resources to make coffee!")
	}
}

func (cm *CoffeeMachine) fillResources() {
	fmt.Println("Write how many ml of water you want to add:")
	var water int
	fmt.Scan(&water)

	fmt.Println("Write how many ml of milk you want to add:")
	var milk int
	fmt.Scan(&milk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	var coffeeBeans int
	fmt.Scan(&coffeeBeans)

	fmt.Println("Write how many disposable cups you want to add:")
	var cups int
	fmt.Scan(&cups)

	cm.availWater += water
	cm.availMilk += milk
	cm.availCoffeeBeans += coffeeBeans
	cm.availCups += cups
}

func (cm *CoffeeMachine) takeMoney() {
	fmt.Printf("I gave you $%d\n", cm.availMoney)
	cm.availMoney = 0
}

func (cm *CoffeeMachine) hasEnoughResources(water, milk, coffeeBeans, cups int) bool {
	return cm.availWater >= water && cm.availMilk >= milk && cm.availCoffeeBeans >= coffeeBeans && cm.availCups >= cups
}
