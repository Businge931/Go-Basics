package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// type str string

type saver interface {
	Save() error
}

// type displayer interface{
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface{
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("test")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)

}

func add(a, b interface{}) interface{} {
	aInt, aIsInt :=a.(int)
	bInt, bIsInt :=b.(int)

	if aIsInt && bIsInt{
		return aInt +bInt
	}

	aFloat, aIsFloat :=a.(float64)
	bFloat, bIsFloat :=b.(float64)

	if aIsFloat && bIsFloat{
		return aFloat +bFloat
	}

	aString, aIsString :=a.(string)
	bString, bIsString :=b.(string)

	if aIsString && bIsString{
		return aString +bString
	}

	return nil
}

func printSomething(value interface{}) { // You can also use the "any alias" insted of  interface{}
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Flaot: ", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Flaot: ", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)

	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
