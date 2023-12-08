// Code generated by MockGen. DO NOT EDIT.
// Source: repo_crew.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockICrewRepo is a mock of ICrewRepo interface.
type MockICrewRepo struct {
	ctrl     *gomock.Controller
	recorder *MockICrewRepoMockRecorder
}

// MockICrewRepoMockRecorder is the mock recorder for MockICrewRepo.
type MockICrewRepoMockRecorder struct {
	mock *MockICrewRepo
}

// NewMockICrewRepo creates a new mock instance.
func NewMockICrewRepo(ctrl *gomock.Controller) *MockICrewRepo {
	mock := &MockICrewRepo{ctrl: ctrl}
	mock.recorder = &MockICrewRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICrewRepo) EXPECT() *MockICrewRepoMockRecorder {
	return m.recorder
}

// FindActor mocks base method.
func (m *MockICrewRepo) FindActor(name, birthDate string, films, career []string, country string) ([]models.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActor", name, birthDate, films, career, country)
	ret0, _ := ret[0].([]models.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActor indicates an expected call of FindActor.
func (mr *MockICrewRepoMockRecorder) FindActor(name, birthDate, films, career, country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActor", reflect.TypeOf((*MockICrewRepo)(nil).FindActor), name, birthDate, films, career, country)
}

// GetActor mocks base method.
func (m *MockICrewRepo) GetActor(actorId uint64) (*models.CrewItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", actorId)
	ret0, _ := ret[0].(*models.CrewItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockICrewRepoMockRecorder) GetActor(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockICrewRepo)(nil).GetActor), actorId)
}

// GetFilmCharacters mocks base method.
func (m *MockICrewRepo) GetFilmCharacters(filmId uint64) ([]models.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmCharacters", filmId)
	ret0, _ := ret[0].([]models.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmCharacters indicates an expected call of GetFilmCharacters.
func (mr *MockICrewRepoMockRecorder) GetFilmCharacters(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmCharacters", reflect.TypeOf((*MockICrewRepo)(nil).GetFilmCharacters), filmId)
}

// GetFilmDirectors mocks base method.
func (m *MockICrewRepo) GetFilmDirectors(filmId uint64) ([]models.CrewItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmDirectors", filmId)
	ret0, _ := ret[0].([]models.CrewItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmDirectors indicates an expected call of GetFilmDirectors.
func (mr *MockICrewRepoMockRecorder) GetFilmDirectors(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmDirectors", reflect.TypeOf((*MockICrewRepo)(nil).GetFilmDirectors), filmId)
}

// GetFilmScenarists mocks base method.
func (m *MockICrewRepo) GetFilmScenarists(filmId uint64) ([]models.CrewItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmScenarists", filmId)
	ret0, _ := ret[0].([]models.CrewItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmScenarists indicates an expected call of GetFilmScenarists.
func (mr *MockICrewRepoMockRecorder) GetFilmScenarists(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmScenarists", reflect.TypeOf((*MockICrewRepo)(nil).GetFilmScenarists), filmId)
}