# Command-Line Applications with Cobra and Viper

Welcome to the repository containing a collection of command-line applications developed using Go, Cobra, and Viper. This project explores various aspects of command-line tool development, covering a wide range of functionalities and advanced features.

## Table of Contents

- [Command-Line Applications with Cobra and Viper](#command-line-applications-with-cobra-and-viper)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Projects](#projects)
  - [Getting Started](#getting-started)
  - [Technologies Used](#technologies-used)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

This repository showcases a series of command-line applications built with Go programming language, emphasizing the use of the Cobra CLI framework for command-line interaction and Viper for configuration management. Each project in this collection addresses different aspects of command-line tool development, providing practical examples and insights into various Go programming concepts.

## Projects

1. [Word Counter](./wordcounter) 
   - Basic implementation of a word counter.
   - Addition of features, exploration of testing, and building for different platforms.

2. [To-Do List Manager](./todoServer) 
   - Command-line tool for managing lists of to-do items.
   - Input from STDIN, parsing command-line parameters, and defining flags.
   - Utilizing environment variables for increased flexibility.

3. [Markdown File Previewer](./markdown_previewer) 
   - Tool to preview Markdown files using a web browser.
   - Handling file paths, using temporary files, and applying file templates.
   - Implementation of Go interfaces for code flexibility.

4. [File System Navigator](./file_system_crawler) 
   - CLI application to find, delete, and back up files.
   - Common file system operations, logging, and table-driven testing.

5. [CSV Data Processor](./colStats) 
   - Command-line tool processing data from CSV files.
   - Benchmarking, profiling, and tracing for performance analysis.
   - Concurrent processing with goroutines and channels.

6. [Network Tool with Cobra](./pScan) 
   - Network tool executing a TCP port scan on remote machines.
   - Utilizing the Cobra CLI framework for flexible subcommands.

8. [REST API Client](./todoServer) 
   - Enhancing a to-do application with a REST API.
   - Developing a command-line client for API interaction and testing.


## Getting Started

To get started with any of the projects, follow the instructions provided in each project's directory. Ensure you have Go (version 1.x) installed on your machine.

1. Clone the repository:

    ```bash
    git clone https://github.com/Micah-Shallom/golang_cli_applications.git
    ```

2. Navigate to the project directory:

    ```bash
    cd golang_cli_applications/project_directory
    ```

3. Follow the specific project's README for detailed instructions on running and testing the application.

## Technologies Used

- Go (Official Go Website)
- Cobra CLI Framework
- Viper Configuration Management

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit pull requests. Contributions, suggestions, and improvements are welcome!

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code for your own purposes.