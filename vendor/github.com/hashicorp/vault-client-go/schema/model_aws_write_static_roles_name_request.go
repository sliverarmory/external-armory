// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsWriteStaticRolesNameRequest struct for AwsWriteStaticRolesNameRequest
type AwsWriteStaticRolesNameRequest struct {
	// Period by which to rotate the backing credential of the adopted user. This can be a Go duration (e.g, '1m', 24h'), or an integer number of seconds.
	RotationPeriod string `json:"rotation_period,omitempty"`

	// The IAM user to adopt as a static role.
	Username string `json:"username,omitempty"`
}