package main

import (
	"github.com/hashicorp/hcl/v2/hcl"
)

type LogBeginCallback func(testName string, testFile *TestFile)
type LogProblemsCallback func(testName string, testFile *TestFile, diags hcl.Diagnostics)