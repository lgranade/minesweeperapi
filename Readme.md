# Minesweeper API

This is a dockerized Minesweeper API requested by deviget team as a code challenge.

## API Contract

The openapi spec is in `/design/contract.yaml`

There is an insomnia file to manually interact with the api in `/design/insomnia.json`

## Architecture

The api runs in an AWS EC2 micro instance (free tier) with docker installed.

The database is an RDS postgres instance (free tier too).

## Database Schema

Migrations are stored in `/migrations`

## Build

For help about how to deal with the repo just run:

```bash
$ make
```

To build the image just run

```bash
$ make build
```
