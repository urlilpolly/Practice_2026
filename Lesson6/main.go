package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Выбери пример:")
	fmt.Println("1 - nil channel: close(c)")
	fmt.Println("2 - nil channel: c <- val")
	fmt.Println("3 - nil channel: val := <-c")
	fmt.Println("4 - closed channel: close(c)")
	fmt.Println("5 - closed channel: c <- val")
	fmt.Println("6 - closed channel: val := <-c")
	fmt.Println("7 - обычный channel: close(c)")
	fmt.Println("8 - обычный channel: c <- val")
	fmt.Println("9 - обычный channel: val := <-c")

	var n int
	fmt.Print("Номер: ")
	fmt.Scan(&n)

	switch n {

	case 1:
		var c chan int

		fmt.Println("Пробуем закрыть nil channel")
		close(c)

	case 2:
		var c chan int

		fmt.Println("Пробуем отправить значение в nil channel")
		c <- 10

		fmt.Println("Эта строка не выполнится")

	case 3:
		var c chan int

		fmt.Println("Пробуем прочитать значение из nil channel")
		val := <-c

		fmt.Println("Получили:", val)

	case 4:
		c := make(chan int)
		close(c)

		fmt.Println("Пробуем закрыть уже закрытый channel")
		close(c)

	case 5:
		c := make(chan int)
		close(c)

		fmt.Println("Пробуем отправить значение в закрытый channel")
		c <- 10

	case 6:
		c := make(chan int)
		close(c)

		fmt.Println("Читаем из закрытого channel")
		val, ok := <-c

		fmt.Println("val =", val)
		fmt.Println("ok =", ok)

	case 7:
		c := make(chan int)

		fmt.Println("Закрываем обычный открытый channel")
		close(c)

		fmt.Println("Канал успешно закрыт")

	case 8:
		c := make(chan int)

		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Получатель готов принять значение")
			val := <-c
			fmt.Println("Получатель получил:", val)
		}()

		fmt.Println("Отправляем значение в обычный channel")
		c <- 10

		fmt.Println("Значение отправлено")

	case 9:
		c := make(chan int)

		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Отправляем значение в channel")
			c <- 20
		}()

		fmt.Println("Ждем значение из channel")
		val := <-c

		fmt.Println("Получили:", val)

	default:
		fmt.Println("Такого пункта нет")
	}
}
