# Employee API

Учебный HTTP API для работы с сотрудниками на Go и PostgreSQL.

## Возможности

- Создание сотрудника.
- Получение списка сотрудников.
- Удаление сотрудника по ID.
- Запуск приложения и БД через Docker Compose.
- Сохранение данных PostgreSQL в Docker volume.

## Технологии

Go, net/http, encoding/json, PostgreSQL, pgxpool, Docker Compose.

## Запуск

Нужен установленный и запущенный Docker с поддержкой Compose.

Из папки проекта создайте файл настроек:

```bash
cp .env.example .env
```

В `.env` замените `CHANGE_ME` своим локальным паролем PostgreSQL.
Для простоты используйте буквы и цифры: пароль подставляется в URL подключения.

Если база уже существует в Docker volume, укажите её текущий пароль.
Изменение `.env` не меняет пароль существующей БД.

Запустите:

```bash
docker compose up --build
```

Приложение доступно по адресу:

```text
http://localhost:8080
```

Compose ожидает готовности PostgreSQL перед запуском приложения.
Таблица `employees` создаётся автоматически при старте.

## API

| Метод | Маршрут | Описание |
|---|---|---|
| POST | `/employees` | Создать сотрудника |
| GET | `/employees` | Получить список |
| DELETE | `/employees?id=1` | Удалить по ID |

### Создание

Пример для Bash / WSL:

```bash
curl -i -X POST http://localhost:8080/employees \
  -H 'Content-Type: application/json' \
  -d '{"ID":1,"FullName":"Иван Иванов","Position":"Go Intern"}'
```

Успешный ответ — `201 Created`.

В текущей реализации ID задаётся клиентом и должен быть уникальным.

### Получение списка

```bash
curl -i http://localhost:8080/employees
```

Успешный ответ — `200 OK`:

```json
[
  {
    "ID": 1,
    "FullName": "Иван Иванов",
    "Position": "Go Intern"
  }
]
```

Если сотрудников нет, возвращается `[]`.
Порядок записей не гарантирован.

### Удаление

```bash
curl -i -X DELETE 'http://localhost:8080/employees?id=1'
```

Успешный ответ — `204 No Content`.
Этот же статус возвращается, если записи с указанным ID нет.

## Структура

```text
main.go                  — запуск и маршруты
Employee_Store_CRUD/      — HTTP-обработчики и SQL
Module_employee/         — структура сотрудника
simple_connection/       — подключение через pgxpool
Dockerfile               — контейнер приложения
docker-compose.yaml      — приложение, БД и volume
.env.example             — пример настроек
```

## Проверка кода

```bash
go fmt ./...
go build ./...
go vet ./...
```

## Остановка

```bash
docker compose down
```

Данные сохраняются в именованном томе между запусками.

## Ограничения

Пока не реализованы обновление сотрудника, пагинация, валидация
обязательных полей и автоматические тесты.

