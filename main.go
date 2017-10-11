package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/ryanuber/columnize"
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
var showCompletedFlag bool
var configFormat *columnize.Config

func main() {
	configureTable()
	parseFlags()
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("could not get current user: %v", err))
	}
	app := todoApp{DataFilePath: fmt.Sprintf("%s/.todo", usr.HomeDir)}

	todos := read(app)
	todos = add(todos)
	todos = complete(todos)
	todos = remove(todos)
	err = write(todos, app)
	if err != nil {
		fmt.Printf("error writing file: %v", err)
		os.Exit(1)
	}
	list(todos)

}

func configureTable() {
	configFormat = columnize.DefaultConfig()
	configFormat.Glue = " | "
}
func parseFlags() {
	flag.StringVar(&addFlag, "a", "", "Add a todo using -a")
	flag.IntVar(&completeFlag, "c", -1, "Complete a todo by index `-c 0`.")
	flag.IntVar(&removeFlag, "r", -1, "Remove a todo by index `-r 0`.")
	flag.BoolVar(&showCompletedFlag, "show-completed", false, "Show Completed Todos `--show-completed`")
	flag.Parse()
}

func add(todos []todo) []todo {
	if len(addFlag) > 0 {
		todos = append(todos, todo{Description: addFlag, Completed: false, CreatedAt: time.Now()})
	}
	return todos
}

func complete(todos []todo) []todo {
	if completeFlag > -1 {
		for i := range todos {
			if i == completeFlag {
				todos[i].Completed = !todos[i].Completed
				todos[i].CompletedAt = time.Now()
			}
		}
	}
	return todos
}

func remove(todos []todo) []todo {
	if removeFlag > -1 {
		todos = append(todos[:removeFlag], todos[removeFlag+1:]...)
	}
	return todos
}

func list(todos []todo) {
	var output = make([]string, 0)
	output = append(output, "Num | Description | Completed | Created At")
	for i, x := range todos {
		if showCompletedFlag {
			output = append(output, printTodo(i, x))
		} else {
			if !x.Completed {
				output = append(output, printTodo(i, x))
			}
		}
	}
	result := columnize.Format(output, configFormat)
	fmt.Println(result)
}

func printTodo(i int, t todo) string {
	return fmt.Sprintf("%d | %s | %v | %v", i, t.Description, t.Completed, t.CreatedAt)
}

func write(todos []todo, app todoApp) error {
	f, err := os.OpenFile(app.DataFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
		os.Exit(1)
	}
	f.Truncate(0)
	f.Seek(0, 0)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, t := range todos {

		b, err := json.Marshal(&t)
		if err != nil {
			fmt.Printf("There was an error marshaling the todo: %v\n", err)
			os.Exit(1)
		}
		_, err = w.WriteString(fmt.Sprintf("%v\n", string(b)))
		w.Flush()
	}
	return err
}

func read(app todoApp) []todo {
	todos := make([]todo, 0)
	f, err := os.Open(app.DataFilePath)
	defer f.Close()
	if err == nil {
		r := bufio.NewScanner(f)
		for r.Scan() {
			var t todo
			json.Unmarshal(r.Bytes(), &t)
			todos = append(todos, t)
		}
	}
	return todos
}
