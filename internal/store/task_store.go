package store

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/sakib-maho/golang-beego-restapi-swagger-v1/internal/model"
)

var ErrNotFound = errors.New("task not found")

type TaskStore struct {
	mu    sync.RWMutex
	tasks map[string]model.Task
	seq   int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: map[string]model.Task{},
		seq:   0,
	}
}

func (s *TaskStore) List() []model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]model.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		items = append(items, t)
	}
	return items
}

func (s *TaskStore) Create(req model.CreateTaskRequest) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seq++
	now := time.Now().UTC()
	status := req.Status
	if status == "" {
		status = "todo"
	}

	task := model.Task{
		ID:          fmt.Sprintf("task-%d", s.seq),
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	s.tasks[task.ID] = task
	return task
}

func (s *TaskStore) Get(id string) (model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[id]
	if !ok {
		return model.Task{}, ErrNotFound
	}
	return task, nil
}

func (s *TaskStore) Update(id string, req model.UpdateTaskRequest) (model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return model.Task{}, ErrNotFound
	}
	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	task.UpdatedAt = time.Now().UTC()
	s.tasks[id] = task
	return task, nil
}

func (s *TaskStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(s.tasks, id)
	return nil
}
