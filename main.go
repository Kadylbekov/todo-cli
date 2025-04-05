package main

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli/todo"
)

const filename = "tasks.json" // Имя файла для хранения задач

func main() {
	args := os.Args    // Чтение аргументов командной строки
	if len(args) < 2 { // Если аргументов недостаточно (меньше 2)
		// Показываем подсказку по использованию программы
		fmt.Println("Usage:")
		fmt.Println("  add <task>       — добавить задачу")
		fmt.Println("  list             — показать все задачи")
		fmt.Println("  done <id>        — отметить задачу как выполненную")
		fmt.Println("  delete <id>      — удалить задачу")
		return // Прерываем выполнение программы, если команда не задана
	}

	command := args[1] // Вторая строка — это команда (add, list, done, delete)

	switch command { // В зависимости от команды выполняем разные действия
	case "add":
		if len(args) < 3 { // Если задачи нет в аргументах
			fmt.Println("Пожалуйста, укажи название задачи.")
			return // Завершаем программу, если не указано название задачи
		}

		title := args[2]                       // Название задачи — третий аргумент в командной строке
		tasks, err := todo.LoadTasks(filename) // Загружаем текущие задачи из файла
		if err != nil {                        // Если не удалось загрузить задачи
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		// Создаём новую задачу с уникальным ID
		newTask := todo.Task{
			ID:    len(tasks) + 1, // ID задачи = текущая длина списка + 1
			Title: title,          // Название задачи — то, что ввёл пользователь
		}

		tasks = append(tasks, newTask) // Добавляем новую задачу в список задач

		// Сохраняем обновлённый список задач в файл
		err = todo.SaveTasks(filename, tasks)
		if err != nil { // Если не удалось сохранить задачи
			fmt.Println("Ошибка сохранения:", err)
			return
		}

		// Выводим сообщение, что задача успешно добавлена
		fmt.Println("Задача добавлена:", title)

	case "list":
		tasks, err := todo.LoadTasks(filename) // Загружаем все задачи
		if err != nil {                        // Если не удалось загрузить задачи
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		if len(tasks) == 0 { // Если задач нет
			fmt.Println("Задачи пока нет.")
			returng
		}

		// Выводим все задачи с их ID и статусом
		for _, task := range tasks {
			status := "[ ]"     // По умолчанию задача не выполнена
			if task.Completed { // Если задача выполнена
				status = "[x]" // Отображаем статус как выполненный
			}
			// Выводим ID задачи, её статус и название
			fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
		}

	case "done":
		if len(args) < 3 { // Если не указан ID задачи
			fmt.Println("Пожалуйста, укажи ID задачи, чтобы отметить её выполненной.")
			return // Завершаем программу
		}

		// Преобразуем третий аргумент (ID задачи) в число
		id, err := strconv.Atoi(args[2])
		if err != nil { // Если ID не является числом
			fmt.Println("Ошибка: ID задачи должно быть числом.")
			return
		}

		tasks, err := todo.LoadTasks(filename) // Загружаем все задачи
		if err != nil {                        // Если не удалось загрузить задачи
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		// Ищем задачу с указанным ID
		for i, task := range tasks {
			if task.ID == id { // Если нашли задачу с таким ID
				tasks[i].Completed = true // Отмечаем задачу как выполненную
				break
			}
		}

		// Сохраняем обновлённый список задач
		err = todo.SaveTasks(filename, tasks)
		if err != nil { // Если не удалось сохранить изменения
			fmt.Println("Ошибка сохранения:", err)
			return
		}

		fmt.Println("Задача отмечена как выполненная.")

	case "delete":
		if len(args) < 3 { // Если не указан ID задачи для удаления
			fmt.Println("Пожалуйста, укажи ID задачи для удаления.")
			return
		}

		// Преобразуем ID задачи в число
		id, err := strconv.Atoi(args[2])
		if err != nil { // Если ID не числовой
			fmt.Println("Ошибка: ID задачи должно быть числом.")
			return
		}

		tasks, err := todo.LoadTasks(filename) // Загружаем все задачи
		if err != nil {                        // Если возникла ошибка при загрузке
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		// Создаём новый список задач без удалённой
		var newTasks []todo.Task
		for _, task := range tasks {
			if task.ID != id { // Если ID задачи не совпадает с удаляемым
				newTasks = append(newTasks, task) // Добавляем её в новый список
			}
		}

		// Сохраняем обновлённый список задач
		err = todo.SaveTasks(filename, newTasks)
		if err != nil { // Если не удалось сохранить
			fmt.Println("Ошибка сохранения:", err)
			return
		}

		// Сообщаем, что задача удалена
		fmt.Println("Задача удалена.")

	default:
		// Если команда неизвестна
		fmt.Println("Неизвестная команда:", command)
	}
}
