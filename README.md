This is a rewrite of [keratin/authn](https://github.com/keratin/authn) from Ruby to Go. It is a work in progress and not ready for usage.

# Getting Started

## Dependencies

* Redis
* SQLite3

## Configuration

All configuration is through environment variables. Please see [docs](https://github.com/keratin/authn/blob/master/docs/config.md) for details.

## Running Migrations

To run migrations for the configured database:

    > authn-server migrate

This command will determine which migrations need to be run and will smartly converge the database in a no-downtime production-quality fashion.

# Developing

Install [glide](https://github.com/Masterminds/glide#install).

Run `glide install` to populate your vendor/ directory.

Run `go install` to compile the vendored dependencies. I think.

Run a Redis server somewhere.

Create a `.env` file in the project root, and add environment variables there as necessary.

# COPYRIGHT & LICENSE

Copyright (c) 2017 Lance Ivy

Keratin AuthN is distributed under the terms of the GPLv3. See [LICENSE-GPLv3](LICENSE-GPLv3) for details.