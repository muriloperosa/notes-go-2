package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/muriloperosa/notes-go-2/note"
	"github.com/muriloperosa/notes-go-2/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	add(1, 3)

	printSomething(":)")
	printSomething(10)
	printSomething(1.9)
	printSomething(true)

	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
	outputData(todo)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value interface{}) {

	intVal, isInt := value.(int)

	if isInt {
		fmt.Println("Value is Int:", intVal)
		return
	}

	stringVal, isString := value.(string)

	if isString {
		fmt.Println("Value is String:", stringVal)
		return
	}

	fmt.Println("Value is not Int or String:", value)

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float64:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// 	// default:
	// 	// 	fmt.Println("Default", value)
	// }

	fmt.Println(value)
}

// func printSomething(value any) {
// 	fmt.Println(value)
// }
