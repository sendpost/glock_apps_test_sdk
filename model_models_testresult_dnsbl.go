/*
 * glockapps
 *
 * glock apps api for inboxing
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsTestresultDnsbl struct {
	Finished bool `json:"finished,omitempty"`
	IpListed bool `json:"ipListed,omitempty"`
	Dnsbl []ModelsTestresultDnsblDnsbl `json:"dnsbl,omitempty"`
}
