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

type UserSimilarInfo struct {
	TypeWeight        string         `json:"typeWeight,omitempty"`
	Weight            int64          `json:"weight,omitempty"`
	CommonGames       int64          `json:"commonGames,omitempty"`
	RateOnePoint      int64          `json:"rateOnePoint,omitempty"`
	RateTwoPoint      int64          `json:"rateTwoPoint,omitempty"`
	SimilarUser       *UserInfoShort `json:"similarUser,omitempty"`
	GamesInCollection int32          `json:"gamesInCollection,omitempty"`
}
