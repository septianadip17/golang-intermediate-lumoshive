package repository

import "session-9/model"

type MockStudentRepository struct {
	Students []model.Student
	ErrGet   error
	ErrSave  error
}

func (mockStudentRepository *MockStudentRepository) GetAll() ([]model.Student, error) {
	return mockStudentRepository.Students, mockStudentRepository.ErrGet
}

func (mockStudentRepository *MockStudentRepository) SaveAll(students []model.Student) error {
	mockStudentRepository.Students = students
	return mockStudentRepository.ErrSave
}
