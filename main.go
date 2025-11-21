package main

import (
	"errors"
	"fmt"
)

func main() {
	m := map[string]string{
		"PurpleSchool": "purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["PurpleSchool"] = "https://purpleschool.ru"
	fmt.Println(m)
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	delete(m, "Y")
	fmt.Println(m["Y"])
	fmt.Println(m)

	m2 := map[string]int{"a": 1, "b": 2}
	for key, value := range m2 {
		fmt.Println(key, value)
	}

	// 1. Посмотреть закладки
	// 2. Добавить закладку
	// 3. Удалить закладку
	// 4. Выход
	setOptions()
}
func setOptions() {
	pack := map[int]string{1: "123", 2: "321"}
outerLoop:
	for {
		request := getInputUser()
		switch request {
		case 1:
			viewPack(pack)
			fmt.Println(request)
		case 2:
			fmt.Println(request)
		case 3:
			deleteValueInPack(pack)
			fmt.Println(request)
		case 4:
			break outerLoop
		}

	}
}

func getInputUser() (request int) {
	for {
		fmt.Printf("___МЕНЮ___\n1. Посмотреть закладки\n2. Добавить закладку\n3. Удалить закладку\n4. Выход\nВведите операцию: ")
		fmt.Scan(&request)
		resultCheck, err := checkUserInputOperation(request)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	return request
}

func checkUserInputOperation(request int) (bool, error) {
	if request > 1 || request < 4 {
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func viewPack(pack map[int]string) {
	for key, value := range pack {
		fmt.Println(key, value)
	}
}

func deleteValueInPack(pack map[int]string) {
	delete(pack, 1)
}
