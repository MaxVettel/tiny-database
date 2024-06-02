NoSQL key-value база данных.
После запуска команды вводятся в bash стиле.

Запуск приложения:
```bash
go run main.go
```
Пример команд:
```bash
set Hello
get edcfa0ad-4f51-4468-b86a-099b9fc76fd9
stop
```

Типы комманд:
- stop - остановка приложения
- get {id} - запрос value по id записи
- getall - запрос всех записей
- set {values ...} - сохранение новых записей
- delete {id} - удаление записи по id
