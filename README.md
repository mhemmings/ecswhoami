# ecswhoami

[![Test](https://github.com/mhemmings/ecswhoami/workflows/Test/badge.svg?branch=master)](https://github.com/mhemmings/ecswhoami/actions?query=workflow%3ATest)
[![GoDoc](https://godoc.org/github.com/apex/log?status.svg)](https://godoc.org/github.com/mhemmings/ecswhoami)


A Go package to do a task metadata lookup from inside AWS Elastic Container Service.

`Lookup()` calls the `ECS_CONTAINER_METADATA_URI` as defined in https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html, and returns the response.


## Usage:

```go
package main

import (
	"log"

	"github.com/mhemmings/ecswhoami"
)

func main() {
	meta, err := ecswhoami.Lookup()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Name:", meta.Name)
	log.Println("Image ID:", meta.ImageID)
	// ... etc
}
```
