package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Println("Enter command (add, remove, complete, list, or exit)")

		userInput := scanner.Scan()
		if !userInput {
			break
		}
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Println("Enter task to add:")
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")
		case "remove":
			fmt.Println("Enter task number to remove:")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
			}
			taskNumber, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if taskNumber < 1 || taskNumber > len(tasks) {
				fmt.Println("Task number out of range.")
				continue
			}
			tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
			fmt.Println("Task removed successfully!")
		case "complete":
			fmt.Println("Enter task number to mark complete")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				continue
			}
			taskNumber, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if taskNumber < 1 || taskNumber > len(tasks) {
				fmt.Println("Task number out of range.")
				continue
			}
			fmt.Printf("Task '%s' marked complete.\n", tasks[taskNumber-1])
			tasks[taskNumber-1] = "**COMPLETED** - " + tasks[taskNumber-1]
		case "list":
			if len(tasks) == 0 {
				fmt.Println("You don't have any tasks yet.")
			} else {
				fmt.Println("Your tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
