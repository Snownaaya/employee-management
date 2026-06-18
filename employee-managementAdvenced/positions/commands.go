package positions

import (
	"fmt"

	"FirstMeaw/Internal/extensionflow"
	users "FirstMeaw/Internal/infrastructure"
)

func AddEmployee(employees map[users.Position]*[]users.FullName) {
	positionInput := extensionflow.UserInput("Введите должность: ")
	position := users.Position{positionInput}

	name := extensionflow.UserInput("Введите имя: ")
	lastName := extensionflow.UserInput("Введите фамилию: ")
	patronymic := extensionflow.UserInput("Введите Отчество: ")

	newUser := users.FullName{
		FirstName:  name,
		LastName:   lastName,
		Patronymic: patronymic,
	}

	if employees == nil {
		return
	}

	if employees[position] == nil {
		emptySlice := []users.FullName{}
		employees[position] = &emptySlice
	}
	*employees[position] = append(*employees[position], newUser)

	fmt.Println("Добавлено: ", newUser, position)
}

func RemoveEmployee(employees map[users.Position]*[]users.FullName) {
	positionInput := extensionflow.UserInput("Введите должность: ")
	position := users.Position{positionInput}

	lastNameInput := extensionflow.UserInput("Введите фамилию сотрудника, чтобы удалить: ")
	lastName := users.FullName{
		LastName: lastNameInput,
	}
	if employees == nil {
		fmt.Println("Такого человека нет: ")
		return
	}

	employeeList, ok := employees[position]
	if ok == false {
		return
	}

	*employeeList = append((*employeeList), lastName)

	delete(employees, position)
	fmt.Println("Удален сотрудник: ", position)
}

func FullInfo(employees map[users.Position]*[]users.FullName) {
	if employees == nil{
		fmt.Println("Список пуст.")
		return
	}

	for	position, userName := range employees{
		fmt.Println("Должность: ", position.JobTitle)
		
		for i := 0; i < len(*userName); i++ {
			fmt.Println("ФИО: ", (*userName)[i])
		}
	}
}