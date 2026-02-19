package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   int
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}
	count := 0
	for {
		cmd := 0
		fmt.Print(commands)

		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("Введите номер команды (число)\n")
			continue
		}

		switch cmd {
		case 1:
			if count >= size {
				fmt.Println("Штат полностью укомплектован! Максимум 512 сотрудников\n")
				continue
			}

			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s", &empl.Name)

			fmt.Println("Возраст:")
			_, err = fmt.Scan(&empl.Age)
			if err != nil || empl.Age <= 0 {
				fmt.Println("Возраст должен быть положительным числом. Добавление невозможно\n")
				clearInput()
				continue
			}

			fmt.Println("Позиция:")
			fmt.Scanf("%s", &empl.Position)

			fmt.Println("Зарплата:")
			_, err = fmt.Scan(&empl.Salary)
			if err != nil || empl.Salary < 0 {
				fmt.Println("Зарплата должна быть числом (>= 0). Добавление невозможно\n")
				clearInput()
				continue
			}

			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					count++
					fmt.Println("Сотрудник добавлен\n")
					break
				}
			}
		case 2:
			if count == 0 {
				fmt.Println("Список сотрудников пуст. Невозможно удалить\n")
				continue
			}

			var nameToDelete string
			fmt.Print("Введите имя сотрудника для удаления: ")
			fmt.Scan(&nameToDelete)

			found := false
			for i := 0; i < size; i++ {
				if empls[i] != nil && empls[i].Name == nameToDelete {
					empls[i] = nil
					count--
					found = true
					fmt.Printf("Сотрудник удален\n")
					break
				}
			}

			if !found {
				fmt.Printf("Сотрудник с таким именем не найден\n")
			}
		case 3:
			if count == 0 {
				fmt.Println("Список сотрудников пуст.\n")
				continue
			}

			fmt.Println("\n----Список сотрудников----")
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					e := empls[i]
					fmt.Printf("Имя: %s | Возраст: %d | Позиция: %s | Зарплата: %d\n", e.Name, e.Age, e.Position, e.Salary)
				}
			}
			fmt.Println("--------------------------")
		case 4:
			return

		default:
			fmt.Println("Ошибка. Введите число от 1 до 4\n")
		}
	}
}

func clearInput() {
	var c rune
	for {
		_, err := fmt.Scanf("%c", &c)
		if err != nil || c == '\n' {
			break
		}
	}
}
