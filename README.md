# gopher-port

A fast, concurrent port scanner written in Go. This programme allows you to scan single or multiple targets for open ports using Go's powerful concurrency features.

## Features

- Concurrent port scanning using goroutines
- Multiple target support
- Controlled concurrency with semaphores to prevent system resource exhaustion
- 2-second timeout for each port check
- Simple command-line programme interface
- Error handling and clean output

## Installation

Ensure you have Go installed on your system. Then:

```bash
# Clone the repository
git clone https://github.com/yourusername/gopher-port
cd gopher-port

# Build the binary
go build -o gopher-port
```

## Usage

Run the programme:

```bash
./gopher-port
```

You will be prompted to:
1. Enter target IP addresses (separated by commas for multiple targets)
2. Enter the number of ports you want to scan (from port 1 to your specified number)

Example:
```
[*] Enter Targets To Scan(split them by ,): 192.168.1.1,192.168.1.2
[*] Enter How Many Ports You Want To Scan: 1000
```

## Technical Details

- Uses a semaphore to limit concurrent scans to 100 at a time
- Each port scan has a 2-second timeout
- Employs Go's WaitGroup for proper goroutine synchronization
- Filters out common "connection refused" errors for cleaner output
- Properly handles multiple targets with whitespace trimming

## Example Output

```
[*] Scanning Multiple Targets

[*] Scanning target: 192.168.1.1
[*] Starting scan of 1000 ports
[+] Port 80 is open
[+] Port 443 is open
[*] Scan completed

[*] Scanning target: 192.168.1.2
[*] Starting scan of 1000 ports
[+] Port 22 is open
[+] Port 3306 is open
[*] Scan completed
```

## Performance

The scanner utilises Go's concurrency features for optimal performance:
- Concurrent port scanning with controlled parallelism
- Semaphore limiting to prevent overwhelming the system
- Non-blocking I/O operations

## Limitations

- The scanner only performs TCP connect scans
- Maximum concurrent scans limited to 100 to prevent resource exhaustion
- Does not support service version detection
- No custom port range specification (always starts from port 1)

## Contributing

Feel free to submit issues and pull requests. Some areas for potential optimisation:

1. Add custom port range support
2. Implement UDP scanning
3. Add service version detection
4. Add command-line flags for configuration
5. Improve error handling and reporting
6. Add progress bar for long scans

## License

MIT License - feel free to use and modify as needed.

## Disclaimer

This programme is intended for network administrators and security professionals to scan their own networks. Always ensure you have permission before scanning any networks or systems you don't own.