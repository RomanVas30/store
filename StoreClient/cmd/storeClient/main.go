package main

import (
	"fmt"
	"github.com/RomanVas30/storeClient/internal/service"
	"os"
)

func main() {
	for {
		currToken, currUserName := AuthMenu()
		StoreMenu(currToken, currUserName)
	}
}

func AuthMenu() (string, string) {
	for {
		fmt.Println("You need to log in:")
		fmt.Println("1. Sign In")
		fmt.Println("2. Sign Up")
		fmt.Println("3. Exit")

		var command int
		if _, err := fmt.Scanf("%d\n", &command); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return "", ""
		}

		if command == 1 {
			var err error
			_, currToken, err := service.SignInUI()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return currToken, ""
			}
			break
		} else if command == 2 {
			if err := service.SignUpUI(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				continue
			}
			fmt.Println("Successful registration!")
		} else if command == 3 {
			fmt.Println("Goodbye!")
			os.Exit(0)
		} else {
			fmt.Println("An unknown command has been entered")
		}
	}
	return "", ""
}

func StoreMenu(token string, username string) {
	fmt.Println("Welcome to the store:)")
	for {
		fmt.Println("Select an action from the list:")
		fmt.Println("1. Show all products")
		fmt.Println("2. Show product info")
		fmt.Println("3. Create order")
		fmt.Println("4. Add product to order")
		fmt.Println("5. Show order")
		fmt.Println("6. Order payment")
		fmt.Println("7. Sign Out")

		var command int
		if _, err := fmt.Scanf("%d\n", &command); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			continue
		}

		if command == 1 {
			if err := service.ShowProducts(token); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				continue
			}
		} else if command == 2 {

		} else if command == 3 {

		} else if command == 4 {

		} else if command == 5 {

		} else if command == 6 {

		} else if command == 7 {
			fmt.Println("Sign out processing...")
			return
		} else {
			fmt.Println("An unknown command has been entered")
		}
	}
}
