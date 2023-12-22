package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

const (
	todoFile = ".todo.json"
)

func main() {
	// Флаг "-add" позволяет добавить новую задачу в список.
	add := flag.Bool("add", false, "добавить новую задачу")

	// Флаг "-complete" позволяет отметить задачу как выполненную, используя индекс задачи.
	complete := flag.Int("complete", 0, "отметить задачу как выполненную")
	// Флаг "-del" позволяет удалить задачу
	del := flag.Int("del", 0, "удалить задачу")
	// Флаг "-list" позволяет увидить to see all todos
	list := flag.Bool("list", false, "список всех задач")

	// Функция flag.Parse() анализирует аргументы командной строки.
	// Должна быть вызвана после определения всех флагов и перед обращением к флагам программы.
	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// Выполнение команды в зависимости от указанных флагов.
	switch {
	case *add:
		//// Добавление новой задачи.
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		//простой пример
		//todos.Add("sample todo")
		//err := todos.Store(todoFile)
		//if err != nil {
		//	fmt.Fprintln(os.Stderr, err.Error())
		//	os.Exit(1)
		//}

	case *complete > 0:
		// Отметить задачу как выполненную.
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		// Отметить задачу как выполненную.
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()

	default:
		// Ошибка: неправильная команда.
		fmt.Fprintln(os.Stderr, "неправильная команда")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	//Если введено одно или более аргументов, возвращаем их склеенные через пробел
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}
	//Если аргументов нет, создаем объект scanner для считывания введенного пользователем текста
	scanner := bufio.NewScanner(r)
	// Считываем текст и проверяем на наличие ошибок
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	// Проверяем, не является ли введенный текст пустым, и возвращаем ошибку, если это так
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("задача не может быть пустым")
	}
	// Возвращаем считанный текст без ошибок
	return text, nil
}
