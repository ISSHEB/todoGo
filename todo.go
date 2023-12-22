package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	// Создание нового экземпляра структуры item
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	// Добавление нового экземпляра структуры item в массив t
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	// Создаем указатель на массив задач.
	ls := *t
	// Проверяем, является ли индекс корректным.
	if index <= 0 || index > len(ls) {
		// Если индекс некорректный, возвращаем ошибку.
		return errors.New("invalid index")
	}

	// Для корректного индекса, устанавливаем дату завершения задачи.
	ls[index-1].CompletedAt = time.Now()

	// Устанавливаем значение поля "Done" равное true.
	ls[index-1].Done = true

	// Возвращаем nil, т.к. операция выполнена успешно.
	return nil
}

func (t *Todos) Delete(index int) error {
	// Создаем указатель на массив задач.
	ls := *t
	// Проверяем, является ли индекс корректным.
	if index <= 0 || index > len(ls) {
		// Если индекс некорректный, возвращаем ошибку.
		return errors.New("invalid index")
	}

	// Для корректного индекса, удаляем задачу из массива.
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	// Открываем файл и считываем его содержимое.
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// Если файл не существует, возвращаем nil.
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		// Если файл существует, но не может быть прочитан, возвращаем ошибку.
		return err
	}

	// Проверяем, пуст ли файл.
	if len(file) == 0 {
		// Если файл пустой, возвращаем ошибку.
		return err
	}

	// Расшифровываем данные из файла и сохраняем их в структуре задач.
	err = json.Unmarshal(file, t)
	if err != nil {
		// Если расшифровка не удалась, возвращаем ошибку.
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	// Конвертируем структуру задач в json-формат.
	data, err := json.Marshal(t)
	if err != nil {
		// Если конвертация не удалась, возвращаем ошибку.
		return err
	}

	// Записываем конвертированные данные в файл.
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		// Если запись в файл не удалась, возвращаем ошибку.
		return err
	}

	// Если все шаги выполнены успешно, возвращаем nil.
	return nil
}

// Метод Print для отображения списка todos на экране
func (t *Todos) Print() {
	// Создаем новый экземпляр таблицы
	table := simpletable.New()

	//Затем настраивается шапка таблицы, добавляя заголовки для каждого столбца
	table.Header = &simpletable.Header{
		// Задаем данные для строки
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "#Task"},
			{Align: simpletable.AlignCenter, Text: "#Done?"},
			{Align: simpletable.AlignCenter, Text: "#CreateAt"},
			{Align: simpletable.AlignCenter, Text: "#CompletedAt"},
		},
	}
	//в README есть обяснения
	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")

		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})

	}
	//в footer таблицы добавляется одна ячейка, содержащая текст
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("you have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()

	//простой вариант
	// Используем цикл for с range для перебора каждого элемента в списке todos
	//for i, item := range *t {
	//	// Используем fmt.Printf для печати индекса и задачи каждого элемента в списке todos
	//	fmt.Printf("%d - %s\n", i, item.Task)
	//}
}

func (t *Todos) CountPending() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}
