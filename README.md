[![GitHub release](https://img.shields.io/github/release/leonelquinteros/hubspot.svg)](https://github.com/leonelquinteros/hubspot)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/leonelquinteros/hubspot?status.svg)](https://godoc.org/github.com/leonelquinteros/hubspot)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/hubspot)](https://goreportcard.com/report/github.com/leonelquinteros/hubspot)

# HubSpot API client SDK for Go

## Install

```
go get github.com/leonelquinteros/hubspot
```

## Quickstart

```go

package main

import "github.com/leonelquinteros/hubspot"

func main() {
    // Create client object with config from environment variables (HUBSPOT_API_HOST, HUBSPOT_API_KEY, HUBSPOT_OAUTH_TOKEN)
    c := NewClient(NewClientConfig())

    // Create new contact
    data := ContactsRequest{
		Properties: []Property{
			Property{
				Property: "email",
				Value:    "contact@example.com",
			},
			Property{
				Property: "firstname",
				Value:    "Contact",
			},
			Property{
				Property: "lastname",
				Value:    "Example",
			},
		},
    }
    r, err := c.Contacts().Create(data)
	if err != nil {
		log.Fatal(err)
	}

    // Get contact by email
    contact, err = c.Contacts().GetByEmail("contact@example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Print contact object
    fmt.Printf("%+v", contact)
}

```

