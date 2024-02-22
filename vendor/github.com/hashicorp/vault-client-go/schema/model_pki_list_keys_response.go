// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListKeysResponse struct for PkiListKeysResponse
type PkiListKeysResponse struct {
	// Key info with issuer name
	KeyInfo map[string]interface{} `json:"key_info,omitempty"`

	// A list of keys
	Keys []string `json:"keys,omitempty"`
}
