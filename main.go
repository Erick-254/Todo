//defines main package as the entry point
package main
//import packages used in the projects
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/
type Todo struct {
	ID     int
	Task   string
	Done   bool
}
//declare variables 
var todos []Todo
var currentID int
//the main entry point of our programme

func main() {
	reader := bufio.NewReader(os.Stdin)
	//loops through and assigns id
	for {
		fmt.Println("== Todo App ==")
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Mark Todo as Done")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			addTodo(task)
		case "2":
			listTodos()// prints all todo
		case "3":
			fmt.Print("Enter the ID of the todo to mark as done: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			markDone(id)
		case "4":
			fmt.Print("Enter the ID of the todo to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			deleteTodo(id)
		//the end 	
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func addTodo(task string) {
	currentID++
	todo := Todo{
		ID:     currentID,
		Task:   task,
		Done:   false,
	}
	todos = append(todos, todo)
	fmt.Println("Todo added successfully!")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	fmt.Println("== Todos ==")
	for _, todo := range todos {
		status := "Not Done"
		if todo.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s - %s\n", todo.ID, todo.Task, status)
	}
}
//hakuna haja
func markDone(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			fmt.Println("Todo marked as done.")
			return
		}
	}
	fmt.Println("Todo not found.")
}

func deleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted.")
			return
		}
	}
	fmt.Println("Todo not found.")
}
