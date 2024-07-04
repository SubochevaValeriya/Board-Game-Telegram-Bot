/*
 * Tesera API
 *
 * API for Tesera
 *
 * API version: v1
 * Contact:
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UnsubscribeMailModel struct {
	// Email
	Email string `json:"email,omitempty"`
	// Key hash
	Key string `json:"key,omitempty"`
}