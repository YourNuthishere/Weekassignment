# Go Password Cracking Labs

This repository contains all completed password cracking labs (Lab 1, Lab 2, Lab 3) as part of the course exercises. Each lab demonstrates how to crack password hashes (MD5, SHA1, SHA512) using Go.

---

## Lab 2 Task1: MD5 Password Cracker
- **Goal:** Crack an MD5 hash using a wordlist.
- **Target MD5:** `6a85dfd77d9cb35770c9dc6728d73d3f`
- **Files:**
  - `main.go` — main program for MD5 cracking
  - `utils/crack/md5.go` — helper functions
  - `nord_vpn.txt` — password candidates
  - `verbose.txt` — program output log


---

## Lab 2 task2: SHA1 Password Cracker
- **Goal:** Crack a SHA1 hash using a wordlist.
- **Target SHA1:** `aa1c7d931cf140bb35a5a16adeb83a551649c3b9`
- **Files:**
  - `main.go` — main program for SHA1 cracking
  - `utils/crack/sha1.go` — helper functions
  - `nord_vpn.txt`
  - `verbose.txt`
 

---

## Lab 3: SHA512 Password Cracker
- **Goal:** Crack a SHA512 hash using a wordlist.
- **Target SHA512:**  
`485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38`
- **Files:**
  - `main.go` — main program for SHA512 cracking
  - `utils/crack/sha512.go` — helper functions
  - `nord_vpn.txt`
  - `verbose.txt`
  - `outputscreenshot/` — screenshots showing program run

---

## all output screenshot are in outputscreenshot in folder lab3
## How to Run
```bash
go run main.go > verbose.txt
