# go_crypto_generator


## About

"go_crypto_generator" is an ui application for generate a string hash with cryptographics functions as :

- MD5
- SHA-1
- SHA-256
- SHA-512

And those who require a secret key :

- HMAC MD5
- HMAC SHA-1
- HMAC SHA-256
- HMAC SHA-515

Without any external library.


## Docker

1. `cd go/src/github.com/EtienneR/go_crypto_generator`
2. `sudo docker build -t go_crypto_generator .`
3. `sudo docker run --publish 3000:3000 --name crypto_generator --rm go_crypto_generator`

