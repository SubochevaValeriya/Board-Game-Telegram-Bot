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

type ReportInfoResponse struct {
	Report  *ReportInfo    `json:"report,omitempty"`
	Prefill *ReportPrefill `json:"prefill,omitempty"`
	Config  *ReportConfig  `json:"config,omitempty"`
}
