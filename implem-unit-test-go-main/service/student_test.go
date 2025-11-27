package service

import (
	"session-9/model"
	"session-9/repository"
	"testing"
)

func newTestService(student []model.Student) (*StudentService, *repository.MockStudentRepository) {
	repo := &repository.MockStudentRepository{Students: student}
	service := NewStudentService(repo)
	return service, repo
}

func TestStudent_Create(t *testing.T) {
	service, repo := newTestService([]model.Student{})

	created, err := service.Create(model.Student{
		Name: "Rudi",
		Age:  20,
	})

	if err != nil {
		t.Fatalf("Created returned error: %v", err)
	}

	if created.ID != 1 {
		t.Error("expected ID 1, got %d", created.ID)
	}

	if len(repo.Students) != 1 {
		t.Fatalf("expected repo to have 1 student, got %d", len(repo.Students))
	}
}
