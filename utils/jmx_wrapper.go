// Package utils contains common utilities used across the integration
package utils

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/jmx"
)

// JMX variable
var (
	// JMXLock is intended to be used to lock around all JMX calls.
	JMXLock sync.Mutex

	// JMXQuery is a wrapper around infra-integrations-sdk/jmx functions to allow
	// easier mocking during tests
	JMXQuery = jmx.Query

	// JMXOpen is a wrapper around infra-integrations-sdk/jmx functions to allow
	// easier mocking during tests
	JMXOpen = jmx.Open

	// JMXClose is a wrapper around infra-integrations-sdk/jmx functions to allow
	// easier mocking during tests
	JMXClose = jmx.Close
)

// SetupJmxTesting sets all JMX wrapper variables to basic shells
func SetupJmxTesting() {
	JMXOpen = func(hostname, port, username, password string) error { return nil }
	JMXClose = func() {}
	JMXQuery = func(query string, timeout int) (map[string]interface{}, error) { return map[string]interface{}{}, nil }
}