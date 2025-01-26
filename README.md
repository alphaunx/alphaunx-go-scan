GitHub automatically adds a copy button to code blocks in the README.md file. If the copy button is not appearing, it might be due to improper formatting of the code blocks. Below is the corrected and properly formatted README.md file to ensure the copy button appears:

---

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

   
bash
   pkg install golang
  

2. Download the code:

   
bash
   git clone https://github.com/alphaunx/alphaunx-go.git
   cd alphaunx-go
  

3. Run the tool:

   
bash
   go run alphaunx.go hosts.txt
  

---

### On Kali Linux

1. Install Go:

   
bash
   sudo apt-get update
   sudo apt-get install golang
  

2. Download the code:

   
bash
   git clone https://github.com/alphaunx/alphaunx-go.git
   cd alphaunx-go
  

3. Run the tool:

   
bash
   go run alphaunx.go hosts.txt
  

---

## Usage

1. Create a `hosts.txt` file containing a list of domains you want to scan, with each domain on a separate line.

2. Run the tool:

   
bash
   go run alphaunx.go hosts.txt
  

3. Enter the number of workers (between 10 and 500) when prompted.

4. The tool will scan the domains and display the results, including the server type, status code, and response time.

---

## Example

### `hosts.txt` File

plaintext
example.com
google.com
github.com

### Command

bash
go run alphaunx.go hosts.txt

### Output

plaintext
[✓] example.com [200] Nginx
[✓] google.com [200] Google Frontend
[✓] github.com [200] GitHub.com

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

### Steps to Ensure Proper Code Block Formatting Code Block Formatting**:
   - Use triple backticks (
   - Specify the language after the opening backticks (e.g., `
```bash ```` for Bash commands or
   Example:
   
markdown
   ```bash
   go run alphaunx.go hosts.txt
  
   
2. **Push the Updated README.md to GitHub**:
   - After making the changes, push the updated README.md file to GitHub.

  
   git add README.md
   git commit -m "Fix code block formatting for copy button"
   git push origin main
   
3. **Check the Copy Button**:
   - Once the file is updated on GitHub, the copy button should appear automatically in the top-right corner of each code block.

---

### Why the Improper Formattingppear

- **Improper Formatting**: If the code block is not wrapped in triple backticks or the language is not specified, GitHub may not Caching Issuescode block.
- **Caching Issues**: Sometimes, GitHub caches the README.md file. Wait a few minutes or refresh the page.

---

Now, the README.md filecopy buttonrmatted, and the **copy button** should appear next to all code blocks when viewed on GitHub.
