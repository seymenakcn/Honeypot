# Honeypot
Lightweight SSH honeypot in Go for observing basic attacker interactions.

# Go SSH Honeypot (Learning Project)

This project is a **simple TCP-based SSH honeypot** written in Go.  
It is designed for **learning purposes**, focusing on how network services,
concurrency, and basic attacker interaction work at a low level.

The honeypot does **not** provide a real SSH service.  
It only simulates an SSH server enough to attract automated scanners
and basic connection attempts.

---

## Purpose

- Learn how TCP servers work in Go
- Understand how real services handle multiple connections
- Observe how attackers interact with exposed services
- Practice goroutines and blocking I/O concepts

---

## How It Works

1. The program listens on a TCP port (`2222`)
2. When a client connects:
   - A fake SSH banner is sent
   - The client’s first input is read
3. Each connection is handled in a **separate goroutine**
4. The server keeps running and accepts new connections continuously

---

## Why Port 2222?

Port `2222` is commonly used as an alternative SSH port.
Using it avoids conflicts with a real SSH service running on port `22`
and makes testing safer.

---

## Features (Current)

- TCP listener using Go `net` package
- Fake OpenSSH banner
- Reads user input line-by-line
- Supports multiple simultaneous connections
- Minimal and easy-to-understand codebase

---

## Running the Honeypot

### Requirements
- Go 1.20+ recommended
- Linux / macOS (Windows works with minor adjustments)

### Run
```bash
go run honeypot.go
