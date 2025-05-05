package main

import (
	"net/http"

	"github.com/IbadT/golang-the-way-of-the-warrior.git/handlers"
)

func main() {
	http.HandleFunc("/tasks", handlers.GetTasks)
	http.HandleFunc("/task-id", handlers.GetTaskById)

	http.HandleFunc("/create/task", handlers.CreateTask)

	http.HandleFunc("/update/task", handlers.UpdateTaskById)
	http.HandleFunc("/update/task/completed", handlers.UpdateTaskCompletedById)

	http.HandleFunc("/delete/task-id", handlers.DeleteTaskById)

	http.HandleFunc("/", handlers.GetUsers)
	http.HandleFunc("/users", handlers.GetUsersByName)
	// http.HandleFunc("/:id", handlers.GetUserById)
	http.HandleFunc("/user", handlers.CreateUser)
	// http.HandleFunc("/patch/user/:id", handlers.UpdateUserById)
	// http.HandleFunc("/delete/user/:id", handlers.DeleteUserById)

	http.ListenAndServe("localhost:8080", nil)
}

// go list -m

//

// - **Создайте структуру `Order` с полями:**
//     - `ID` (уникальный номер заказа),
//     - `Quantity` (количество единиц товара, которое заказывают).
// - **Напишите функцию `supplier(ordersChan chan<- Order)` — это ваш «поставщик»:**
//     - Эта функция (будущая горутина) в цикле отправляет несколько (скажем, 5) заказов в канал `ordersChan`, указывая разные значения `ID` и `Quantity`.
//     - По завершении отправки всех заказов **закрывайте** канал `ordersChan` (операция `close(ordersChan)`), сообщая, что новых заказов больше не будет.
// - **Напишите функцию `warehouse(ordersChan <-chan Order, resultsChan chan<- string)` — это ваш «склад»:**
//     - Убедимся, что это **просто функция**, но при запуске через `go warehouse(...)` она станет «горутиной-складом», работающей параллельно.
//     - Внутри функции примите условный «остаток товара» (например, 10 штук).
//     - С помощью цикла `for order := range ordersChan` считайте заказы из канала:
//         - Если товара достаточно, вычитайте `order.Quantity` из остатка и пишите в `resultsChan` строку вида `"Order <ID> processed, left <остаток> items"`.
//         - Если товара не хватает, пишите в `resultsChan` строку вида `"Order <ID> rejected: not enough items"`.
//     - Когда канал `ordersChan` закрывается, цикл `for range` завершится. После этого **закройте** `resultsChan`,
// 		чтобы «сказать» читающей стороне, что сообщений больше не будет.
// - **В функции `main()`:**
//     - Объявите два канала: `ordersChan := make(chan Order)` и `resultsChan := make(chan string)`.
//     - **Запустите** «поставщика» (функцию `supplier`) через `go supplier(ordersChan)`.
//     - **Запустите** «склад» (функцию `warehouse`) через `go warehouse(ordersChan, resultsChan)`.
//     - В `main()` сделайте цикл `for msg := range resultsChan`, где `msg` — строки, приходящие от «склада». Выводите их в консоль.
//     - Когда канал `resultsChan` закроется, цикл завершится, и программа может завершать работу.
// - **(Дополнительно) Добавьте задержки:**
//     - В `supplier` после отправки каждого заказа — `time.Sleep(300 * time.Millisecond)`.
//     - В `warehouse` перед обработкой заказа — `time.Sleep(500 * time.Millisecond)`.

//

// type Person struct {
// 	ID   int
// 	Name string
// }

// func main() {
// 	chanLen := 2
// 	personsChan := make(chan Person, chanLen)

// 	persons := []Person{
// 		{ID: 1, Name: "John"},
// 		{ID: 2, Name: "Eduard"},
// 		{ID: 3, Name: "Brew"},
// 		{ID: 4, Name: "Matwey"},
// 		{ID: 5, Name: "Andrew"},
// 	}

// 	for index, person := range persons {
// 		if index+1 <= chanLen {
// 			personsChan <- person
// 		} else {
// 			break
// 		}
// 	}

// 	close(personsChan)

// 	fmt.Println(<-personsChan)
// 	fmt.Println(<-personsChan)
// }

// func main() {
// 	mailboxOne := make(chan int)
// 	mailboxTwo := make(chan int)

// 	// Запускаем горутину для отправки данных
// 	go func() {
// 		for value := range 20 {
// 			if value%2 == 0 {
// 				mailboxTwo <- value
// 			} else {
// 				mailboxOne <- value
// 			}
// 			time.Sleep(time.Millisecond * 100) // Имитация задержки
// 		}
// 		close(mailboxOne)
// 		close(mailboxTwo)
// 	}()

// 	// Читаем данные из каналов, пока они не закрыты
// 	for mailboxOne != nil || mailboxTwo != nil {
// 		select {
// 		case msg, ok := <-mailboxOne:
// 			if ok {
// 				fmt.Println("📬 Получено из mailboxOne:", msg)
// 			} else {
// 				mailboxOne = nil // Очищаем канал после закрытия
// 			}
// 		case msg, ok := <-mailboxTwo:
// 			if ok {
// 				fmt.Println("📬 Получено из mailboxTwo:", msg)
// 			} else {
// 				mailboxTwo = nil // Очищаем канал после закрытия
// 			}
// 		default:
// 			time.Sleep(100 * time.Millisecond) // Ждём новые данные
// 		}
// 	}

// 	fmt.Println("✅ Все данные получены, завершаем программу.")
// }

// // - **Создайте структуру `Order` с полями:**
// //     - `ID` (уникальный номер заказа),
// //     - `Quantity` (количество единиц товара, которое заказывают).
// // - **Напишите функцию `supplier(ordersChan chan<- Order)` — это ваш «поставщик»:**
// //     - Эта функция (будущая горутина) в цикле отправляет несколько (скажем, 5) заказов в канал `ordersChan`, указывая разные значения `ID` и `Quantity`.
// //     - По завершении отправки всех заказов **закрывайте** канал `ordersChan` (операция `close(ordersChan)`), сообщая, что новых заказов больше не будет.
// // - **Напишите функцию `warehouse(ordersChan <-chan Order, resultsChan chan<- string)` — это ваш «склад»:**
// //     - Убедимся, что это **просто функция**, но при запуске через `go warehouse(...)` она станет «горутиной-складом», работающей параллельно.
// //     - Внутри функции примите условный «остаток товара» (например, 10 штук).
// //     - С помощью цикла `for order := range ordersChan` считайте заказы из канала:
// //         - Если товара достаточно, вычитайте `order.Quantity` из остатка и пишите в `resultsChan` строку вида `"Order <ID> processed, left <остаток> items"`.
// //         - Если товара не хватает, пишите в `resultsChan` строку вида `"Order <ID> rejected: not enough items"`.
// //     - Когда канал `ordersChan` закрывается, цикл `for range` завершится. После этого **закройте** `resultsChan`,
// // чтобы «сказать» читающей стороне, что сообщений больше не будет.
// // - **В функции `main()`:**
// //     - Объявите два канала: `ordersChan := make(chan Order)` и `resultsChan := make(chan string)`.
// //     - **Запустите** «поставщика» (функцию `supplier`) через `go supplier(ordersChan)`.
// //     - **Запустите** «склад» (функцию `warehouse`) через `go warehouse(ordersChan, resultsChan)`.
// //     - В `main()` сделайте цикл `for msg := range resultsChan`, где `msg` — строки, приходящие от «склада». Выводите их в консоль.
// //     - Когда канал `resultsChan` закроется, цикл завершится, и программа может завершать работу.
// // - **(Дополнительно) Добавьте задержки:**
// //     - В `supplier` после отправки каждого заказа — `time.Sleep(300 * time.Millisecond)`.
// //     - В `warehouse` перед обработкой заказа — `time.Sleep(500 * time.Millisecond)`.

// type Order struct {
// 	ID       int
// 	Quantity int
// }

// func supplier(ordersChan chan<- Order) {
// 	for i := 1; i <= 5; i++ {
// 		ordersChan <- Order{ID: i, Quantity: i * 2}
// 		time.Sleep(300 * time.Millisecond)
// 	}
// 	close(ordersChan)
// }

// func warehouse(ordersChan <-chan Order, resultsChan chan<- string) {
// 	stock := 20
// 	for order := range ordersChan {
// 		time.Sleep(500 * time.Millisecond)
// 		if stock >= order.Quantity {
// 			stock -= order.Quantity
// 			resultsChan <- fmt.Sprintf("Order %d processed, left %d items", order.ID, stock)
// 		} else {
// 			resultsChan <- fmt.Sprintf("Order %d rejected: not enough items", order.ID)
// 		}
// 	}
// 	close(resultsChan)
// }

// func main() {
// 	ordersChan := make(chan Order)
// 	resultsChan := make(chan string)

// 	go supplier(ordersChan)
// 	go warehouse(ordersChan, resultsChan)

// 	for msg := range resultsChan {
// 		fmt.Println(msg)
// 	}
// }
