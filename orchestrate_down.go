package rancher

import (
	"errors"

	api_orchestrate "github.com/wunderkraut/radi-api/operation/orchestrate"
	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Orchestrate Down operation for rancher
 */

type RancherOrchestrateDownOperation struct {
	api_orchestrate.BaseOrchestrationDownOperation
	RancherBaseClientOperation
}

// Alter the ID of the parent operation
func (down *RancherOrchestrateDownOperation) Id() string {
	return "rancher." + down.BaseOrchestrationDownOperation.Id()
}

// Run a validation check on the Operation
func (down *RancherOrchestrateDownOperation) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// What settings/values does the Operation provide to an implemenentor
func (down *RancherOrchestrateDownOperation) Properties() api_property.Properties {
	return api_property.New_SimplePropertiesEmpty().Properties()
}

// Execute the operation
func (down *RancherOrchestrateDownOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	res.AddError(errors.New("RANCHER DOWN OPERATION NOT YET WRITTEN"))
	res.MarkFailed()

	res.MarkFinished()
	return res.Result()
}
