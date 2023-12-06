# File Management Utility

This is a simple command-line utility for managing files in a directory. The utility provides various options for filtering, listing, archiving, and deleting files based on different criteria.

## Usage

```bash
./fileutil -root <directory_path> [options]
```

### Options:

- `-root`: Root directory to start file operations (default is the current directory).
- `-log`: Log deletes to this file (default is STDOUT).
- `-list`: List files only.
- `-del`: Delete files.
- `-archive`: Archive directory.
- `-ext`: File extension(s) to filter out (comma-separated).
- `-size`: Minimum file size.
- `-before`: Return files created before the specified date and time.
- `-after`: Return files created after the specified date and time.

## Examples

### List Files

```bash
./fileutil -root /path/to/directory -list
```

### Delete Files

```bash
./fileutil -root /path/to/directory -del
```

### Archive Files

```bash
./fileutil -root /path/to/directory -archive /path/to/archive
```

### Filter by Extension and Size

```bash
./fileutil -root /path/to/directory -ext .txt,.pdf -size 1024
```

### Filter by Date

```bash
./fileutil -root /path/to/directory -before "2023-01-01 12:00:00" -after "2022-01-01 12:00:00"
```

## Notes

- The utility provides flexible options for file management based on your needs.
- Use caution when using the delete (`-del`) option, as it permanently removes files.
- Logging can be redirected to a file using the `-log` option.

Feel free to explore and customize the utility based on your specific requirements.