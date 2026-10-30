package tasks

type Service struct {
	repository *List
}

func NewService(repository *List) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateTask(title, description string) (Task, error) {
	task := NewTask(title, description)

	if err := s.repository.AddTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *Service) GetTask(title string) (Task, error) {
	return s.repository.GetTask(title)
}

func (s *Service) ListTasks() (map[string]Task, error) {
	return s.repository.ListTasks()
}

func (s *Service) ListUncompletedTasks() (map[string]Task, error) {
	return s.repository.ListUncompletedTasks()
}

func (s *Service) CompleteTask(title string) (Task, error) {
	return s.repository.CompleteTask(title)
}

func (s *Service) UncompleteTask(title string) (Task, error) {
	return s.repository.UncompleteTask(title)
}

func (s *Service) DeleteTask(title string) error {
	return s.repository.DeleteTask(title)
}