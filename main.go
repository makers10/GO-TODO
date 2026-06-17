package main
import ( "fmt"
"encoding/json"
"os"
"bufio"
"strings"

)
var tasks []Task

type Task struct {
	ID int
	Title string
	Completed bool

}
func getNextID() int {
    maxID := 0

    for _, task := range tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }

    return maxID + 1
}

func main(){
	loadTasks()
	scanner := bufio.NewScanner(os.Stdin)

 for {
	fmt.Println("\nTodo App")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Tasks")
	fmt.Println("4. Delete task")
	fmt.Println("5. Exit")

    fmt.Print("choose: ")
    scanner.Scan()
	var choice int
	fmt.Sscanf(scanner.Text(), "%d", &choice)

	switch choice {
	case 1:
	    fmt.Print("Task title: ")
		scanner.Scan()
		title := strings.TrimSpace(scanner.Text())

		addTask(title)
		saveTasks()

	case 2:
		listTasks()

	case 3:
		var id int
		fmt.Print("Task ID: ")
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &id)

		completeTask(id)
		saveTasks()

	case 4:
		var id int
		fmt.Print("Task ID: ")
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &id)

		deleteTask(id)
		saveTasks()

	case 5:
		fmt.Println("See you later")
		return
	}
  }
}

// addTask("aaj Paneer bnega")
	// addTask("Socha to yhi hai")
	// fmt.Println("Todo App")

// var tasks []Task
func addTask(title string){
	task := Task{
		// ID: len(tasks) + 1,
		ID:getNextID(),
		Title: title,
		Completed: false,
	}
	tasks = append(tasks, task)
}
func listTasks(){
	for _, task := range tasks {
		status := ""
		if task.Completed {
			status = "-/"
		}
		fmt.Printf("[%s] %d - %s\n",
	status,
    task.ID,
    task.Title)
	}
}
func saveTasks(){
	data, _ := json.Marshal(tasks)
	os.WriteFile("tasks.json", data, 0644)
}
func completeTask(id int) {
	for i:= range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return
		}
	}
}
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		return
		}
	}
}
func loadTasks(){
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &tasks)
  }

