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

type FeedItem struct {
	FirstImg       string    `json:"firstImg,omitempty"`
	Title          string    `json:"title,omitempty"`
	SecondImg      string    `json:"secondImg,omitempty"`
	CreateDateTime time.Time `json:"createDateTime,omitempty"`
	SubscribeType  string    `json:"subscribeType,omitempty"`
	Alias          string    `json:"alias,omitempty"`
}
