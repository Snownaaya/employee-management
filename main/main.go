package main

import (
	"fmt"

	"FirstMeaw/Internal/extensionflow"
	"FirstMeaw/Internal/infrastructure"
)

const (
	CommandAddDossier    string = "1"
	CommandPrintDossier  string = "2"
	CommandDeleteDossier string = "3"
	CommandSearchName    string = "4"
	CommanfExit          string = "5"
)

func main() {

	var names []users.FullName
	var positions []users.Position
	isWork := true

	for isWork {

		fmt.Println("Введите, чтобы добавить досье - ", CommandAddDossier)
		fmt.Println("Введите, показать все досье - ", CommandPrintDossier)
		fmt.Println("Введите, чтобы удалить досье - ", CommandDeleteDossier)
		fmt.Println("Введите, поиск по Фамилии - ", CommandSearchName)
		fmt.Println("Выход.")

		userInput := extensionflow.UserInput("Введите число: ")

		switch userInput {

		case CommandAddDossier:
			extensionflow.Add(&names, &positions)

		case CommandPrintDossier:
			extensionflow.PrintDossiers(&names, &positions)

		case CommandDeleteDossier:
			extensionflow.Delete(&names, &positions)

		case CommandSearchName:
			extensionflow.SearchLastName(&names, &positions)

		case CommanfExit:
			isWork = false

		default:
			fmt.Println("Такой команды нет.")
		}
	}
}