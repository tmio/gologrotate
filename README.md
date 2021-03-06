# Gologrotate

This go program finds all .log files under a directory recursively, and rotate the file contents in .gz files next to them.

[![Build status](https://travis-ci.org/tmio/gologrotate.svg)](https://travis-ci.org/tmio/gologrotate)
[![Join the chat at https://gitter.im/tmio/gologrotate](https://badges.gitter.im/tmio/gologrotate.svg)](https://gitter.im/tmio/gologrotate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![DockerHub Badge](http://dockeri.co/image/tmio/gologrotate)](https://hub.docker.com/r/tmio/gologrotate/)

## Usage

`gologrotate [-now] [-time <time>] [-format <format>] <searchDir>`

`-now`: optional, runs a one-time log rotate. If this parameter is omitted, gologrotate runs as a cron job.

`-time <time>`: optional, sets the time of day according to the local clock at which the cron job will run. Defaults to 23:55.

`-format <format>`: optional, sets the format of the timestamp used to suffix the gzipped log file. Defaults to 2006-01-02 (YYYY-MM-DD).
See [Go time formats](https://golang.org/src/time/format.go).

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

# License

Copyright 2016 Antoine Toulme

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.