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

type NewsInfo struct {
	TeseraId           int64     `json:"teseraId,omitempty"`
	Title              string    `json:"title,omitempty"`
	Title2             string    `json:"title2,omitempty"`
	Alias              string    `json:"alias,omitempty"`
	Content            string    `json:"content,omitempty"`
	CreationDateUtc    time.Time `json:"creationDateUtc,omitempty"`
	PublicationDateUtc time.Time `json:"publicationDateUtc,omitempty"`
	IsDraft            bool      `json:"isDraft,omitempty"`
	PhotoUrl           string    `json:"photoUrl,omitempty"`
	Rating             float64   `json:"rating,omitempty"`
	RatingMy           float64   `json:"ratingMy,omitempty"`
	Source             string    `json:"source,omitempty"`
	SourceUrl          string    `json:"sourceUrl,omitempty"`
	CommentsTotal      int64     `json:"commentsTotal,omitempty"`
	NumVotes           int64     `json:"numVotes,omitempty"`
}
