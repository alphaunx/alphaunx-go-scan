# Alphaunx (alphaunx-go)
AlphaUnx is a powerful, fast, and efficient network scanning toolkit written in Go. It provides a comprehensive suite of tools for server and network analysis with an easy-to-use interface. The current version is v1.5 Stable .


## Features
AlphaUnx includes eight specialized tools for different network scanning and analysis tasks:

1. Host Scanner - Scan large batches of domains to check their status and server information
2. Port Scanner - Discover open ports on servers and devices
3. Subdomain Finder - Search for subdomains of a specific domain using multiple sources
4. CIDR Scanner - Scan IP ranges using CIDR notation
5. IP Domain Lookup - Discover domains hosted on a specific IP address
6. CDN Finder - Identify CDN providers for domains
7. Ping Tool - Check host or IP connectivity
8. Update Checker - Stay up-to-date with the latest tool versions
## Requirements
- Go 1.16 or higher
- Internet connection for advanced features like subdomain finding
- Git (for update functionality)
## Installation
### On Termux
1. Install Go:
   
   ```
   pkg install golang
   ```
2. Install the tool using wget:
   
   ```
   wget https://raw.githubusercontent.com/alphaunx/alphaunx-go-scan/refs/heads/main/aunx.go
   ```
3. Run the tool:
   
   ```
   go run aunx.go
   ```
### On Kali Linux
1. Install Go:
   
   ```
   sudo apt-get update
   sudo apt-get install 
   golang
   ```
2. Install the tool using wget:
   
   ```
   wget wget https://raw.githubusercontent.com/alphaunx/alphaunx-go-scan/refs/heads/main/aunx.go
   ```
3. Run the tool:
   
   ```
   go run aunx.go
   ```
## Usage
1. Run the tool:
   
   ```
   go run aunx.go
   ```
2. Select the appropriate mode from the main menu:
   
   - [1] Host Scanner - Scan domains from a file or single domain
   - [2] Port Scanner - Scan ports for IP addresses/domains
   - [3] Subdomain Finder - Find subdomains for a domain
   - [4] CIDR Scanner - Scan CIDR ranges
   - [5] IP Domain Lookup - Collect domains from an IP address
   - [6] CDN Finder - Detect CDN networks
   - [7] Ping Tool - Check connectivity
   - [8] Update Checker - Check for updates
3. Follow the on-screen instructions for each mode.
## Advanced Features
- Multitasking : The tool supports concurrent operations to speed up scanning processes with configurable worker counts (10-500)
- Result Saving : Ability to save scan results to text files
- Colored Interface : User-friendly interface with ANSI colors for enhanced user experience
- Progress Indicators : Display progress bars for long scanning operations
## Tool-Specific Features
### Host Scanner
- Supports single domains, comma-separated lists, or input from files
- Displays status codes and server information for active hosts
- Real-time progress tracking
### Port Scanner
- Scans common ports on target hosts
- Concurrent connections with configurable worker count
- Immediate display of open ports
### Subdomain Finder
- Real-time display of found subdomains
- Progress spinner during search
### CIDR Scanner
- Generates and scans all IPs in a CIDR range
- Concurrent scanning with progress bar
- Displays active hosts with status codes
### IP Domain Lookup
- Searches multiple online sources for domains associated with an IP
- Parallel querying using goroutines
- Filters out invalid or resource-related domains
### CDN Finder
- Identifies CDN providers by analyzing HTTP headers and CNAME records
- Supports a wide range of CDN patterns
- Displays detected CDNs and relevant headers
### Ping Tool
- Checks host connectivity using system ping commands
- Colored output based on success or failure
### Update Checker
- Checks for updates from GitHub repository
- Verifies Git installation
- Provides force update option
## Notes
- The tool uses HTTP HEAD requests to minimize data transfer
- Internet connection is required for features like subdomain finding and IP domain lookup
- Results are saved to /sdcard/ on Android/Termux or the current directory on other systems
- This is a v1.5 Stable release that is regularly updated
## Developer
- Developer : ᴅᴀʀᴋ 404 - ɴᴏᴛ ꜰᴏɴᴜᴅ
- Version : v1.5 Stable
## Contact and Issue Reporting
For reporting any issues or inquiries, please contact via:

- Telegram Channel: t.me/alphaunx
- Technical Reports: t.me/l300e
## Disclaimer
This tool is intended for legal and ethical use only. The developer is not responsible for any illegal use of the tool.
