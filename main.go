package main

import "fmt"

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
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

func scanTransaction() (transaction float64) {
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}
