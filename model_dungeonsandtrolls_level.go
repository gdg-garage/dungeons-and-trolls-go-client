/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsLevel struct {
	Level int32 `json:"level,omitempty"`
	Width int32 `json:"width,omitempty"`
	Height int32 `json:"height,omitempty"`
	Free [][]interface{} `json:"free,omitempty"`
	Objects []DungeonsandtrollsMapObjects `json:"objects,omitempty"`
	Ascii []string `json:"ascii,omitempty"`
}
