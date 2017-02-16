package rancher

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Monitor Handler using Rancher
 */

// Rancher Monitor Handler
type RancherMonitorHandler struct {
	RancherBaseClientHandler
}

// Rturn a string identifier for the Handler (not functionally needed yet)
func (monitor *RancherMonitorHandler) Id() string {
	return "rancher.monitor"
}

// Initialize and activate the Handler
func (monitor *RancherMonitorHandler) Operations() api_operation.Operations {
	base := New_RancherBaseClientOperation(monitor.ConfigSource())

	ops := api_operation.New_SimpleOperations()

	ops.Add(api_operation.Operation(&RancherMonitorListEnvironmentsOperation{RancherBaseClientOperation: *base}))

	return ops.Operations()
}
