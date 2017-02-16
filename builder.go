package rancher

import (
	log "github.com/Sirupsen/logrus"

	api_api "github.com/wunderkraut/radi-api/api"
	api_builder "github.com/wunderkraut/radi-api/builder"
	api_handler "github.com/wunderkraut/radi-api/handler"
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_config "github.com/wunderkraut/radi-api/operation/config"
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Handler builder and settings to create rancher handlers
 */

type RancherBuilder struct {
	settings RancherSettings

	parent   api_api.API
	handlers api_handler.Handlers
}

// Return a string identifier for the Handler (not functionally needed yet)
func (builder *RancherBuilder) Id() string {
	return "rancher"
}

// Set a API for this Handler
func (builder *RancherBuilder) SetAPI(parent api_api.API) {
	builder.parent = parent
}

// Initialize and activate the Handler
func (builder *RancherBuilder) Activate(implementations api_builder.Implementations, settingsProvider api_builder.SettingsProvider) api_result.Result {
	if &builder.handlers == nil {
		builder.handlers = api_handler.Handlers{}
	}

	// create a base handler
	base := builder.base_RancherBaseClientHandler(settingsProvider)

	// overall result
	res := api_result.New_StandardResult()

	for _, implementation := range implementations.Order() {
		var handler api_handler.Handler

		switch implementation {
		case "command":
			handler = api_handler.Handler(&RancherCommandHandler{RancherBaseClientHandler: *base})
		case "monitor":
			handler = api_handler.Handler(&RancherMonitorHandler{RancherBaseClientHandler: *base})
		case "orchestrate":
			handler = api_handler.Handler(&RancherOrchestrateHandler{RancherBaseClientHandler: *base})
		}

		if handler == nil {
			log.WithFields(log.Fields{"implementation": implementation}).Warn("Rancher:Builder: Unknown implementation in Rancher builder")
		} else {
			// Initialize any new handler
			initRes := handler.Validate()
			<-initRes.Finished()

			if initRes.Success() {
				builder.handlers.Add(api_handler.Handler(handler))

			} else {
				log := log.WithFields(log.Fields{"implementation": implementation})
				for _, err := range initRes.Errors() {
					log.WithError(err)
				}
				log.Warn("Rancher:Builder: Unknown implementation in Rancher builder")
			}

			res.Merge(initRes)
		}

	}

	return res
}

// Validate the builder after Activation is complete
func (builder *RancherBuilder) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// Return a list of Operations from the Handler
func (builder *RancherBuilder) Operations() api_operation.Operations {
	return builder.handlers.Operations()
}

// Return a shared BaseUpcloudServiceOperation for any operation that needs it
func (builder *RancherBuilder) base_RancherBaseClientHandler(settingsProvider api_builder.SettingsProvider) *RancherBaseClientHandler {
	// Builder a configwrapper, which will be used to build upcloud service structs
	ops := builder.parent.Operations()
	configWrapper := api_config.New_SimpleConfigWrapper(ops)

	// try to convert the settings to RancherSettings
	builder.settings = RancherSettings{}
	settingsProvider.AssignSettings(&builder.settings)

	// make a new config handler from the configWrapper (for now just assume yaml)
	ymlConfigSource := New_RancherConfigSourceYaml(configWrapper)

	// Make a YAML based config wrapper
	return New_RancherBaseClientHandler(ymlConfigSource)
}
