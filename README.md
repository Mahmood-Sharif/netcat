# net-cat

## Overview
`net-cat` is a TCP-based client–server chat application written in Go.  
It implements a simplified NetCat-like server that supports multiple concurrent clients communicating through a shared chat room.

The project focuses on TCP networking, concurrency, and connection lifecycle management.

---

## Features
- TCP server accepting multiple concurrent client connections
- Client name registration with validation
- Broadcast messaging between all connected clients
- Message timestamps and sender identification
- Chat history synchronization for newly connected clients
- Join and leave notifications for all participants
- Connection limit enforcement (maximum 10 clients)
- Graceful handling of client disconnects
- Default port fallback when no port is specified

---

## Architecture
- Server–client model using TCP sockets
- Goroutines for handling concurrent client connections
- Channels and mutexes for synchronized message handling
- Central message dispatcher to broadcast events and messages

---

## Technologies
- Go
- TCP networking (`net` package)
- Concurrency (goroutines, channels, mutexes)
- Buffered I/O
- Time-based message formatting

---

## Usage

### Start the server
```bash
go run .

### Starts the server on the default port 8989.

go run . 2525


### Starts the server on a custom port.

### If invalid arguments are provided:

[USAGE]: ./TCPChat $port

### Connect a client

### Use NetCat from another terminal:

nc localhost 8989

### Upon connection, the server prompts for a non-empty username before joining the chat.

### Message Format:

### Messages are broadcast to all clients using the following format:

[YYYY-MM-DD HH:MM:SS][username]: message

### Notes:

The project is intended as a learning and demonstration tool.

It does not aim to fully replace the standard nc utility.

The focus is on correctness, concurrency safety, and connection management.