# alphaunx-go
AlphaUnix-Go Scan is a powerful and flexible Host Scanner tool Built using the Go programming language, this tool is designed for fast and efficient network scanning, making it ideal for discovering active hosts and open ports on a network.

Here’s the updated README.md file with proper formatting and a copy button suggestion for code blocks. GitHub supports copy buttons for code blocks natively, so no additional setup is required. Just ensure the code blocks are properly formatted.

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

example.com
google.com
github.com

### Command

bash
go run alphaunx.go hosts.txt

### Output

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

### HCreate a New Repository on GitHubw Repository on GitHub**:
   - Log in to your GitHub account.
   - Click on "New" to create a new repository.
   - Name the repository, for example, alphaunx-go.
   - Choose the appropriate settings (public or private) and click "Upload the Code to the RepositoryCode to the Repository**:
   - Use Git to upload the code to the repository.

  
   git init
   git add .
   git commit -m "Initial commit"
   git branch -M main
   git remote add origin https://github.com/alphaunx/alphaunx-go.git
   git push -u origin main
   
3. **Create the README.md File**:
   - Create a README.md file in the root directory of the repository and add the Push the Changesve.

4. **Push the Changes**:
   - Commit and push the README.md file to GitHub.

  
   git add README.md
   git commit -m "Add README.md file with tool details"
   git push origin main
   
---

### Copy Button for Code Blocks

copy buttoncally adds a **copy button** to code blocks in the README.md file. Users can click the copy button to copy the commands directly.

For example:

go run alphaunx.go hosts.txt
The copy button will appear in the top-right corner of the code block when viewed on GitHub.

---

Now, your tool alphaunx.go is ready to be shared on GitHub with clear installation and usage instructions, along with version and developer information!
