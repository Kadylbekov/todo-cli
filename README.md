# 📝 ToDo CLI на Go

Простой консольный менеджер задач, написанный на языке Go.  
Позволяет создавать, просматривать, отмечать как выполненные и удалять задачи.  
Все данные хранятся в локальном JSON-файле.

---

## 📦 Возможности

- `add <task>` — добавить новую задачу
- `list` — показать все задачи
- `done <id>` — отметить задачу как выполненную
- `delete <id>` — удалить задачу по ID
- `stats` — показать статистику задач

---

## ▶ Пример использования

```bash
go run main.go add "Купить хлеб"
go run main.go list
go run main.go done 1
go run main.go delete 1
go run main.go stats
