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

## Building

For help about how to deal with the repo just run:

```bash
$ make
```

### Unit tests

```bash
$ make unit-test-local
```

### Binary

```bash
$ make build-local
```

### Image

```bash
$ make build
```

## Important Notes about implementation

* This authentication is a poors man auth (would rather do oauth2 with a dedicated authorizer). It's only here to allow the api to reject users accessing other user's game and to know which user owns the game being created.

* Since auth endpoint is not implemented, there is no token generation. This token is used by the api to know what user is making the call. Currently all game endpoints (not user management) ignore token and assume the call is being made by the same hardcoded user (id: e341410d-752a-404f-9acc-904764fd38f3). This hardcoded user was created in a migration.

* Randomizer for mine positioning might not be random enought. Should improve this.

* On board creation the mines are created, but it should be done such as to avoid the first click of the user, to make sure he doesn't loose on the first move. Not implemented in first release.

* In db the name used for user is accout because user is kind of a reserved word in postgresql.

* Pause endpoint wasn't finally implemented. The model does have everything needed for it to work.