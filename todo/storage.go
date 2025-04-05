package todo

import (
	"encoding/json"
	"os"
)

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ") //Превращаем список задач в красиво отформатированный JSON (отступ 2 пробела).
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644) //Записываем JSON в файл. Права доступа 0644 означают: владелец может читать и писать, остальные — только читать
}
func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename) //Читаем весь файл с задачами в переменную data.
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err //Если файла ещё нет — это не ошибка, просто вернём пустой список. Если другая ошибка — возвращаем её.
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks) //Создаём пустой список задач. Распаковываем JSON из data в tasks.
	return tasks, err
}
