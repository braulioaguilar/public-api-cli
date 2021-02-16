# Public API CLI

Simple CLI to use [Public API for Public APIs](https://api.publicapis.org/).

More info about parameters [here](https://api.publicapis.org/).

## Endpoints

* `/categories`: List all categories
* `/entries`: List all entries currently cataloged in the project
* `/random`: List a single entry selected at random

## Test code

```sh
$ go run cmd/*.go <endpoint> <filter> 

## Example
$ go run cmd/*.go entries --filter cors=yes,https=true
```

## Generate binary

```sh
$ make build
```

## Usage binary

```sh
$ ./bin/publicapi entries --filter cors=no
```