```markdown
# Markdown Preview Tool

This tool allows you to preview Markdown files by converting them to HTML and displaying them in your default web browser. It also provides an option to customize the output using an alternate template.

## Usage

```sh
markdown-preview -file <filename> [-s] [-t <template_file>]
```

- `-file`: Markdown file to preview (required).
- `-s`: Skip auto-preview in the default web browser (optional).
- `-t`: Specify an alternate template file (optional).

## Example

```sh
markdown-preview -file example.md
```

This command converts the `example.md` file to HTML, opens it in the default web browser, and deletes the temporary HTML file after a short delay.

## Customizing Output

You can customize the output by providing an alternate template using the `-t` flag. Create a template file with the desired structure and pass it to the tool.

```sh
markdown-preview -file example.md -t custom_template.html
```

## Installation

Ensure you have Go installed, then run:

```sh
go install github.com/username/markdown-preview
```

Replace `github.com/username/markdown-preview` with the actual repository path.

## License

This tool is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
