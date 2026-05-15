package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const serverAddress = "localhost:8080"

func main() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите уникальный ник: ")
	nick := readLine(consoleReader)

	if nick == "" || strings.Contains(nick, " ") {
		fmt.Println("Ник не должен быть пустым и не должен содержать пробелы")
		return
	}

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "NICK %s\n", nick)

	serverScanner := bufio.NewScanner(conn)

	if !serverScanner.Scan() {
		fmt.Println("Сервер закрыл соединение")
		return
	}

	response := serverScanner.Text()
	fmt.Println(response)

	if strings.HasPrefix(response, "ERR") {
		return
	}

	done := make(chan struct{})

	go func() {
		for serverScanner.Scan() {
			message := serverScanner.Text()
			fmt.Println()
			fmt.Println(message)
		}

		if err := serverScanner.Err(); err != nil {
			fmt.Println("Ошибка соединения с сервером:", err)
		}

		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("Соединение с сервером закрыто")
			return
		default:
		}

		fmt.Print("Введите ник получателя или /exit: ")
		toNick := readLine(consoleReader)

		if toNick == "/exit" {
			fmt.Println("Вы вышли из чата")
			return
		}

		if toNick == "" {
			continue
		}

		fmt.Print("Введите сообщение: ")
		message := readLine(consoleReader)

		if message == "" {
			continue
		}

		_, err := fmt.Fprintf(conn, "SEND %s %s\n", toNick, message)
		if err != nil {
			fmt.Println("Не удалось отправить сообщение:", err)
			return
		}
	}
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
