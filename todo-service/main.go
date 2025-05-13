//
// @package Showcase-Oras
//
// @file Todo main
// @copyright 2023-present Christoph Kappel <christoph@unexist.dev>
// @version $Id$
//
// This program can be distributed under the terms of the Apache License v2.0.
// See the file LICENSE for details.
//

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unexist/showcase-oras/adapter"
	"github.com/unexist/showcase-oras/domain"
	"github.com/unexist/showcase-oras/infrastructure"

	"log"
)

func main() {
	/* Create business stuff */
	var todoRepository *infrastructure.TodoListRepository

	todoRepository = infrastructure.NewTodoListRepository()

	defer todoRepository.Close()

	todoService := domain.NewTodoService(todoRepository)
	todoResource := adapter.NewTodoResource(todoService)

	/* Finally start Gin */
	engine := gin.Default()

	todoResource.RegisterRoutes(engine)

	log.Fatal(http.ListenAndServe(":8080", engine))
}
