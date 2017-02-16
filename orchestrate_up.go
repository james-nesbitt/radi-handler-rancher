package rancher

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	api_orchestrate "github.com/wunderkraut/radi-api/operation/orchestrate"
	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Orchestrate Up operation for rancher
 */

type RancherOrchestrateUpOperation struct {
	api_orchestrate.BaseOrchestrationUpOperation
	RancherBaseClientOperation
}

// Alter the ID of the parent operation
func (up *RancherOrchestrateUpOperation) Id() string {
	return "rancher." + up.BaseOrchestrationUpOperation.Id()
}

// Run a validation check on the Operation
func (up *RancherOrchestrateUpOperation) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// What settings/values does the Operation provide to an implemenentor
func (up *RancherOrchestrateUpOperation) Properties() api_property.Properties {
	return api_property.New_SimplePropertiesEmpty().Properties()
}

// Execute the operation
func (up *RancherOrchestrateUpOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	log.WithFields(log.Fields{"clientsettings": up.RancherClientSettings(), "envsettings": up.RancherEnvironmentSettings()}).Info("SETTINGS")

	res.AddError(errors.New("RANCHER UP OPERATION NOT YET WRITTEN"))
	res.MarkFailed()

	res.MarkFinished()
	return res.Result()
}
