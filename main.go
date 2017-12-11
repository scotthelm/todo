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

var configFormat *columnize.Config

func main() {
	configureTable()

	addFlag := flag.String("a", "", "Add a todo")
	completeFlag := flag.Int("c", -1, "Complete a todo by Num")
	removeFlag := flag.Int("r", -1, "Remove a todo by Num")
	showCompletedFlag := flag.Bool("show-completed", false, "Show Completed Todos")
	completedOnlyFlag := flag.Bool("only-completed", false, "Show Only Completed Todos")
	backupFlag := flag.String("b", "", "Backup Todo File To Path")
	flag.Parse()

	app := newAppForUser(".todo")

	todos := read(app)
	todos = add(todos, *addFlag)
	todos = complete(todos, *completeFlag)
	todos = remove(todos, *removeFlag)
	err := write(todos, app.DataFilePath)
	err = backup(todos, *backupFlag)
	if err != nil {
		fmt.Printf("error writing file: %v", err)
		os.Exit(1)
	}
	output := list(todos, *showCompletedFlag, *completedOnlyFlag)
	result := columnize.Format(output, configFormat)
	fmt.Println(result)
}

func configureTable() {
	configFormat = columnize.DefaultConfig()
	configFormat.Glue = " | "
}

func newAppForUser(dataFileName string) todoApp {
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("could not get current user: %v", err))
	}
	return todoApp{DataFilePath: fmt.Sprintf("%s/%s", usr.HomeDir, dataFileName)}
}

func add(todos []todo, addFlag string) []todo {
	if len(addFlag) > 0 {
		todos = append(todos, todo{Description: addFlag, Completed: false, CreatedAt: time.Now()})
	}
	return todos
}

func complete(todos []todo, completeFlag int) []todo {
	if completeFlag < 0 || completeFlag > len(todos)-1 {
		return todos
	}
	for i := range todos {
		if i == completeFlag {
			todos[i].Completed = !todos[i].Completed
			todos[i].CompletedAt = time.Now()
		}
	}
	return todos
}

func remove(todos []todo, removeFlag int) []todo {
	if removeFlag > -1 {
		todos = append(todos[:removeFlag], todos[removeFlag+1:]...)
	}
	return todos
}

func list(todos []todo, showCompletedFlag, completedOnlyFlag bool) []string {
	var output = make([]string, 0)
	output = append(output, "Num | Description | Completed | Created At | Completion Time")
	for i, x := range todos {
		if showCompletedFlag {
			output = append(output, printTodo(i, x))
		} else if completedOnlyFlag {
			if x.Completed {
				output = append(output, printTodo(i, x))
			}
		} else {
			if !x.Completed {
				output = append(output, printTodo(i, x))
			}
		}
	}
	return output
}

func printTodo(i int, t todo) string {
	var elapsedTime time.Duration
	if t.CompletedAt.IsZero() {
		since := time.Since(t.CreatedAt)
		elapsedTime = since - (since % time.Minute)
	} else {
		since := t.CompletedAt.Sub(t.CreatedAt)
		elapsedTime = since - (since % time.Minute)
	}
	return fmt.Sprintf(
		"%d | %s | %v | %v | %v",
		i,
		t.Description,
		t.Completed,
		t.CreatedAt.Format("2006-01-02@03:04:05"),
		elapsedTime,
	)
}

func write(todos []todo, path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		reportErrorAndExit("There was an error opening the file: %v\n", err)
	}
	f.Truncate(0)
	f.Seek(0, 0)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, t := range todos {

		b, err := json.Marshal(&t)
		if err != nil {
			reportErrorAndExit("There was an error marshaling the todo: %v\n", err)
		}
		_, err = w.WriteString(fmt.Sprintf("%v\n", string(b)))
		w.Flush()
	}
	return err
}

func read(app todoApp) []todo {
	todos := make([]todo, 0)
	f, err := os.OpenFile(app.DataFilePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		reportErrorAndExit("unable to read todo file: %v", err)
	}
	defer f.Close()
	r := bufio.NewScanner(f)
	for r.Scan() {
		var t todo
		json.Unmarshal(r.Bytes(), &t)
		todos = append(todos, t)
	}
	return todos
}

func backup(todos []todo, path string) error {
	if path != "" && path != "./todo" {
		return write(todos, path)
	}
	return nil
}

func reportErrorAndExit(message string, err error) {
	os.Stderr.WriteString(fmt.Sprintf(message, err))
	os.Exit(1)
}
