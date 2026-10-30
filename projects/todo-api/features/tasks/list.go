package tasks

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type List struct {
	pool *pgxpool.Pool
}

func NewList(pool *pgxpool.Pool) *List {
	return &List{
		pool: pool,
	}
}

func (l *List) AddTask(task Task) error {
	_, err := l.pool.Exec(
		context.Background(),
		`INSERT INTO tasks (
			title, description, completed, created_at, complete_at
		) VALUES ($1, $2, $3, $4, $5)`,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompleteAt,
	)

	return err
}

func (l *List) GetTask(title string) (Task, error) {
	var task Task

	err := l.pool.QueryRow(
		context.Background(),
		`SELECT title, description, completed, created_at, complete_at
		 FROM tasks WHERE title = $1`,
		title,
	).Scan(
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.CompleteAt,
	)

	if err == pgx.ErrNoRows {
		return Task{}, ErrTaskNotFound
	}

	return task, err
}

func (l *List) ListTasks() (map[string]Task, error) {
	rows, err := l.pool.Query(
		context.Background(),
		`SELECT title, description, completed, created_at, complete_at
		 FROM tasks`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]Task)

	for rows.Next() {
		var task Task

		if err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompleteAt,
		); err != nil {
			return nil, err
		}

		tasks[task.Title] = task
	}

	return tasks, rows.Err()
}

func (l *List) ListUncompletedTasks() (map[string]Task, error) {
	tasks, err := l.ListTasks()
	if err != nil {
		return nil, err
	}

	uncompletedTasks := make(map[string]Task)

	for title, task := range tasks {
		if !task.Completed {
			uncompletedTasks[title] = task
		}
	}

	return uncompletedTasks, nil
}

func (l *List) CompleteTask(title string) (Task, error) {
	_, err := l.pool.Exec(
		context.Background(),
		`UPDATE tasks
		 SET completed = TRUE, complete_at = $2
		 WHERE title = $1`,
		title,
		time.Now(),
	)
	if err != nil {
		return Task{}, err
	}

	return l.GetTask(title)
}

func (l *List) UncompleteTask(title string) (Task, error) {
	_, err := l.pool.Exec(
		context.Background(),
		`UPDATE tasks
		 SET completed = FALSE, complete_at = NULL
		 WHERE title = $1`,
		title,
	)
	if err != nil {
		return Task{}, err
	}

	return l.GetTask(title)
}

func (l *List) DeleteTask(title string) error {
	result, err := l.pool.Exec(
		context.Background(),
		`DELETE FROM tasks WHERE title = $1`,
		title,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}