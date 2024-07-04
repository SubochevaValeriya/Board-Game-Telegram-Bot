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

type TagInfo struct {
	Id       int64  `json:"id,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	Name     string `json:"name,omitempty"`
}