package main

import "testing"

func TestAddTodo(t *testing.T) {
	app := todoApp{"~/.todo"}
	todo := todo{Description: "it adds a todo", Completed: false}
	err := app.add(todo)
	if err != nil {
		t.Errorf("could not add a todo: %v", err)
	}
}
