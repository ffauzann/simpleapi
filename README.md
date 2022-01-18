# simpleapi
A simple API contains login and read data with Go.

## Overview
This project has **only** been tested on ParrotOS 4.11 amd64 (Debian Buster) with Go 1.17 and MySQL 8.0.27 installed.
Kindly submit new issue if there is an error occurred in your machine.

## Installation
For **Linux** or **OSX** user, use these following steps:
- Clone the repository.
- Make a copy of a file `internal/config/config.yaml.example`, name it as `config.yaml`.
- Open `config.yaml`, configure the database credentials with your own.
- Execute sql script(s) in `migration` directory manually.
- Run `make init`.
- Run `swag init`.

For **Windows** user, use the same steps as mentioned above. Since the machine won't recognized `make` command, you might need to copy each commands manually on `Makefile` file at the `init` section.

## Usage
- Start the server with `make run` or `air` command for development purpose.
- Open `localhost:3000/swagger/` in your browser to start testing.