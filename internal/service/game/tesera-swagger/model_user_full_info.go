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

import (
	"time"
)

type UserFullInfo struct {
	TeseraId            int64        `json:"teseraId,omitempty"`
	Login               string       `json:"login,omitempty"`
	Name                string       `json:"name,omitempty"`
	Comment             string       `json:"comment,omitempty"`
	AvatarUrl           string       `json:"avatarUrl,omitempty"`
	Gender              string       `json:"gender,omitempty"`
	Rating              int64        `json:"rating,omitempty"`
	Experience          int64        `json:"experience,omitempty"`
	CreationDateUtc     time.Time    `json:"creationDateUtc,omitempty"`
	ModificationDateUtc time.Time    `json:"modificationDateUtc,omitempty"`
	CollectionCount     int64        `json:"collectionCount,omitempty"`
	GamesAdded          int64        `json:"gamesAdded,omitempty"`
	NewsAdded           int64        `json:"newsAdded,omitempty"`
	ArticlesAdded       int64        `json:"articlesAdded,omitempty"`
	JournalsAdded       int64        `json:"journalsAdded,omitempty"`
	ThoughtsAdded       int64        `json:"thoughtsAdded,omitempty"`
	Country             *CountryInfo `json:"country,omitempty"`
	City                *CityInfo    `json:"city,omitempty"`
}
