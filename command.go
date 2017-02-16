package rancher

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Command Handler using Rancher
 */

// Rancher Command Handler
type RancherCommandHandler struct {
	RancherBaseClientHandler
}

// Rturn a string identifier for the Handler (not functionally needed yet)
func (command *RancherCommandHandler) Id() string {
	return "rancher.command"
}

// Validate the Handler
func (command *RancherCommandHandler) Operations() api_operation.Operations {
	ops := api_operation.New_SimpleOperations()

	// ops.Add(api_operation.Operation(&RancherCommandUpOperation{BaseRancherServiceOperation: *baseOperation}))
	// ops.Add(api_operation.Operation(&RancherCommandStopOperation{BaseRancherServiceOperation: *baseOperation}))
	// ops.Add(api_operation.Operation(&RancherCommandDownOperation{BaseRancherServiceOperation: *baseOperation}))

	return ops.Operations()
}
