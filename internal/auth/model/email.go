/*
 * Otto user service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// Email struct for Email
type Email struct {
	User           User           `json:"user"`
	Email          string         `json:"email"`
	VerifiedStatus VerifiedStatus `json:"verifiedStatus"`
}
