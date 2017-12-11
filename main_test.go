package main

import "testing"

func TestAddTodo(t *testing.T) {
	todos := []todo{}
	addFlag := "some todo"
	val := add(todos, addFlag)
	if len(val) != 1 {
		t.Errorf("could not add a todo")
	}
}

func TestCompleteTodoWithCompleteFlag(t *testing.T) {
	todos := []todo{todo{Description: "a", Completed: false}}
	completeFlag := 0
	val := complete(todos, completeFlag)
	if val[0].Completed != true {
		t.Error("expected completed to be true, got false")
	}
}

func TestUnCompleteTodoWithCompleteFlag(t *testing.T) {
	todos := []todo{todo{Description: "a", Completed: true}}
	completeFlag := 0
	val := complete(todos, completeFlag)
	if val[0].Completed != false {
		t.Error("expected completed to be false, got true")
	}
}

func TestNoCompleteTodoWithNegativeOneCompleteFlag(t *testing.T) {
	todos := []todo{todo{Description: "a", Completed: false}}
	completeFlag := -1
	val := complete(todos, completeFlag)
	if val[0].Completed != false {
		t.Error("expected completed to be false, got true")
	}
}

func TestRemoveTodoWithRemoveFlag(t *testing.T) {
	todos := defaultTodos()
	removeFlag := 1
	val := remove(todos, removeFlag)
	if len(val) != 2 {
		t.Error("did not remove todo")
	}
}

func TestNoRemoveTodoWithNoRemoveFlag(t *testing.T) {
	todos := defaultTodos()
	removeFlag := -1
	val := remove(todos, removeFlag)
	if len(val) != 3 {
		t.Error("unexpectedly removed todo")
	}
}

var listWithCompletion = []struct {
	List    []todo
	Show    bool
	Only    bool
	Length  int
	Message string
}{
	{todosWithCompletion(), false, false, 3, "showed completed when should not"},
	{todosWithCompletion(), true, false, 4, "wrong number of todos with show completed"},
	{todosWithCompletion(), false, true, 2, "wrong number of todos with only completed"},
}

func TestListWithCompletion(t *testing.T) {
	for _, lwc := range listWithCompletion {

		val := list(lwc.List, lwc.Show, lwc.Only)
		// list adds a header row
		if len(val) != lwc.Length {
			t.Error(lwc.Message)
		}
	}
}

func todosWithCompletion() []todo {
	return []todo{
		todo{Description: "a", Completed: false},
		todo{Description: "b", Completed: false},
		todo{Description: "c", Completed: true},
	}
}

func defaultTodos() []todo {
	return []todo{
		todo{Description: "a"},
		todo{Description: "b"},
		todo{Description: "c"},
	}
}
