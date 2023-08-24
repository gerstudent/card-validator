# Introduction

Credit card number validator written in Go using the Luhn algorithm.

> The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, named after its creator, IBM scientist Hans Peter Luhn, is a simple checksum formula used to validate a variety of identification numbers.

# Prerequisites

- You need to have Go installed on your computer. The version used in this project is 1.20.

# Run

1. Clone this repo: `git clone https://github.com/gerstudent/card-validator`
2. To run this project locally, write the following command in the terminal: `go run main.go validation.go 8080`
3. Use Postman or Curl to send a POST request to the server at route "/"
4. Set the headers: Content-Type: application/json
5. Send a POST request to http://localhost:8080/
6. Set the request body to { "number": "4539 3195 0343 6467" }

- If all goes well, you should receive this response {"valid": true}
