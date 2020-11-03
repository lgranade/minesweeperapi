# Minesweeper API

This is a dockerized Minesweeper API requested by deviget team as a code challenge.

## API Contract

The openapi spec is in `/design/contract.yaml`

There is an insomnia file to manually interact with the api in `/design/insomnia.json`

## Architecture

The api runs in an AWS EC2 micro instance (free tier) with docker installed.

The database is an RDS postgres instance (free tier too).

## Database Schema

Migrations are stored in `/sql/migrations`

## Build

For help about how to deal with the repo just run:

```bash
$ make
```

To build the image just run

```bash
$ make build
```

## Notes

* User should be taken from access token by design but first version has user id hardcoded.

* Randomizer for mine positioning might not be random enought. Should improve this.

* On board creation the mines are created, but it should be done such as to avoid the first click of the user, to make sure he doesn't loose on the first move. Not implemented in first release.

* In db the name used for user is accout because user is kind of a reserved word in postgresql.