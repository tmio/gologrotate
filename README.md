[![Build status](https://travis-ci.org/tmio/gologrotate.svg)](https://travis-ci.org/tmio/gologrotate)
[![Join the chat at https://gitter.im/tmio/gologrotate](https://badges.gitter.im/tmio/gologrotate.svg)](https://gitter.im/tmio/gologrotate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

# Gologrotate

This go program finds all .log files under a directory recursively, and rotate the file contents in .gz files next to them.

## Usage

`gologrotate [-now] <searchDir>`

`-now`: optional, runs a one-time log rotate. If this parameter is omitted, gologrotate runs at midnight.

`searchDir` : the directory which logrotate will search recursively for .log files.

# Contribute

## Develop

* `make build` builds the binary gologrotate.
* Tests come in 2 flavors:
** `make test` runs the tests
** `make it` runs the tests, builds the Docker image locally and tests it too.

## Get in touch

[![Join the chat at https://gitter.im/tmio/gologrotate](https://badges.gitter.im/tmio/gologrotate.svg)](https://gitter.im/tmio/gologrotate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

# Credits

Original code from [copytruncate](https://github.com/jamesandariese/copytruncate) by James Andariese

