#!/bin/sh
go test support.go \
  aligncheck_test.go \
  deadcode_test.go \
  errcheck_test.go \
  golint_test.go \
  gosimple_test.go \
  interfacer_test.go \
  staticcheck_test.go \
  structcheck_test.go \
  unconvert_test.go \
  unused_test.go \
  varcheck_test.go \
