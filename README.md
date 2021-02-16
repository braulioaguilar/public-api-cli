# Public API CLI

Simple CLI to use [Public API for Public APIs](https://api.publicapis.org/).

More info about parameters [here](https://api.publicapis.org/).

## Endpoints

`entries`, `random`, `categories`

## Test code

```sh
$ go run cmd/*.go <endpoint> <filter> 

## Example
$ go run cmd/*.go entries --filter cors=yes,https=true
```

## Makefile

```sh
$ make build
```

## Bin

```sh
$ ./bin/publicapi entries --filter cors=no
```