# simpleapi
A simple API contains login and read data with Go.

## Overview
This project has been tested on ParrotOS 4.11 amd64 (Debian Buster) with Go 1.17 and MySQL 8.0.27 installed.

## Installation
For **Linux** or **Mac** user, use these following steps:
- Clone the repository.
- Make a copy of a file `internal/config/config.yaml.example`, name it as `config.yaml`.
- Open `config.yaml`, configure the database credentials with your own.
- Run `make init`.
- Start the server with `make run` or `air` for development purpose.

For **Windows** user, since the machine won't recognized `make` command. You might need to copy each commands manually on `Makefile` file at the `init` section.

## Usage
- Start the server with `make run` or `air` command for development purpose.
- Open `localhost:3000/swagger/` in your browser.