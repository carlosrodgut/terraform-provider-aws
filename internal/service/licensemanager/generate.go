// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/tags/main.go -AWSSDKVersion=2 -ServiceTagsSlice -UpdateTags
//go:generate go run ../../generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListLicenseConfigurations,ListLicenseSpecificationsForResource,ListReceivedLicenses,ListDistributedGrants,ListReceivedGrants
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package licensemanager
