# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.10.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://dt.garage-trip.cz/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsAssignSkillPoints**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsassignskillpoints) | **Post** /v1/assign-skill-points | Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsBuy**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsbuy) | **Post** /v1/buy | Buy Items identified by the provided ID for the Character bound to the logged user.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsCommands**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollscommands) | **Post** /v1/commands | Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsGame**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsgame) | **Get** /v1/game | Sends all info about the game.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsGameLevel**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsgamelevel) | **Get** /v1/game/{level} | Sends all info about the game level.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsLevels**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollslevels) | **Get** /v1/levels | Sends info about
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsMonstersCommands**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsmonsterscommands) | **Post** /v1/monsters-commands | Control monsters. Admin only.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsMove**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsmove) | **Post** /v1/move | Assign skill point to the attribute for the Character bound to the logged user.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsPickUp**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollspickup) | **Post** /v1/pick-up | Equip the Item from the ground identified by the provided ID for the Character bound to the logged user (unused).
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsPlayers**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsplayers) | **Get** /v1/players | Sends all info about all players.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsRegister**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsregister) | **Post** /v1/register | Register provided User to the Game and create a character.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsRespawn**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsrespawn) | **Post** /v1/respawn | Respawn the Character bound to the logged user.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsSkill**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsskill) | **Post** /v1/skill | Use a skill (provided by an item) by the Character bound to the logged user.
*DungeonsAndTrollsApi* | [**DungeonsAndTrollsYell**](docs/DungeonsAndTrollsApi.md#dungeonsandtrollsyell) | **Post** /v1/yell | The Character bound to the logged user yells a messages (visible for everyone).

## Documentation For Models

 - [DungeonsandtrollsAttributes](docs/DungeonsandtrollsAttributes.md)
 - [DungeonsandtrollsAvailableLevels](docs/DungeonsandtrollsAvailableLevels.md)
 - [DungeonsandtrollsCharacter](docs/DungeonsandtrollsCharacter.md)
 - [DungeonsandtrollsCommandsBatch](docs/DungeonsandtrollsCommandsBatch.md)
 - [DungeonsandtrollsCommandsForMonsters](docs/DungeonsandtrollsCommandsForMonsters.md)
 - [DungeonsandtrollsCoordinates](docs/DungeonsandtrollsCoordinates.md)
 - [DungeonsandtrollsDamageType](docs/DungeonsandtrollsDamageType.md)
 - [DungeonsandtrollsDecoration](docs/DungeonsandtrollsDecoration.md)
 - [DungeonsandtrollsDroppable](docs/DungeonsandtrollsDroppable.md)
 - [DungeonsandtrollsEffect](docs/DungeonsandtrollsEffect.md)
 - [DungeonsandtrollsEvent](docs/DungeonsandtrollsEvent.md)
 - [DungeonsandtrollsEventType](docs/DungeonsandtrollsEventType.md)
 - [DungeonsandtrollsFogOfWarMap](docs/DungeonsandtrollsFogOfWarMap.md)
 - [DungeonsandtrollsGameState](docs/DungeonsandtrollsGameState.md)
 - [DungeonsandtrollsIdentifier](docs/DungeonsandtrollsIdentifier.md)
 - [DungeonsandtrollsIdentifiers](docs/DungeonsandtrollsIdentifiers.md)
 - [DungeonsandtrollsItem](docs/DungeonsandtrollsItem.md)
 - [DungeonsandtrollsItemType](docs/DungeonsandtrollsItemType.md)
 - [DungeonsandtrollsKey](docs/DungeonsandtrollsKey.md)
 - [DungeonsandtrollsLevel](docs/DungeonsandtrollsLevel.md)
 - [DungeonsandtrollsMap](docs/DungeonsandtrollsMap.md)
 - [DungeonsandtrollsMapObjects](docs/DungeonsandtrollsMapObjects.md)
 - [DungeonsandtrollsMessage](docs/DungeonsandtrollsMessage.md)
 - [DungeonsandtrollsMonster](docs/DungeonsandtrollsMonster.md)
 - [DungeonsandtrollsPlayerSpecificMap](docs/DungeonsandtrollsPlayerSpecificMap.md)
 - [DungeonsandtrollsPlayersInfo](docs/DungeonsandtrollsPlayersInfo.md)
 - [DungeonsandtrollsPosition](docs/DungeonsandtrollsPosition.md)
 - [DungeonsandtrollsRegistration](docs/DungeonsandtrollsRegistration.md)
 - [DungeonsandtrollsSimpleItem](docs/DungeonsandtrollsSimpleItem.md)
 - [DungeonsandtrollsSkill](docs/DungeonsandtrollsSkill.md)
 - [DungeonsandtrollsSkillAttributes](docs/DungeonsandtrollsSkillAttributes.md)
 - [DungeonsandtrollsSkillEffect](docs/DungeonsandtrollsSkillEffect.md)
 - [DungeonsandtrollsSkillGenericFlags](docs/DungeonsandtrollsSkillGenericFlags.md)
 - [DungeonsandtrollsSkillSpecificFlags](docs/DungeonsandtrollsSkillSpecificFlags.md)
 - [DungeonsandtrollsSkillUse](docs/DungeonsandtrollsSkillUse.md)
 - [DungeonsandtrollsStun](docs/DungeonsandtrollsStun.md)
 - [DungeonsandtrollsUser](docs/DungeonsandtrollsUser.md)
 - [DungeonsandtrollsWaypoint](docs/DungeonsandtrollsWaypoint.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [SkillTarget](docs/SkillTarget.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


