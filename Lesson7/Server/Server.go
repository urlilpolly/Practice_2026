package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

const serverAddress = ":8080"

type Client struct {
	nick string
	conn net.Conn
	out  chan string
}

var (
	clients = make(map[string]*Client)
	mu      sync.RWMutex
)

func main() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	defer listener.Close()
	log.Println("TCP сервер запущен на", serverAddress)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка подключения:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		conn.Close()
		return
	}

	firstLine := strings.TrimSpace(scanner.Text())
	parts := strings.Fields(firstLine)

	if len(parts) != 2 || parts[0] != "NICK" {
		sendRaw(conn, "ERROR! первое сообщение должно быть: NICK <ник>")
		conn.Close()
		return
	}

	nick := strings.TrimSpace(parts[1])

	if nick == "" || strings.Contains(nick, " ") {
		sendRaw(conn, "ERROR! некорректный ник")
		conn.Close()
		return
	}

	client := &Client{
		nick: nick,
		conn: conn,
		out:  make(chan string, 32),
	}

	mu.Lock()
	if _, exists := clients[nick]; exists {
		mu.Unlock()
		sendRaw(conn, "ERROR! такой ник уже занят")
		conn.Close()
		return
	}

	clients[nick] = client
	mu.Unlock()

	log.Println("Подключился пользователь:", nick)

	go writeLoop(client)

	client.out <- "Теперь вы подключены как " + nick

	defer func() {
		mu.Lock()
		if clients[nick] == client {
			delete(clients, nick)
		}
		mu.Unlock()

		close(client.out)
		conn.Close()

		log.Println("Отключился пользователь:", nick)
	}()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		command, rest, ok := strings.Cut(line, " ")
		if !ok || command != "SEND" {
			client.out <- "ERROR! команда должна быть: SEND <ник_получателя> <сообщение>"
			continue
		}

		toNick, message, ok := strings.Cut(strings.TrimSpace(rest), " ")
		if !ok || strings.TrimSpace(toNick) == "" || strings.TrimSpace(message) == "" {
			client.out <- "ERROR! команда должна быть: SEND <ник_получателя> <сообщение>"
			continue
		}

		sendPrivateMessage(client, toNick, message)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Ошибка чтения от клиента", nick+":", err)
	}
}

func writeLoop(client *Client) {
	writer := bufio.NewWriter(client.conn)

	for message := range client.out {
		_, err := fmt.Fprintln(writer, message)
		if err != nil {
			return
		}

		err = writer.Flush()
		if err != nil {
			return
		}
	}
}

func sendPrivateMessage(from *Client, toNick string, message string) {
	mu.RLock()
	recipient, exists := clients[toNick]

	if exists {
		select {
		case recipient.out <- fmt.Sprintf("Сообщение от %s: %s", from.nick, message):
		default:
			from.out <- "ERROR! пользователь сейчас не может принять сообщение"
		}

		mu.RUnlock()
		return
	}

	mu.RUnlock()

	from.out <- "ERROR! пользователя с ником '" + toNick + "' нет в системе"
}

func sendRaw(conn net.Conn, message string) {
	writer := bufio.NewWriter(conn)
	fmt.Fprintln(writer, message)
	writer.Flush()
}
