package main

import (
	"github.com/cascades-fbp/cascades/library"
)

var registryEntry = &library.Entry{
	Description: "Add two numbers and supported type is float32",
	Elementary:  true,
	Inports: []library.EntryPort{
		library.EntryPort{
			Name:        "INA",
			Type:        "string",
			Description: "Input port for receiving IPs",
			Required:    true,
		},
		library.EntryPort{
			Name:        "INB",
			Type:        "string",
			Description: "Input port for receiving IPs",
			Required:    true,
		},
	},
	Outports: []library.EntryPort{
		library.EntryPort{
			Name:        "SUM",
			Type:        "all",
			Description: "Output port for captured submatching map in JSON",
			Required:    true,
		},
	},
}
