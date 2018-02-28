# go-underarmour
> Golang library for interacting with the [Nest API](https://developers.nest.com/).

## Install

```
$ go get github.com/blaskovicz/go-nest
```

## Use

```go
import (
  nest "github.com/blaskovicz/go-nest"
)

// initialize a default client
client := nest.New()

```

See the [cmd/](cmd/) directory for examples of generating an access token
and using the client.

## Test

```
$ go test ./...
```
