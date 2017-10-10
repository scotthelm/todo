package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"time"
)

type todo struct {
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type todoApp struct {
	DataFilePath string
}

var addFlag string
var completeFlag int
var removeFlag int

func (t *todoApp) add(td todo) error {
	b, err := json.Marshal(&td)
	if err != nil {
		fmt.Printf("There was an error marshaling the todo: %v\n", err)
	}
	f, err := os.OpenFile(t.DataFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(fmt.Sprintf("%v\n", string(b)))
	w.Flush()

	return err
}
func main() {
	parseFlags()
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("could not get current user: %v", err))
	}
	app := todoApp{DataFilePath: fmt.Sprintf("%s/.todo", usr.HomeDir)}

	add(app)
	todos := read(app)
	todos = complete(todos)
	todos = remove(todos)
	write(todos, app)
	list(todos)

}

func parseFlags() {
	flag.StringVar(&addFlag, "a", "", "Add a todo using -a")
	flag.IntVar(&completeFlag, "c", -1, "Complete a todo by index `-c 0`.")
	flag.IntVar(&removeFlag, "r", -1, "Remove a todo by index `-r 0`.")
	flag.Parse()
}

func add(app todoApp) {
	if len(addFlag) > 0 {
		app.add(todo{Description: addFlag, Completed: false})
	}
}

func complete(todos []todo) []todo {
	if completeFlag > -1 {

	}
	return todos
}

func remove(todos []todo) []todo {
	if removeFlag > -1 {

	}
	return todos
}

func list(todos []todo) {
	for i, x := range todos {
		printTodo(i, x)
	}
}

func write(todos []todo, app todoApp) {
	os.Remove(app.DataFilePath)
	for _, t := range todos {
		app.add(t)
	}
}

func printTodo(i int, t todo) {
	fmt.Printf("%d\t%s : %v\n", i, t.Description, t.Completed)
}

func read(app todoApp) []todo {
	todos := make([]todo, 0)
	f, err := os.Open(app.DataFilePath)
	defer f.Close()
	if err != nil {
		panic(fmt.Sprintf("Unable to open todo file. Try adding a todo? :%v", err))
	}
	r := bufio.NewScanner(f)
	for r.Scan() {
		var t todo
		json.Unmarshal(r.Bytes(), &t)
		todos = append(todos, t)
	}

	return todos
}
