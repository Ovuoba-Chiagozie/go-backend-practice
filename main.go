package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func calculateAge(birthYear int) int {
	return 2026 - birthYear
}

func isAdult(age int) bool {
	return age >= 18
}
func main() {
	// name := "Chiagozie"
	// birthYear := 2009
	// isActive := true

	// age := calculateAge(birthYear)
	// isAdult := isAdult(age)

	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Active:", isActive)
	// fmt.Println("Is an adult:", isAdult)

	// array := make([]int, 2, 5)

	// array[0] = 4
	// nonpointer := append(array, 2, 4, 5, 6, 9, 90)
	// fmt.Println(nonpointer, array)

	// main2()
	// main3()
	// main4()
	// main5()

	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	printLetters()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	printNumbers()
	// }()

	// wg.Wait()

	// fmt.Println("It is Finished")

	var wg sync.WaitGroup

	wg.Add(1)

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

	go func() {
		defer wg.Done()
		doWork(ctx)
	}()

	// time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("Main function finished")
}

func doWork(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled")
			return
		case <-ticker.C:
			fmt.Println("Processing task..")
		}
	}
}

func main2() {

	users := make(map[string]string)

	users["John"] = "admin"
	users["Mary"] = "editor"
	users["Peter"] = "user"

	users["Mary"] = "admin"

	value, exists := users["Peter"]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("U know u dont have sense this guy")
	}

	delete(users, "John")

	for key, value := range users {
		fmt.Println(key, value)
	}
	// Add these users:
	// John -> admin
	// Mary -> editor
	// Peter -> user

	// Update Mary's role to admin

	// Check if Peter exists and print his role

	// Delete John

	// Loop through the remaining users and print their names and roles
}

func main3() {

	account := BankAccount{
		AccountHolder: "Chiagozie",
		AccountNumber: "123456789",
		Balance:       9000000.99,
	}

	amount := 6700.00
	account.Deposit(&amount)
	err := account.Withdraw(amount)

	if errors.Is(err, ErrInsuffienctFunds) {
		fmt.Println(err)
	}

	fmt.Println("The current balance is", account.getBalance())
	fmt.Println("It has sufficent balance for this amount", account.HasSufficentFunds(amount))
	fmt.Println("Current amount", amount)

}

type BankAccount struct {
	AccountHolder string
	AccountNumber string
	Balance       float64
}

func (b *BankAccount) Deposit(amount *float64) float64 {
	b.Balance = b.Balance + *amount
	return b.Balance
}

var ErrInsuffienctFunds = errors.New("Insufficient funds")

func (b *BankAccount) Withdraw(amount float64) error {

	if amount <= 0 {
		return errors.New("amount cannot be negative")
	}

	if amount <= b.Balance {
		b.Balance = b.Balance - amount
		return nil
	}
	return ErrInsuffienctFunds
}

func (b BankAccount) getBalance() float64 {
	return b.Balance
}

func (b BankAccount) HasSufficentFunds(amount float64) bool {
	return b.Balance >= amount
}

type NotificationSender interface {
	Send(message string) error
}

type EmailSender struct {
}

func (e EmailSender) Send(message string) error {
	if message == "" {
		return errors.New("Message cannot be empty")
	}
	fmt.Println("This message:", message, "was sent")
	return nil
}

type SmsSender struct {
}

func (s SmsSender) Send(message string) error {
	if message == "" {
		return errors.New("Message cannot be empty")
	}
	fmt.Println("This message:", message, "was fing sent")
	return nil
}

func sendNotification(notificationProcessor NotificationSender) {
	err := notificationProcessor.Send("")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main4() {
	sms := SmsSender{}
	email := EmailSender{}

	sendNotification(sms)
	sendNotification(email)

}

func ProcessPayment(amount int) {
	defer fmt.Println("Payment processing finished")
	fmt.Println("Processing payment...")

	if amount <= 0 {
		panic("So u no get shii guy")
	}

	fmt.Println("Payment fucking SUccessful")
}

func main5() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/tasks", TaskHandler)
	http.HandleFunc("GET /tasks/{id}", getTaskHandler)
	http.HandleFunc("PUT /tasks/{id}", updateTaskHandler)
	http.HandleFunc("DELETE /tasks/{id}", deleteTaskHandler)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server running bitch")

}

func printNumbers() {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {

		fmt.Println(letter)
	}
}
