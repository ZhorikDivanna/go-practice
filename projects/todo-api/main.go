package main

import (
	"Rest-API/core/database"
	httpserver "Rest-API/core/http"
	"Rest-API/core/logger"
	"Rest-API/features/tasks"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
	)
	defer stop()

	conn, err := database.NewConnection(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("Не удалось создать пул подключения", zap.Error(err))
		return
	}
	defer conn.Close()

	if err := conn.Ping(ctx); err != nil {
		log.Error("Не удалось подключиться к PostgreSQL", zap.Error(err))
		return
	}

	log.Info("Подключение к PostgreSQL установлено")

	todoList := tasks.NewList(conn)
	service := tasks.NewService(todoList)
	httpHandlers := tasks.NewHTTPHaldlers(service)
	httpServer := httpserver.NewHTTPServer(httpHandlers, log)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- httpServer.StartServer()
	}()

	log.Info("Запуск HTTP-сервера",
		zap.String("address", ":9091"),
		zap.Int("pid", os.Getpid()),
	)

	select {
	case err := <-serverErrors:
		if err != nil {
			log.Error("Ошибка HTTP-сервера", zap.Error(err))
		}
		return

	case <-ctx.Done():
		log.Info("Получен SIGTERM, останавливаем HTTP-сервер")
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error("Ошибка завершения HTTP-сервера", zap.Error(err))

		if err := httpServer.Close(); err != nil {
			log.Error("Ошибка принудительной остановки", zap.Error(err))
		}
		return
	}

	log.Info("HTTP-сервер корректно остановлен")
}