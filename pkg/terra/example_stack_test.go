// Copyright (c) Volvo Car AB
// SPDX-License-Identifier: Apache-2.0

package terra_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/volvo-cars/lingon/pkg/terra"
	"golang.org/x/exp/slog"
)

// EmptyStack shows how to create a Terraform stack.
// The only catch is that this one is empty!
type EmptyStack struct {
	// Embed terra.Stack to implement the terra.Exporter interface
	terra.Stack

	// Add some resources here from our generated code, e.g.
	// VPC aws.vpc `validate:"required"
}

func Example() {
	stack := EmptyStack{}

	// Typically you would use terra.Export() and write to a file. We will
	// write to a buffer for test purposes
	var b bytes.Buffer
	if err := terra.ExportWriter(&stack, &b); err != nil {
		slog.Error("exporting stack", "err", err.Error())
		os.Exit(1)
	}
	fmt.Println(b.String())

	// Output:
	// terraform {
	//
	// }
}