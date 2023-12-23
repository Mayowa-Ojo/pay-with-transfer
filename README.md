# Pay With Transfer

> This is an attempt to implment a simple version of the 'Bank Transfer' payment option you'll see in Paystack sdks and Interswitch's Quickteller.

![](https://res.cloudinary.com/devmayor/image/upload/w_0.5,c_scale/v1703369164/u0aykovwfpv7ow97xshu.png)
![](https://res.cloudinary.com/devmayor/image/upload/w_0.5,c_scale/v1703369163/nxkbwf8tvtetfddci2we.png)

### How it works

This payment method leverages the instant bank settlement infrastructure that's available in Nigeria. The service provider generates an ephemeral bank account (read virtual account) that is only valid for a single payment (based on the amount) and timeframe. You then go to make a transfer from your bank to the generated account, after which the service provider confirms that transaction has been received.

### Architecture

This demo is built on a Go backend with postgres DB layer and temporal for asynchronous job execution. There's no real 3rd party service implementation for the payments and accounts so it's 100% based on mocked data.

### Installation

Requirements:
- Go (1.16+)
- Node.js (v16+)
- Npm/Yarn/Pnpm
- Docker (optional)
- Temporal (In-memory)

- Clone the repository
```bash
$ git clone https://github.com/Mayowa-Ojo/pay-with-transfer
```
- Install client dependencies
```bash
$ cd client
$ yarn
$ touch .env
```
- Install server dependencies
```bash
$ cd server
$ go mod tidy
$ make setup
$ touch .env
```
- Start containers
```bash
$ docker compose up
```
- Run migration
```bash
$ make migrate-up
```
- Start temporal server
```bash
$ make temporal-start
$ make temporal-ns-create
```
- Start the app
```bash
$ make run
$ cd ../client
$ yarn dev
```
