## Hello World Program: ✅

A simple program that prints "Hello, World!" to the console.
Key Concepts: Basic syntax, fmt package, main function.

## Simple Calculator:✅

Create a command-line calculator that can perform basic arithmetic operations like addition, subtraction, multiplication, and division.
Key Concepts: User input, conditional statements, functions.

## Todo List CLI: ✅

A command-line interface (CLI) application to manage a todo list. Users can add, view, and delete tasks.
Key Concepts: Slices, struct, basic CRUD operations.

```go
// Para que los inputs str x consola puedan usar espacio
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var strInput string
    fmt.Println("Enter a string ")
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        strInput = scanner.Text()
    }

    fmt.Println(strInput)
}
```

## Guess the Number: ✅

A game where the program randomly selects a number between 1 and 100, and the user has to guess it. The program provides hints if the guess is too high or too low.
Key Concepts: Random number generation, loops, user input.

```go
//To generate random numbers
func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
```

## HTTP Server: ✅

Create a basic HTTP server that responds with "Hello, World!" to any GET request.
Key Concepts: net/http package, handling HTTP requests and responses.

## URL Shortener:

Develop a simple URL shortener service that takes a long URL and returns a shortened version.
Key Concepts: Maps, HTTP server, handling JSON.

## Weather CLI:

A command-line application that fetches the current weather information for a given city using a weather API.
Key Concepts: Making HTTP requests, working with APIs, JSON parsing.

## Simple Chat Application:

A basic chat application where multiple clients can connect to a server and send messages to each other.
Key Concepts: Concurrency, Goroutines, net package, TCP/UDP protocols.

## Markdown to HTML Converter:

A tool that reads a Markdown file and converts it to HTML.
Key Concepts: File handling, string manipulation, and text parsing.

## Contact Book CLI:

A command-line application to manage contacts with features to add, view, update, and delete contacts.
Key Concepts: CRUD operations, struct, file storage. SQLITE & Some CLI Framework

## Image Resizer:

A program to resize images to specified dimensions.
Key Concepts: Image processing, working with external libraries (e.g., image package).

## JSON to CSV Converter:

Convert JSON data into CSV format and vice versa.
Key Concepts: File handling, JSON and CSV parsing.

## CSV File Processor

Create a program that reads a CSV file, processes the data, and outputs the results. Use loops to read each line of the CSV and perform operations such as calculating averages or filtering data.

## Currency Converter:

A tool to convert amounts between different currencies using real-time exchange rates from an API.
Key Concepts: HTTP requests, working with APIs, JSON parsing.

## Basic Web Scraper:

A program to scrape data from a website and display or save it.
Key Concepts: HTTP requests, HTML parsing.

## Email Sender:

An application to send emails using an SMTP server.
Key Concepts: Working with external packages, email protocols.

## Simple Chatbot:

A basic chatbot that can respond to user input with predefined responses.
Key Concepts: String matching, basic AI concepts.

## Password Generator:

A tool to generate random, strong passwords based on user-defined criteria.
Key Concepts: Random number generation, string manipulation.

## Unit Converter:

A command-line tool to convert between different units (e.g., length, weight, temperature).
Key Concepts: User input, basic arithmetic operations.

# Resources

[General Resources About Go](https://yourbasic.org/golang/)

[Go cheatsheet for using printf](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#default)

[Go debugging with VSCode](https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code)

[Go debugging with VSCode 2](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)
