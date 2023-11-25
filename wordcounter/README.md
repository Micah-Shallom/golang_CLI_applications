## Word Counter

This is a simple command-line utility in Go for counting lines, words, or bytes in a given text file or from standard input. The tool provides flexibility in the choice of operation, making it versatile for different use cases.

## Usage

The tool accepts various command-line flags:

- `-l`: Count lines.
- `-b`: Count bytes.
- `-file`: Specify the file to process.

You can also provide multiple filenames as arguments to process multiple files.

### Examples

#### Count Lines in a File
```bash
./wordcounter -l -file example.txt
```

#### Count Words from Standard Input
```bash
echo "This is a sample text" | ./wordcounter
```

#### Count Bytes in Multiple Files
```bash
./wordcounter -b file1.txt file2.txt
```

## How It Works

1. **Reading Files**: The tool reads content from either the specified file(s) or from standard input if no files are provided.

2. **Counting Operation**: Based on the flags provided, the tool counts lines, words, or bytes using a scanner that reads from the input source.

3. **Displaying Results**: The results, along with the type of operation performed, are displayed for each file processed.

## Build and Run

To build the tool, use the `go build` command:

```bash
go build -o wordcounter
```

After building, you can run the tool as described in the examples above.

## Contribution

Feel free to contribute to this project by submitting issues or pull requests. Your suggestions and improvements are welcome.

Happy counting!