
# Golit

This repository implements a simple TCP listener in Go that listens for incoming connections on a specified port and prints the received data. It also includes graceful shutdown functionality to handle termination signals (`Ctrl+C`) cleanly.

**Features:**

- Accepts incoming TCP connections on a user-defined port.
- Reads data from connected clients.
- Prints received data to the console.
- Handles `Ctrl+C` gracefully by closing connections before exiting.

**Requirements:**

- Go 1.17 or later ([https://go.dev/](https://go.dev/))

**Getting Started:**

1. Clone this repository:
```bash
git clone https://github.com/antomfdez/golit.git
```

2. Navigate to the project directory:
```bash
cd golit
```

3. Run the listener:
```bash
go run ./cmd/golit -p 8080
```

- Replace 8080 with your desired port number.
**Usage:**

- The script accepts a single optional flag:

- -p: The port number on which to listen (default: 8080).

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
