package extensionflow

import (
	"FirstMeaw/Internal/users"
	"fmt"
	"strconv"
)

func Add(names *[]users.FullName, jobTitle *[]users.Position) {
	name := UserInput("Введите имя: ")
	lastName := UserInput("Введите фамилию: ")
	patronymic := UserInput("Введите Отчество: ")

	position := UserInput("Введите должность: ")

	newUser := users.FullName{
		FirstName:  name,
		LastName:   lastName,
		Patronymic: patronymic,
	}

	newJobTitle := users.Position{
		JobTitle: position,
	}

	*names = append((*names), newUser)
	*jobTitle = append((*jobTitle), newJobTitle)

	fmt.Println("Досье добавлено: ", names, jobTitle)
}

func Delete(names *[]users.FullName, jobTitles *[]users.Position) {
	name := UserInput("Введите номер человека, который вы хотите удалить: ")

	number, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println("Ошибка: введено не число. Попробуйте снова.")
		return
	}

	index := number - 1

	if index < 0 || index < len(*names) {
		fmt.Printf("Ошибка: досье под номером %d не существует.\n", number)
		return
	}

	*names = append(
		(*names)[:index],
		(*names)[index:+1]...)
		
	*jobTitles = append(
		(*jobTitles)[:number],
		(*jobTitles)[number:+1]...,
	)

	fmt.Println("Досье удалено: ", names, jobTitles)
}

func SearchLastName(names *[]users.FullName, jobTitle *[]users.Position) {

	surname := UserInput("Введите Фамилию, чтобы найти человека: ")
	canFind := true

	for i := 0; i < len(*names); i++ {

		if surname == (*names)[i].LastName {
			fmt.Println("Информация по поиску: ", (*names)[i], (*jobTitle)[i])

		} else if canFind == false {
			fmt.Println("Такой фамилии не существует.")
		}
	}
}

func PrintDossiers(names *[]users.FullName, jobTitle *[]users.Position) {
	for i := 0; i < len(*names); i++ {

		fmt.Printf("%d) ФИО, Должность: %s %s\n", i+1, (*names)[i], (*jobTitle)[i])
	}
}
