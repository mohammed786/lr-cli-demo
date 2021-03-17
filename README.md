# Introduction

This is the LoginRadius CLI project

## How to Build from Source

### Installation from source

0. Verify that you have Go 1.13+ installed

   ```sh
   $ go version
   ```

   If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).

1. Clone this repository

   ```sh
   $ git clone https://github.com/LoginRadius/lr-cli.git
   $ cd lr-cli
   ```

2. Build and install

   #### Unix-like systems
   ```sh
   # installs to '/usr/local' by default; sudo may be required
   $ make install
   
   ```

## List of Commands supported in Beta

```
lr
    - help
    - register
    - login
    - get
        - servertime
        - config
        - social <provider>
        - domain
        - account
        - sites
        - theme
    - add
        - social <provider>
        - domain
        - account
    - delete
        - social <provider>
        - domain
        - account --uid, --email
    - set
        - social <provider>
        - domain
        - account --uid
        - account-email --uid
        - account-phone --uid
        - sec-ques --uid
        - theme
    - verify
        - email <email>
        - username <username>
        - resend <username>, <email>
        - invalidate 
    
    - get-password 
    - set-password
    - reset-secret
```   