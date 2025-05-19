//
// @package Showcase-OCI-Registries
//
// @file Todo list repository
// @copyright 2023-present Christoph Kappel <christoph@unexist.dev>
// @version $Id$
//
// This program can be distributed under the terms of the Apache License v2.0.
// See the file LICENSE for details.
//

package infrastructure

import (
	"errors"

	"braces.dev/errtrace"

	"github.com/unexist/showcase-oras/domain"
)

type TodoListRepository struct {
	todos []domain.Todo
}

func NewTodoListRepository() *TodoListRepository {
	return &TodoListRepository{
		todos: make([]domain.Todo, 0),
	}
}

func (repository *TodoListRepository) Open(_ string) error {
	return nil
}

func (repository *TodoListRepository) GetTodos() ([]domain.Todo, error) {
	return repository.todos, nil
}

func (repository *TodoListRepository) CreateTodo(todo *domain.Todo) error {
	newTodo := domain.Todo{
		ID:          len(repository.todos) + 1,
		Title:       todo.Title,
		Description: todo.Description,
	}

	todo.ID = newTodo.ID

	repository.todos = append(repository.todos, newTodo)

	return nil
}

func (repository *TodoListRepository) GetTodo(todoId int) (*domain.Todo, error) {
	for i := 0; i < len(repository.todos); i++ {
		if repository.todos[i].ID == todoId {
			return &repository.todos[i], nil
		}
	}

	return nil, errtrace.Wrap(errors.New("Not found"))
}

func (repository *TodoListRepository) UpdateTodo(todo *domain.Todo) error {
	for i := 0; i < len(repository.todos); i++ {
		if repository.todos[i].ID == todo.ID {
			repository.todos[i].Title = todo.Title
			repository.todos[i].Description = todo.Description

			return nil
		}
	}

	return errtrace.Wrap(errors.New("Not found"))
}

func (repository *TodoListRepository) DeleteTodo(todoId int) error {
	for i := 0; i < len(repository.todos); i++ {
		if repository.todos[i].ID == todoId {
			repository.todos = removeIndex(repository.todos, i)

			return nil
		}
	}

	return errtrace.Wrap(errors.New("Not found"))
}

func (repository *TodoListRepository) Close() error {
	return nil
}

func (repository *TodoListRepository) Clear() error {
	repository.todos = make([]domain.Todo, 0)

	return nil
}

func removeIndex(s []domain.Todo, idx int) []domain.Todo {
	return append(s[:idx], s[idx+1:]...)
}
