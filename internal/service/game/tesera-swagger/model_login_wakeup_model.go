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

type LoginWakeupModel struct {
	// User's login
	Login string `json:"login,omitempty"`
	// Secret user hash
	Hash      string `json:"hash,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
}
