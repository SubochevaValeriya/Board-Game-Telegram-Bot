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

type ReportCommentInfo struct {
	Id                  int64          `json:"id,omitempty"`
	ReportId            int64          `json:"reportId,omitempty"`
	AuthorId            int64          `json:"authorId,omitempty"`
	Content             string         `json:"content,omitempty"`
	Level               int64          `json:"level,omitempty"`
	LeftKey             int64          `json:"leftKey,omitempty"`
	RightKey            int64          `json:"rightKey,omitempty"`
	CreationDateUtc     time.Time      `json:"creationDateUtc,omitempty"`
	ModificationDateUtc time.Time      `json:"modificationDateUtc,omitempty"`
	Rating              int64          `json:"rating,omitempty"`
	Status              int64          `json:"status,omitempty"`
	Author              *UserInfoShort `json:"author,omitempty"`
}