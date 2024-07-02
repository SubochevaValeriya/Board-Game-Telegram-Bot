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

type UserFullInfoResponse struct {
	User              *UserFullInfo        `json:"user,omitempty"`
	Photos            []PhotoInfo          `json:"photos,omitempty"`
	GamesInCollection int32                `json:"gamesInCollection,omitempty"`
	Collection        []CollectionGameInfo `json:"collection,omitempty"`
	Sales             []AdvertInfo         `json:"sales,omitempty"`
	Purchases         []AdvertInfo         `json:"purchases,omitempty"`
}
