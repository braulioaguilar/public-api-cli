# Public API CLI

Simple CLI to use [Public API for Public APIs](https://api.publicapis.org/).

More info about parameters [here](https://api.publicapis.org/).

## Endpoints

`entries`, `random`, `categories`

## Usage

```sh
$ go run cmd/main.go <endpoint> <filter> 

## Example
$ go run cmd/main.go entries --filter cors=yes,https=true
```

## Makefile

```sh
$ make build
```

## Bin

```sh
$ ./bin/puplicapi entries --filter cors=no
```