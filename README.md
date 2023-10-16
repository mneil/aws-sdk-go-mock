# AWS SDK GO Mock

A complete mock implementation of the aws-sdk-go and aws-sdk-go-v2 clients.

## Contributing

### Requirements

 - go >= 1.20.5
 - node >= 20.x

I use node for development because it has some advantages for:

 1. writing automation scripts
 2. running git patches

I'm not a fan of Makefile and while I can easily traverse a folder and apply git patches it is far easier to use a known, tested piece of software like `patch-package`

### Set Up

Clone the repository and run

```
npm i
```
