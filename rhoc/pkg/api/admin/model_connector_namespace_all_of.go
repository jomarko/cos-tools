/*
 * Connector Service Fleet Manager Admin APIs
 *
 * Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

// ConnectorNamespaceAllOf struct for ConnectorNamespaceAllOf
type ConnectorNamespaceAllOf struct {
	Name       string                   `json:"name"`
	ClusterId  string                   `json:"cluster_id"`
	Expiration string                   `json:"expiration,omitempty"`
	Tenant     ConnectorNamespaceTenant `json:"tenant"`
	Status     ConnectorNamespaceStatus `json:"status"`
}
