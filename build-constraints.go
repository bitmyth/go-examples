// https://riptutorial.com/go/example/8607/separate-integration-tests
// Build constraints are commonly used to separate normal unit tests from
// integration tests that require external resources, like a database or network access.
// To do this, add a custom build constraint to the top of the test file:

// +build integration
 
package main
 
import (
    "testing"
)
 
func TestThatRequiresNetworkAccess(t *testing.T) {
    t.Fatal("It failed!")
}

// The test file will not compile into the build executable unless the following invocation of go test is used:

// go test -tags "integration"
