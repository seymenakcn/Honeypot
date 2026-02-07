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
   - A login: prompt is displayed and the username is read
   - A password: prompt is displayed and the password is read
3. The captured credentials are logged together with:
   -Source IP address
   -Timestamp
4. Each connection is handled in a **separate goroutine**
5. The server keeps running and accepts new connections continuously

---

## Features (Current)

- TCP listener using Go `net` package
- Fake OpenSSH banner
- Simulated 'login' and 'password' prompts
- Captures username and password attempts
- Logs events to:
  -Terminal(real-time)
  -Log File ('honeypot.log')
- Reads user input line-by-line
- Supports multiple simultaneous connections
- Minimal and easy-to-understand codebase

---

## Running the Honeypot

```bash
go run honeypot.go
