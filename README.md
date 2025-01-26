If the **copy button** is not appearing next to the code blocks in your GitHub `README.md` file, it is likely due to **improper formatting** of the code blocks. GitHub automatically adds a copy button to properly formatted code blocks. Below are the steps to fix this issue and ensure the copy button appears for all embedded commands:

---

### **Steps to Fix the Issue**

1. **Use Proper Code Block Formatting**:
   - Ensure all code blocks are wrapped in **triple backticks** (```` ``` ````).
   - Specify the **language** after the opening backticks (e.g., ````bash```` for Bash commands or ````plaintext```` for plain text).

   Example:
   ```markdown
   ```bash
   go run alphaunx.go hosts.txt
   ```
   ```

2. **Check for Consistent Formatting**:
   - Make sure every code block in your `README.md` file follows the same formatting rules.
   - Avoid mixing spaces or tabs in the formatting.

3. **Push the Updated README.md to GitHub**:
   - After fixing the formatting, push the updated `README.md` file to your GitHub repository.

   ```bash
   git add README.md
   git commit -m "Fix code block formatting for copy button"
   git push origin main
   ```

4. **Clear GitHub Cache (if needed)**:
   - Sometimes, GitHub caches the `README.md` file. To ensure the changes are reflected:
     - Wait a few minutes.
     - Refresh the GitHub page.
     - If the issue persists, try opening the repository in an incognito/private browser window.

---

### **Example of Correctly Formatted README.md**

```markdown
# AlphaUnx (alphaunx.go)

AlphaUnx (`alphaunx.go`) is a fast and efficient tool for scanning servers and extracting information about them. The tool is written in Go and supports concurrent scanning using workers. It is currently in **Beta v1.0** and is actively maintained.

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

1. Create a `hosts.txt` file containing a list of domains you want to scan, with each domain on a separate line.

2. Run the tool:

   ```bash
   go run alphaunx.go hosts.txt
   ```

3. Enter the number of workers (between 10 and 500) when prompted.

4. The tool will scan the domains and display the results, including the server type, status code, and response time.

---

## Example

### `hosts.txt` File

```plaintext
example.com
google.com
github.com
```

### Command

```bash
go run alphaunx.go hosts.txt
```

### Output

```plaintext
[✓] example.com [200] Nginx
[✓] google.com [200] Google Frontend
[✓] github.com [200] GitHub.com
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
```

---

### **Why the Copy Button Might Not Appear**

- **Improper Formatting**: If the code block is not wrapped in triple backticks or the language is not specified, GitHub may not recognize it as a code block.
- **Caching Issues**: GitHub sometimes caches the `README.md` file. Wait a few minutes or refresh the page.
- **Mixed Formatting**: Avoid mixing spaces or tabs in the formatting, as this can break the code block rendering.

---

### **Final Notes**

After following the steps above, the **copy button** should appear next to all code blocks in your `README.md` file when viewed on GitHub. If the issue persists, double-check the formatting and ensure there are no hidden characters or syntax errors.
