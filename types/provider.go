package types

import (
	"github.com/hashicorp/terraform/config/configschema"
	"github.com/hashicorp/terraform/tfdiags"
	"github.com/zclconf/go-cty/cty"
)

// Provider represents the set of methods required for a complete resource
// provider plugin.
type Provider interface {
	// GetSchema returns the complete schema for the provider.
	GetSchema() GetSchemaResponse

	// ValidateProviderConfig allows the provider to validate the provider
	// configuration values.
	ValidateProviderConfig(ValidateProviderConfigRequest) ValidateProviderConfigResponse

	// ValidateResourceTypeConfig allows the provider to validate the resource
	// configuration values.
	ValidateResourceTypeConfig(ValidateResourceTypeConfigRequest) ValidateResourceTypeConfigResponse

	// ValidateDataSource allows the provider to validate the data source
	// configuration values.
	ValidateDataSourceConfig(ValidateDataSourceConfigRequest) ValidateDataSourceConfigResponse

	// UpgradeResourceState is called when the state loader encounters an
	// instance state whose schema version is less than the one reported by the
	// currently-used version of the corresponding provider, and the upgraded
	// result is used for any further processing.
	UpgradeResourceState(UpgradeResourceStateRequest) UpgradeResourceStateResponse

	// Configure configures and initialized the provider.
	Configure(ConfigureRequest) ConfigureResponse

	// Stop is called when the provider should halt any in-flight actions.
	//
	// Stop should not block waiting for in-flight actions to complete. It
	// should take any action it wants and return immediately acknowledging it
	// has received the stop request. Terraform will not make any further API
	// calls to the provider after Stop is called.
	//
	// The error returned, if non-nil, is assumed to mean that signaling the
	// stop somehow failed and that the user should expect potentially waiting
	// a longer period of time.
	Stop() error

	// ReadResource refreshes a resource and returns its current state.
	ReadResource(ReadResourceRequest) ReadResourceResponse

	// PlanResourceChange takes the current state and proposed state of a
	// resource, and returns the planned final state.
	PlanResourceChange(PlanResourceChangeRequest) PlanResourceChangeResponse

	// ApplyResourceChange takes the planned state for a resource, which may
	// yet contain unknown computed values, and applies the changes returning
	// the final state.
	ApplyResourceChange(ApplyResourceChangeRequest) ApplyResourceChangeResponse

	// ImportResourceState requests that the given resource be imported.
	ImportResourceState(ImportResourceStateRequest) ImportResourceStateResponse

	// ReadDataSource returns the data source's current state.
	ReadDataSource(ReadDataSourceRequest) ReadDataSourceResponse
}

type GetSchemaResponse struct {
	Provider      *configschema.Block
	ResourceTypes map[string]*configschema.Block
	DataSources   map[string]*configschema.Block
	Diagnostics   tfdiags.Diagnostics
}

type ValidateProviderConfigRequest struct {
	Config cty.Value
}

type ValidateProviderConfigResponse struct {
	Diagnostics tfdiags.Diagnostics
}

type ValidateResourceTypeConfigRequest struct {
	Name   string
	Config cty.Value
}

type ValidateResourceTypeConfigResponse struct {
	Diagnostics tfdiags.Diagnostics
}

type ValidateDataSourceConfigRequest struct {
	Name   string
	Config cty.Value
}

type ValidateDataSourceConfigResponse struct {
	Diagnostics tfdiags.Diagnostics
}

type UpgradeResourceStateRequest struct {
	Name    string
	Version int
	state   cty.Value
}

type UpgradeResourceStateResponse struct {
	State       cty.Value
	Diagnostics tfdiags.Diagnostics
}

type ConfigureRequest struct {
	Config cty.Value
}

type ConfigureResponse struct {
	Diagnostics tfdiags.Diagnostics
}

type ReadResourceRequest struct {
	Name       string
	PriorState cty.Value
}

type ReadResourceResponse struct {
	NewState    cty.Value
	Diagnostics tfdiags.Diagnostics
}

type PlanResourceChangeRequest struct {
	Name         string
	PriorState   cty.Value
	PriorPrivate []byte
}

type PlanResourceChangeResponse struct {
	PlannedState   cty.Value
	PlannedPrivate []byte
	Diagnostics    tfdiags.Diagnostics
}

type ApplyResourceChangeRequest struct {
	Name           string
	PriorState     cty.Value
	PlannedState   cty.Value
	PlannedPrivate []byte
}

type ApplyResourceChangeResponse struct {
	NewState    cty.Value
	Private     []byte
	Diagnostics tfdiags.Diagnostics
}

type ImportResourceStateRequest struct {
	Name string
	ID   string
}

type ImportResourceStateResponse struct {
	State       []cty.Value
	Diagnostics tfdiags.Diagnostics
}

type ReadDataSourceRequest struct {
	Name string
}

type ReadDataSourceResponse struct {
	State       cty.Value
	Diagnostics tfdiags.Diagnostics
}
