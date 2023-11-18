# Word Counter

This simple command-line tool, developed by Micah Shallom, counts the number of words, lines, or bytes from an input source. It provides flexibility by allowing users to choose what to count.

## Usage

```bash
./word-counter [options]
```

### Options:

- `-l`: Count lines instead of words.
- `-b`: Count bytes.

By default, the tool counts words.

## Examples:

1. **Count Words:**

   ```bash
   ./word-counter
   ```

2. **Count Lines:**

   ```bash
   ./word-counter -l
   ```

3. **Count Bytes:**

   ```bash
   ./word-counter -b
   ```

## Note:

- If neither `-l` nor `-b` is provided, the tool counts words.

## License

This tool is provided under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Copyrigt 2023 Micah Shallom