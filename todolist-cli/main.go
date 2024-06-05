package main

import (
	"fmt"
)

// ## Todo List CLI:

// A command-line interface (CLI) application to manage a todo list. Users can add, view, and delete tasks.
// Key Concepts: Slices, struct, basic CRUD operations.

/*
Inicializa el programa:
1. crea la variable que contiene Todos
2. Opciones posibles:
	1. agregar todo
		a. poner una breve descripcion
	2. ver los todos pendientes
		a. iterar los todos q sean pendientes
	3. ver los todos completados
	 a. iterar los todos q esten completados
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

var todos = make([]Todo, 0)

func addNewTodo(idx int) {
	var newTodoTitle string
	fmt.Println("Ingrese una Todo: ")
	fmt.Scan(&newTodoTitle)
	newTodo := Todo{
		id:           idx,
		title:        newTodoTitle,
		is_completed: false,
	}
	todos = append(todos, newTodo)
}

func todoOptions(opt string, idx int) {
	switch opt {
	case "1":
		fmt.Println("[1]AGREGAR un TODO")
		addNewTodo(idx)
	case "2":
		fmt.Println("[2]VER TODOs pendientes")
	case "3":
		fmt.Println("[3]VER TODOs completados")
	case "4":
		fmt.Println("[4]ELIMINAR un TODO")
	case "0":
		fmt.Println("[0]Terminar")
	}

}

func initializeMenu() {
	fmt.Println("Bienvenido a TODO CLI")
	fmt.Println("Que desea hacer?")
	fmt.Println("[1] Agregar TODO")
	fmt.Println("[2] Ver TODO pendiente")
	fmt.Println("[3] Ver TODO completedos")
	fmt.Println("[4] Eliminar TODO")
	fmt.Println("[0] Terminar")
}

func main() {
	var opt string
	idx := 1
	for opt != "0" {
		initializeMenu()
		fmt.Scan(&opt)
		fmt.Printf("Ha seleccionado la opcion %s \n", opt)
		todoOptions(opt, idx)
		idx++
		fmt.Println(todos)
	}

}
