/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsGameState struct {
	Map_ *DungeonsandtrollsMap `json:"map,omitempty"`
	ShopItems []DungeonsandtrollsItem `json:"shopItems,omitempty"`
	Character *DungeonsandtrollsCharacter `json:"character,omitempty"`
	CurrentPosition *DungeonsandtrollsPosition `json:"currentPosition,omitempty"`
	CurrentLevel int32 `json:"currentLevel,omitempty"`
	Tick string `json:"tick,omitempty"`
	// List of events which occurred in the previous tick. Useful for visualising effects, debugging and communication.
	Events []DungeonsandtrollsEvent `json:"events,omitempty"`
	Score float32 `json:"score,omitempty"`
}
