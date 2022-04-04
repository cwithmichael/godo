[![Go Report Card](https://goreportcard.com/badge/github.com/cwithmichael/godo)](https://goreportcard.com/report/github.com/cwithmichael/godo)
# Godo

A basic Todo app written with Go. Based on the Let's Go book by Alex Ewards.

## Example screenshot

<img width="1680" alt="Screen Shot 2021-08-22 at 3 55 18 PM" src="https://user-images.githubusercontent.com/1703143/130369834-49980373-4648-49ec-87a0-f23012592314.png">



## How to Run

The easiest way to run this is with `docker-compose`. Please see the [official documentation](https://docs.docker.com/compose/install/) for instructions on how to install it on your machine.

1. Generate a self-signed TLS certificate

    We'll use the `generate_cert.go` tool included with Go installations. Run these commands from inside the root directory of this project.
    ```
    $ mkdir tls
    $ cd tls
    # On Linux:
    $ go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    # On Mac (assuming you installed Go with brew):
    $ go run /usr/local/Cellar/go/<version>/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    # On Windows: 🤷
    ```
2. `docker-compose up`
3. Go to http://localhost:4000 in your web browser
