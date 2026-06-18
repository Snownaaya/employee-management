package main

import (
	"fmt"

	"FirstMeaw/Internal/employee-managementAdvenced/positions"
	"FirstMeaw/Internal/extensionflow"
	users "FirstMeaw/Internal/infrastructure"
)

const (
	CommandAddDossier    string = "1"
	CommandPrintDossier  string = "2"
	CommandDeleteDossier string = "3"
	CommanfExit          string = "4"
)

func main() {
	employees := make(map[users.Position]*[]users.FullName)
	isWork := true

	commands := map[string]func(){

		CommandAddDossier: func() {
			positions.AddEmployee(employees)
		},

		CommandPrintDossier: func() {
			positions.FullInfo(employees)
		},

		CommandDeleteDossier: func() {
			positions.RemoveEmployee(employees)
		},
		CommanfExit: func() {
			isWork = false
		},
	}

	for isWork {
		fmt.Println("Введите, чтобы добавить досье - ", CommandAddDossier)
		fmt.Println("Введите, показать все досье - ", CommandPrintDossier)
		fmt.Println("Введите, чтобы удалить досье - ", CommandDeleteDossier)
		fmt.Println("Выход - ", CommanfExit)

		userInput := extensionflow.UserInput("Введите число: ")

		action, ok := commands[userInput]
		if ok == true {
			action()
		} else {
			fmt.Println("Такой команды нет.")
		}
	}
}