/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RpcStatus struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details []map[string]interface{} `json:"details,omitempty"`
}
