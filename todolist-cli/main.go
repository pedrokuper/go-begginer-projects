package main

import (
	"bufio"
	"fmt"
	"os"
)

// ## Todo List CLI:

// A command-line interface (CLI) application to manage a todo list. Users can add, view, and delete tasks.
// Key Concepts: Slices, struct, basic CRUD operations.

/*
Inicializa el programa:

2. Opciones posibles:

	4. borrar un todo
		a. recibir un numerico
		b. buscarlo
		c. eliminarlo
*/

type Todo struct {
	id           int
	title        string
	is_completed bool
}

var todos = make([]Todo, 0) //Slice de todos

func main() {
	var opt string
	idx := 1
	for opt != "0" {
		initializeMenu()
		fmt.Scan(&opt)
		fmt.Printf("Ha seleccionado la opcion [%s] \n", opt)
		todoOptions(opt, &idx)
	}
}

func addNewTodo(idx *int) {
	var newTodoTitle string
	fmt.Println("Ingrese una Todo: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newTodoTitle = scanner.Text()
	}
	newTodo := Todo{
		id:           *idx,
		title:        newTodoTitle,
		is_completed: false,
	}
	todos = append(todos, newTodo)
	*idx++
}

func showTodos(isCompleted bool, todoList []Todo) {
	//Si no hay todos mostrar "No hay todos pendientes | completados"
	todoType := "pendientes"
	if isCompleted {
		todoType = "completados"
	}
	fmt.Printf("Listando todos %s: \n", todoType)
	fmt.Println("*****************")
	for i := range todoList {
		if todoList[i].is_completed == isCompleted {
			fmt.Printf("ID: %d \nTitulo: %s \nCompletado: %t\n", todos[i].id, todos[i].title, todos[i].is_completed)
			fmt.Println("*****************")

		}
	}
}

func searchTodo(idx int, todoList []Todo) (*Todo, bool) {
	for i := range todoList {
		if todoList[i].id == idx {
			return &todoList[i], true
		}
	}
	return &Todo{}, false
}

func completeTodo() {
	var idxToFind int
	showTodos(false, todos)
	fmt.Println("Ingrese el todo que desea completar")
	fmt.Scan(&idxToFind)
	todo, found := searchTodo(idxToFind, todos)
	if found {
		todo.is_completed = true
	}
}

func todoOptions(opt string, idx *int) {
	switch opt {
	case "1":
		fmt.Println("Agregue la descripcion del TODO")
		addNewTodo(idx)
	case "2":
		fmt.Println("[2]VER TODOs pendientes")
		showTodos(false, todos)
	case "3":
		fmt.Println("[3]VER TODOs completados")
		showTodos(true, todos)
	case "4":
		fmt.Println("[4]COMPLETAR un TODO")
		completeTodo()
	case "5":

	case "0":
		fmt.Println("[0]Terminar")
	}

}

func initializeMenu() {
	fmt.Println("")
	fmt.Println("Bienvenido a TODO CLI")
	fmt.Println("Que desea hacer?")
	fmt.Println("1. Agregar TODO")
	fmt.Println("2. Ver TODO pendiente")
	fmt.Println("3. Ver TODO completedos")
	fmt.Println("4. Completar TODO")
	fmt.Println("0. Terminar")
	fmt.Println("")
}
