/*
 * Connector Service Fleet Manager Admin APIs
 *
 * Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

// ConnectorState the model 'ConnectorState'
type ConnectorState string

// List of ConnectorState
const (
	CONNECTORSTATE_ASSIGNING      ConnectorState = "assigning"
	CONNECTORSTATE_ASSIGNED       ConnectorState = "assigned"
	CONNECTORSTATE_UPDATING       ConnectorState = "updating"
	CONNECTORSTATE_READY          ConnectorState = "ready"
	CONNECTORSTATE_STOPPED        ConnectorState = "stopped"
	CONNECTORSTATE_FAILED         ConnectorState = "failed"
	CONNECTORSTATE_DELETING       ConnectorState = "deleting"
	CONNECTORSTATE_DELETED        ConnectorState = "deleted"
	CONNECTORSTATE_PROVISIONING   ConnectorState = "provisioning"
	CONNECTORSTATE_DEPROVISIONING ConnectorState = "deprovisioning"
)
