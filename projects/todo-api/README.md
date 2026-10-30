# ToDo REST API

Учебное приложение на Go для управления задачами.

## Возможности

- Создание, просмотр и удаление задач.
- Отметка выполнения и её отмена.
- Фильтрация невыполненных задач.
- Хранение данных в PostgreSQL.
- Логирование HTTP-запросов и завершение по SIGTERM.

## Стек

Go, net/http, Gorilla Mux, PostgreSQL, pgxpool, zap,
golang-migrate, Docker Compose, Makefile.

Код разделён на HTTP-обработчики, сервис и репозиторий.

## Запуск

Из папки проекта:

1. Скопировать `.env.example` в `.env`.
2. Запустить PostgreSQL: `make postgres-up`.
3. В другом терминале применить миграции: `make migrate-up`.
4. Запустить приложение: `make run-local`.

API: `http://localhost:9091/tasks`.

Для запуска приложения в Docker вместо `make run-local`
используется `make run-docker`.

## О проекте

Выполнен в рамках обучения с использованием учебных материалов. Работа API проверялась вручную через curl и pgAdmin.