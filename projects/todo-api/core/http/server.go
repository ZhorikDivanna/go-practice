package http

import (
	"Rest-API/features/tasks"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type HTTPServer struct {
	server *http.Server
	logger *zap.Logger
}

func NewHTTPServer(
	handlers *tasks.HTTPHandlers,
	logger *zap.Logger,
) *HTTPServer {
	s := &HTTPServer{
		logger: logger,
	}

	router := mux.NewRouter()

	router.Path("/tasks").
		Methods("POST").
		HandlerFunc(handlers.HandleCreateTask)

	router.Path("/tasks/{title}").
		Methods("GET").
		HandlerFunc(handlers.HandleGetTask)

	router.Path("/tasks").
		Methods("GET").
		Queries("completed", "false").
		HandlerFunc(handlers.HandleGetAllUncompletedTasks)

	router.Path("/tasks").
		Methods("GET").
		HandlerFunc(handlers.HandleGetAllTasks)

	router.Path("/tasks/{title}").
		Methods("PATCH").
		HandlerFunc(handlers.HandleCompleteTask)

	router.Path("/tasks/{title}").
		Methods("DELETE").
		HandlerFunc(handlers.HandleDeteleTask)

	s.server = &http.Server{
		Addr:    ":9091",
		Handler: s.logRequests(router),
	}

	return s
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (s *HTTPServer) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		s.logger.Info("Получен HTTP-запрос",
			zap.String("method", r.Method),
			zap.String("path", r.URL.RequestURI()),
			zap.Time("received_at", start),
		)

		response := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(response, r)

		fields := []zap.Field{
			zap.String("method", r.Method),
			zap.String("path", r.URL.RequestURI()),
			zap.Int("status", response.status),
			zap.Duration("duration", time.Since(start)),
			zap.Time("finished_at", time.Now()),
		}

		if response.status >= http.StatusInternalServerError {
			s.logger.Error("HTTP-запрос завершился с ошибкой", fields...)
		} else {
			s.logger.Info("HTTP-запрос завершён", fields...)
		}
	})
}

func (s *HTTPServer) StartServer() error {
	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *HTTPServer) Close() error {
	return s.server.Close()
}