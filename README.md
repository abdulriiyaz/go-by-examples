# Go Learning Journey

Welcome to my Go learning journey! In this repository, I document my progress and showcase examples as I learn and explore the Go programming language. The main focus of this repository is to demonstrate different concepts and features of Go through practical examples.

## Table of Contents

-   [Getting Started](#getting-started)
    -   [Go Setup and Installation](#go-setup-and-installation)
    -   [Running the Examples](#running-the-examples)
    -   [Making Requests](#making-requests)
-   [Examples](#examples)
-   [Contributing](#contributing)
-   [License](#license)

## Getting Started

### Go Setup and Installation

To get started with Go, you need to set up the Go development environment on your machine. Follow the instructions below to install Go:

1. Visit the official Go website: [https://golang.org](https://golang.org)
2. Download the Go installer appropriate for your operating system.
3. Run the installer and follow the on-screen instructions to complete the installation process.

Once Go is installed, you can verify the installation by opening a terminal or command prompt and executing the following command:

If the installation was successful, it will display the installed Go version.

### Running the Examples

To run the Go examples in this repository, follow these steps:

1. Clone this repository using the following command:

2. Navigate to the specific example directory that you want to run. For example, to run the TODO CRUD API example, go to the `todo-crud-api` directory:

3. Start the Go program by running the following command:

This will start the program, and you will see output indicating that the server is running.

### Making Requests

You can use a service like Postman to make HTTP requests to the API endpoints provided by the examples. Follow the steps below to make requests using Postman:

1. Install Postman: Visit the Postman website ([https://www.postman.com](https://www.postman.com)) and download the appropriate version of Postman for your operating system. Install Postman following the provided instructions.

2. Open Postman and create a new request.

3. Set the request URL to the appropriate endpoint. For example, for the TODO CRUD API, you can use `http://localhost:8080/todos` to retrieve all TODO items.

4. Choose the appropriate HTTP method (GET, POST, PUT, DELETE) for the desired operation.

5. Optionally, set request headers and add request body if required.

6. Click the "Send" button to make the request and see the response.

You can repeat these steps for different API endpoints and operations as demonstrated in each example.

## Examples

### Example 1: TODO CRUD API

The `todo-crud-api` directory contains an example of a simple TODO CRUD API implemented in Go. This example showcases how to build a basic API using the Gorilla Mux package. It demonstrates how to perform CRUD operations on TODO items using arrays as the data storage mechanism.

To run the example, navigate to the `todo-crud-api` directory and execute the following command:

This will start the API server, and you can access the API endpoints using an HTTP client such as Postman.

### Example 2: ...

## Contributing

Contributions are welcome! If you have any improvements, bug fixes, or additional examples to share,
