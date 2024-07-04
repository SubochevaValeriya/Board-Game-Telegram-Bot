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

type PostMessageModel struct {
	// User's message
	Comment string `json:"comment,omitempty"`
	// Conversation object
	ObjectType string `json:"objectType,omitempty"`
	// Conversation object id
	ObjectId int64 `json:"objectId,omitempty"`
}