# TODO CLI Tool

This is a simple command-line tool developed by Micah Shallom for managing your to-do list. You can add, list, complete, and delete tasks conveniently through the command line.

## Usage

```bash
./todo-cli [options]
```

### Options:

- `-add`: Add a new task to the to-do list.
- `-list`: List all tasks in the to-do list.
- `-complete`: Mark a task as completed by providing the item number.
- `-del`: Delete a task by providing the item number.

### Environment Variables:

- `TODO_FILENAME`: Specify a custom filename for your to-do list. Default is ".todo.json".

## Examples:

1. **List all tasks:**

    ```bash
    ./todo-cli -list
    ```

2. **Add a new task:**

    ```bash
    ./todo-cli -add "Finish reading a book"
    ```

3. **Complete a task:**

    ```bash
    ./todo-cli -complete 2
    ```

4. **Delete a task:**

    ```bash
    ./todo-cli -del 3
    ```

## Note:

- When using the `-add` option, you can either provide the task as command line arguments or enter it through STDIN.
    - To add a task via STDIN:
      - ```bash
        echo "Go to the gym" | ./todo-cli -add
        ```
## Testing
This project includes unit and integration tests. To run the tests, execute the following command:

```bash
go test -v
```


## License

This tool is provided under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Copyrigt 2023 Micah Shallom