# PortScanner

**Concurrent TCP port scanner in Go**

[![CI](https://github.com/ispectr3/portscanner/actions/workflows/ci.yml/badge.svg)](https://github.com/ispectr3/portscanner/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go: 1.21+](https://img.shields.io/badge/Go-1.21%2B-00ADD8.svg?logo=go)](https://golang.org)

PortScanner is a command-line utility that attempts TCP connections to a host or a CIDR range. It uses a configurable goroutine worker pool, reports responsive ports in the terminal, and can write the results to JSON or CSV.

## Features

- **TCP Connect scanning:** uses `net.DialTimeout`; this is not a SYN or stealth scanner.
- **Concurrent workers:** `-threads` controls the worker count and defaults to 100.
- **Host selection:** accepts a single host/IP or a CIDR range. For ranges with more than two addresses, the network and broadcast addresses are omitted.
- **Port selection:** accepts comma-separated ports (`80,443`), ranges (`1-1024`), or the built-in `top100` alias.
- **Opportunistic banners:** reads an initial response from an open port. For ports 80, 443, and 8080, it falls back to a generic HTTP `HEAD` request when the first read is empty.
- **Output:** always prints open-port results to the terminal; `-output` can additionally write JSON or CSV.

## Requirements

- Go 1.21 or newer
- No third-party Go modules are declared in `go.mod`.

## Installation

```bash
git clone https://github.com/ispectr3/portscanner.git
cd portscanner
go build -o portscanner main.go
```

## Usage

```bash
./portscanner -host <target> [flags]
```

| Flag | Default | Description |
| --- | --- | --- |
| `-host` | empty (required) | Host, IP, or CIDR range to scan. |
| `-ports` | `top100` | Comma-separated ports, ranges, or the built-in list selected by `top100`. |
| `-threads` | `100` | Number of concurrent worker goroutines. |
| `-timeout` | `1000` | Per-connection timeout in milliseconds. |
| `-format` | `json` | File format used with `-output`: `json` or `csv`. |
| `-output` | empty | Optional destination file. |

### Examples

```bash
# Scan one host using the built-in common-port list
./portscanner -host 192.168.1.10 -ports top100

# Scan a range of ports with 250 workers
./portscanner -host 10.0.0.1 -ports 1-1000 -threads 250

# Scan a CIDR range and write JSON
./portscanner -host 192.168.1.0/24 -ports 80,443,8080 -format json -output results.json
```

## Tests and CI

Run the repository's Go tests with:

```bash
go test -v ./...
```

The included GitHub Actions workflow runs these tests and builds the binary with Go 1.21 and 1.22. The current unit tests cover port-list parsing in `utils/utils_test.go`.

## Repository layout

| Path | Role |
| --- | --- |
| `main.go` | CLI flags, input parsing, scan orchestration, and output dispatch. |
| `scanner/scanner.go` | Worker pool, TCP connections, and result collection. |
| `scanner/banner.go` | Initial-read and HTTP fallback banner handling. |
| `output/output.go` | Terminal, JSON, and CSV output. |
| `utils/utils.go` | Host/port parsing, common-port list, and terminal banner. |
| `utils/utils_test.go` | Unit tests for port parsing. |
| `.github/workflows/ci.yml` | Go test and build checks on pushes and pull requests to `main`. |

## Limitations

- The scanner completes a TCP three-way handshake; it does not implement raw-socket SYN/stealth scanning.
- The `top100` name selects a hard-coded short list of common ports; the current list contains 20 entries, not 100.
- Only open ports are appended to the result slice. Closed connection attempts are not included in terminal, JSON, or CSV output.
- Banner collection is best-effort: it reads at most 1024 bytes, does not identify arbitrary protocols, and only has the HTTP fallback described above.
- Port, worker-count, and timeout bounds are not comprehensively validated by the CLI. Use valid ports and positive worker/timeout values.
- The repository does not currently include scanner or output-package tests; the included tests cover parsing utilities.

## Authorized use

> Use this tool only against systems and networks you own or are explicitly authorized to assess. Port scanning can trigger alerts, affect services, or violate policy and law when performed without permission.

## License

PortScanner is released under the [MIT License](LICENSE).
