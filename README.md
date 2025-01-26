Here’s the text rewritten in a clean and consistent format in English:

---

# AlphaUnx (alphaunx.go)

**AlphaUnx** (`alphaunx.go`) is a fast and efficient tool for scanning servers and extracting information about them. Written in Go, it supports concurrent scanning using workers. It is currently in **Beta v1.0** and is actively maintained.

---

## Requirements

- **Go 1.16 or higher**
- A file containing a list of domains to scan (e.g., `hosts.txt`)

---

## Installation

### On Termux

1. Install Go:

   ```bash
   pkg install golang
   ```

2. Installing using wget :

   ```bash
    https://raw.githubusercontent.com/alphaunx/alphaunx-go/refs/heads/main/alphaunx.go
   ```

3. Run the tool:

   ```bash
   go run alphaunx.go hosts.txt
   ```

---

### On Kali Linux

1. Install Go:

   ```bash
   sudo apt-get update
   sudo apt-get install golang
   ```

2. Download the code:

   ```bash
   git clone https://github.com/alphaunx/alphaunx-go.git
   cd alphaunx-go
   ```

3. Run the tool:

   ```bash
   go run alphaunx.go hosts.txt
   ```

---

## Usage

1. Create a `hosts.txt` file containing a list of domains you want to scan, each on a separate line.

2. Run the tool:

   ```bash
   go run alphaunx.go hosts.txt
   ```

3. Enter the number of workers (between 10 and 500) when prompted.

4. The tool will scan the domains and display the results, including the server type, status code, and response time.

---

## Example

### `hosts.txt` File

```
example.com
google.com
github.com
```

### Running the Tool

```bash
go run alphaunx.go hosts.txt
```

---

## Notes

- The tool uses HTTP `HEAD` requests to minimize data transfer.
- Ensure the `hosts.txt` file is in the same directory as the script.
- This is a **Beta v1.0** release. Please report any issues or bugs.

---

## Credits

- **Developer**: Dark503
- **Version**: Beta v1.0
- **Telegram**: @1300e

---

Now, everything is clearly formatted and in English!
