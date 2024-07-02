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

type UpdateMapSettingsModel struct {
	LocationMessage string `json:"locationMessage,omitempty"`
	ViewEmail       bool   `json:"viewEmail,omitempty"`
	InGuest         bool   `json:"inGuest,omitempty"`
	InHome          bool   `json:"inHome,omitempty"`
	InClub          bool   `json:"inClub,omitempty"`
	IamPlayer       bool   `json:"iamPlayer,omitempty"`
	IamClub         bool   `json:"iamClub,omitempty"`
	IamShop         bool   `json:"iamShop,omitempty"`
	IamEvent        bool   `json:"iamEvent,omitempty"`
}
