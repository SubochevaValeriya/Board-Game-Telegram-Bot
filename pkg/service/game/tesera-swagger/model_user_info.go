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

type UserInfo struct {
	TeseraId      int64  `json:"teseraId,omitempty"`
	Login         string `json:"login,omitempty"`
	Name          string `json:"name,omitempty"`
	AvatarUrl     string `json:"avatarUrl,omitempty"`
	CommentsTotal int64  `json:"commentsTotal,omitempty"`
	Rating        int64  `json:"rating,omitempty"`
	Experience    int64  `json:"experience,omitempty"`
	RightGroup    string `json:"rightGroup,omitempty"`
	TeseraUrl     string `json:"teseraUrl,omitempty"`
}