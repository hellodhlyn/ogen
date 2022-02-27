# OGen

OpenGraph image generator. Inspired by [vercel/og-image](https://github.com/vercel/og-image) but added supports for CJK text.

## Usage

### Run server

You can use the docker image.

```shell
docker run -p 8080:8080 gchr.io/hellodhlyn/ogen:latest
```

Add `--platform linux/x86_64` if you are using macOS devices with Apple Silicon chip.

### Make a request

`http://127.0.0.1:8080?title=<title>&author=<name>&profile_image=<url>`

## Development

### Prerequisites

* Golang 1.17
* NodeJS (tested on NodeJS 16 or greater)
* svgexport

### Run server locally

```shell
go run ./cmd/server
```
