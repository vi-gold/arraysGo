package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// var transactions [3]int
	// transactions := [5]int{1, 2, 3, 4, 5}
	// banks := [2]string{}

	// fmt.Println(transactions[1])
	// banks[0] = "Тинькофф"
	// fmt.Println(banks)

	// patrial := transactions[1:]
	// fmt.Println(patrial)
	// ------------------------------------
	// transactions := [6]int{1, 2, 3, 4, 5, 6}
	// transactionsPartial := transactions[1:5]
	// transactionsNewPartial := transactionsPartial[:1]
	// // transactionsPartial[0] = 30
	// transactionsNewPartial[0] = 30
	// transactionsNewPartial = transactionsNewPartial[0:4]

	// fmt.Println(transactions)
	// fmt.Println(transactionsNewPartial)
	// fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	// fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))
	// --------------------------------------

	// transactions := []int{0, 20, 35}
	// //slice := transactions[:2]
	// //slice[0] = 100
	// temp := transactions
	// transactions = append(transactions, 100)
	// fmt.Println(temp)
	// fmt.Println(transactions)

	//fmt.Println(slice)
	// ----------------------------------------------
	// temp1 := 0
	// maxTr := 0
	// tr := []float64{}
	// fmt.Printf("Введите колличество транзакций: ")
	// fmt.Scan(&maxTr)
	// for i := 0; i < maxTr; i++ {
	// 	fmt.Printf("Введите сумму транзакции: ")
	// 	fmt.Scan(&temp1)
	// 	tr = append(tr, float64(temp1))
	// }
	// fmt.Println(tr)
	// ------------------------------------------------
	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, tr2...)
	// fmt.Println(tr1)

	// for index, value := range tr1 {
	// 	fmt.Println(index, value)
	// }

	// tr := make([]string, 0, 2)
	// fmt.Println(len(tr), cap(tr))
	// tr = append(tr, "1")
	// fmt.Println(len(tr), cap(tr))
	// tr = append(tr, "2")
	// fmt.Println(len(tr), cap(tr))
	// tr = append(tr, "3")
	// fmt.Println(len(tr), cap(tr))
	// fmt.Println(tr)
	// tr[5] = "1"
	// fmt.Println(tr)

	// GENERAL CODE

	// 	transactions := []float64{}
	// 	for {
	// 		transaction := scanTransaction()
	// 		if transaction == 0 {
	// 			break
	// 		}
	// 		transactions = append(transactions, transaction)
	// 	}
	// 	balance := calculateBalance(transactions)
	// 	fmt.Printf("Ваш баланс: %.2f", balance)
	// }

	// func scanTransaction() (transaction float64) {
	// 	fmt.Print("Введите транзакцию (n для выхода): ")
	// 	fmt.Scan(&transaction)
	// 	return transaction
	// }

	// func calculateBalance(transactions []float64) (balance float64) {
	// 	for _, value := range transactions {
	// 		balance += value
	// 	}
	// 	return balance

	// END OF GENERAL CODE

	// TEST array homework
	operation, listNum := getInputUser()
	fmt.Println("Operation - ", operation)
	fmt.Println("Numbers - ", listNum)
	arrayNum := convertStringToArray(listNum)
	resultCalc := calculateOperation(arrayNum, operation)
	fmt.Printf("Операция %s с числами %s равна %.2f", operation, listNum, resultCalc)
}

func getInputUser() (operation, listNum string) {
	for {
		fmt.Print("Введите операцию (avg, sum, med): ")
		fmt.Scan(&operation)
		resultCheck, err := checkUserInputOperation(operation)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&listNum)
	return operation, listNum

}

func checkUserInputOperation(operation string) (bool, error) {
	if operation == "avg" || operation == "sum" || operation == "med" {
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func convertStringToArray(listNum string) []float64 {
	arrayStrings := strings.Split(listNum, ",")
	arrayNum := make([]float64, 0, len(arrayStrings))
	for _, value := range arrayStrings {
		convertedNum, err := strconv.ParseFloat(value, 64)
		if err == nil {
			arrayNum = append(arrayNum, convertedNum)
		}
	}
	return arrayNum
}

func calculateOperation(arrayNum []float64, operation string) (result float64) {
	if operation == "avg" {
		for _, value := range arrayNum {
			result += value
		}
		result = result / float64(len(arrayNum))
	}
	if operation == "sum" {
		for _, value := range arrayNum {
			result += value
		}
	}
	if operation == "med" {
		sort.Float64s(arrayNum)
		if len(arrayNum)%2 == 1 {
			result = arrayNum[len(arrayNum)/2]
		} else {
			median1 := arrayNum[len(arrayNum)-1]
			median2 := arrayNum[len(arrayNum)]
			result = (median1 + median2) / 2
		}
	}
	return result
}
