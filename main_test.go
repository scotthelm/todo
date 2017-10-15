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
	todos := []todo{
		todo{Description: "a"},
		todo{Description: "b"},
		todo{Description: "c"},
	}
	removeFlag := 1
	val := remove(todos, removeFlag)
	if len(val) != 2 {
		t.Error("did not remove todo")
	}
}

func TestNoRemoveTodoWithNoRemoveFlag(t *testing.T) {
	todos := []todo{
		todo{Description: "a"},
		todo{Description: "b"},
		todo{Description: "c"},
	}
	removeFlag := -1
	val := remove(todos, removeFlag)
	if len(val) != 3 {
		t.Error("unexpectedly removed todo")
	}
}

func TestListWithNoCompleted(t *testing.T) {
	todos := []todo{
		todo{Description: "a", Completed: false},
		todo{Description: "b", Completed: false},
		todo{Description: "c", Completed: true},
	}
	showCompletedFlag := false
	onlyCompletedFlag := false
	val := list(todos, showCompletedFlag, onlyCompletedFlag)
	// list adds a header row
	if len(val) > 3 {
		t.Error("showed completed when should not")
	}
}

func TestListWithShowCompleted(t *testing.T) {
	todos := []todo{
		todo{Description: "a", Completed: false},
		todo{Description: "b", Completed: false},
		todo{Description: "c", Completed: true},
	}
	showCompletedFlag := true
	onlyCompletedFlag := false
	val := list(todos, showCompletedFlag, onlyCompletedFlag)
	// list adds a header row
	if len(val) != 4 {
		t.Error("wrong number of todos with show completed")
	}
}

func TestListWithOnlyCompleted(t *testing.T) {
	todos := []todo{
		todo{Description: "a", Completed: false},
		todo{Description: "b", Completed: false},
		todo{Description: "c", Completed: true},
	}
	showCompletedFlag := false
	onlyCompletedFlag := true
	val := list(todos, showCompletedFlag, onlyCompletedFlag)
	// list adds a header row
	if len(val) != 2 {
		t.Error("wrong number of todos with only completed")
	}
}
