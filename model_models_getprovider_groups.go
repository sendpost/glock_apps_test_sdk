/*
 * glockapps
 *
 * glock apps api for inboxing
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsGetproviderGroups struct {
	GroupID int32 `json:"GroupID,omitempty"`
	GroupName string `json:"GroupName,omitempty"`
	ISPs []ModelsGetproviderIsPs `json:"ISPs,omitempty"`
}