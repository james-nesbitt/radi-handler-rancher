package rancher

import (
	rancher_client "github.com/rancher/go-rancher/client"

	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Some base structs to share upcloud functionality
 */

// Shared base handler
type RancherBaseClientHandler struct {
	configSource RancherConfigSource
}

// Constructor for RancherBaseClientHandler
func New_RancherBaseClientHandler(configSource RancherConfigSource) *RancherBaseClientHandler {
	return &RancherBaseClientHandler{
		configSource: configSource,
	}
}

// Retrieve the base settings
func (base *RancherBaseClientHandler) ConfigSource() RancherConfigSource {
	return base.configSource
}

// Validate the Handler
func (base *RancherBaseClientHandler) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// Share base operation
type RancherBaseClientOperation struct {
	configSource RancherConfigSource
}

// constructor for configSource RancherConfigSource
func New_RancherBaseClientOperation(configSource RancherConfigSource) *RancherBaseClientOperation {
	return &RancherBaseClientOperation{
		configSource: configSource,
	}
}

// Retrieve a rancher client
func (base *RancherBaseClientOperation) RancherClient() *rancher_client.RancherClient {
	return base.configSource.RancherClient()
}

// Retrieve the base settings
func (base *RancherBaseClientOperation) RancherClientSettings() RancherClientSettings {
	return base.configSource.RancherClientSettings()
}

// Retrieve the base settings
func (base *RancherBaseClientOperation) RancherEnvironmentSettings() RancherEnvironmentSettings {
	return base.configSource.RancherEnvironmentSettings()
}
