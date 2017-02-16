package rancher

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Orchestration Handler using Rancher
 */

// Rancher Orchestration Handler
type RancherOrchestrateHandler struct {
	RancherBaseClientHandler
}

// Return a string identifier for the Handler (not functionally needed yet)
func (orchestrate *RancherOrchestrateHandler) Id() string {
	return "rancher.orchestrate"
}

// Initialize and activate the Handler
func (orchestrate *RancherOrchestrateHandler) Operations() api_operation.Operations {
	base := New_RancherBaseClientOperation(orchestrate.ConfigSource())

	ops := api_operation.New_SimpleOperations()

	ops.Add(api_operation.Operation(&RancherOrchestrateUpOperation{RancherBaseClientOperation: *base}))
	ops.Add(api_operation.Operation(&RancherOrchestrateDownOperation{RancherBaseClientOperation: *base}))

	return ops.Operations()
}
