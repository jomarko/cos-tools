/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public
import (
	"time"
)
// ObjectMeta struct for ObjectMeta
type ObjectMeta struct {
	Owner string `json:"owner,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
}
