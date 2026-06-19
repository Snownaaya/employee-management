package positions

import (
	"fmt"
	"strconv"

	"FirstMeaw/Internal/extensionflow"
	"FirstMeaw/Internal/users"
)

func AddEmployee(employees map[users.Position]*[]users.FullName) {
	positionInput := extensionflow.UserInput("Введите должность: ")
	position := users.Position{JobTitle: positionInput}

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
	fullNames := employees[position]
	*fullNames = append(*fullNames, newUser)

	fmt.Println("Добавлено: ", newUser, position)
}

func RemoveEmployee(employees map[users.Position]*[]users.FullName) {
	positionInput := extensionflow.UserInput("Введите должность: ")
	position := users.Position{JobTitle: positionInput}

	employeeList, ok := employees[position]
	if ok != true {
		fmt.Println("Должность не найдена.")
		return
	}

	if len(*employeeList) == 0 {
		fmt.Println("Список сотрудников пуст.")
		return
	}

	for i, employee := range *employeeList {
		fmt.Println(i+1, ".", employee)
	}

	numberInput := extensionflow.UserInput("Введите номер сотрудника для удаления: ")

	choice, err := strconv.Atoi(numberInput)
	if err != nil {
		fmt.Println("Некорректный номер.")
		return
	}

	index := choice - 1

	if index < 0 || index >= len(*employeeList) {
		fmt.Println("Ошибка: сотрудника не существует.")
		return
	}

	*employeeList = append(
		(*employeeList)[:index],
		(*employeeList)[index+1:]...,
	)

	fmt.Println("Сотрудник удалён.")

	if len(*employeeList) == 0 {
		delete(employees, position)
		fmt.Println("Должность удалена, так как сотрудников больше не осталось.")
	}
}

func FullInfo(employees map[users.Position]*[]users.FullName) {
	if len(employees) == 0 {
		fmt.Println("Список пуст.")
		return
	}

	number := 1

	for position, userName := range employees {
		for _, name := range *userName {
			fmt.Println(number, ".", position.JobTitle, "-", name)
			number++
		}
	}
}
