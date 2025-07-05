![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
# Cross-Language RPC System
A lightweight Remote Procedure Call (RPC) system enabling seamless communication between Go servers and TypeScript clients over TCP sockets.

## Why This Project?
Built to explore distributed systems concepts and demonstrate cross-language communication patterns used in microservices architectures.

## Features
- **Cross-Language Communication**: Go server and TypeScript client
- **TCP Socket Communication**: Direct, efficient network protocol
- **JSON-based Protocol**: Human-readable message format
- **Type Safety**: Strongly typed interfaces and structs
- **Interactive Client**: User-friendly command-line interface
- **Concurrent Server**: Handles multiple clients simultaneously using goroutines
- **Error Handling**: Robust error management and validation

## Tech Stack
- **Backend**: Go with native `net` package
- **Frontend**: TypeScript/Node.js with built-in networking
- **Protocol**: TCP sockets with JSON serialization
- **Concurrency**: Go goroutines for handling multiple clients

## Available RPC Functions
| Function | Parameters | Description | Example |
|----------|------------|-------------|---------|
| `floor` | `float` | Returns floor of a number | `floor(3.7)` returns `3` |
| `nroot` | `int, int` | Calculates nth root | `nroot(3, 27)` returns `3` |
| `reverse` | `string` | Reverses a string | `reverse("hello")` returns `"olleh"` |
| `validAnagram` | `string, string` | Checks if two strings are anagrams | `validAnagram("listen", "silent")` returns `true` |
| `sort` | `string[]` | Sorts array of strings | `sort(["c", "a", "b"])` returns `["a", "b", "c"]` |

## Quick Start
### Prerequisites
- Go 1.19+ 
- Node.js 16+
- npm

### 1. Clone the Repository
```bash
git clone https://github.com/ngthecoder/simple_rpc_system
cd simple_rpc_system
```

### 2. Start the Go Server
```bash
cd server
go run main.go
```
Server will start listening on `localhost:8090`

### 3. Set Up TypeScript Client
```bash
cd client
npm install
```

### 4. Run the Client
```bash
tsc client.ts
node client.js
```

### 5. Make RPC Calls
```
Enter method name: floor
Enter params (comma separated): 3.7
```

## Protocol Specification
### Request Format
```json
{
  "method": "floor",
  "params": [3.7],
  "param_types": ["float"],
  "id": 1
}
```

### Response Format
```json
{
  "result": "3",
  "result_type": "float64", 
  "id": 1
}
```

## Example Usage
**Client Input:**
```
Enter method name: nroot
Enter params (comma separated): 3, 27
```

**Network Communication:**
```json
// Request
{"method":"nroot","params":[3,27],"param_types":["int","int"],"id":42}

// Response  
{"result":"3","result_type":"float64","id":42}
```

**Output:**
```
Sending: { method: 'nroot', params: [3, 27], param_types: ['int', 'int'] }
Connection Established
{"result":"3","result_type":"float64","id":42}
```

## Project Structure
```
rpc-system/
├── server/
│   └── main.go              # Go RPC server
├── client/
│   ├── client.ts            # TypeScript RPC client
│   ├── package.json         # Node.js dependencies
│   └── package-lock.json    # Dependency lock file
├── README.md
└── .gitignore
```

## Technical Implementation
### Server (Go)
- **Concurrent Handling**: Each client connection spawned in separate goroutine
- **JSON Unmarshaling**: Automatic parsing using struct tags
- **Function Registry**: Map-based function lookup and execution
- **Type Safety**: Interface{} with type assertions for dynamic typing

### Client (TypeScript)
- **Interactive CLI**: Readline interface for user input
- **Smart Parameter Parsing**: Function-specific input handling
- **Type Definitions**: Interfaces matching server protocol
- **Error Handling**: Network and parsing error management

## Future Enhancements
- [ ] Authentication and authorization
- [ ] HTTP REST API wrapper
- [ ] Client connection pooling
- [ ] Message compression
- [ ] Async/await client API
- [ ] Load balancing support
- [ ] Metrics and monitoring

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Learning Outcomes
This project demonstrates:
- **Network Programming**: TCP socket implementation
- **Serialization**: JSON protocol design
- **Concurrency**: Goroutine-based server architecture  
- **Cross-Language Integration**: Communication between Go and TypeScript
- **System Design**: RPC architecture patterns
- **Error Handling**: Production-ready error management
