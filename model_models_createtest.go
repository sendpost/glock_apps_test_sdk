/*
 * glockapps
 *
 * glock apps api for inboxing
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsCreatetest struct {
	TestID string `json:"TestID,omitempty"`
	InsertHeader string `json:"InsertHeader,omitempty"`
	InsertInBody string `json:"InsertInBody,omitempty"`
	SeedList []string `json:"SeedList,omitempty"`
}
