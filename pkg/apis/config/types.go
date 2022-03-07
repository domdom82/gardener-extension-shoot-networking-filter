// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Configuration contains information about the policy filter configuration.
type Configuration struct {
	metav1.TypeMeta

	// EgressFilter contains the configuration for the egress filter
	EgressFilter *EgressFilter

	// HealthCheckConfig is the config for the health check controller.
	HealthCheckConfig *healthcheckconfig.HealthCheckConfig
}

// EgressFilter contains the configuration for the egress filter.
type EgressFilter struct {
	// BlackholingEnabled is a flag to set blackholing or firewall approach.
	BlackholingEnabled bool

	// FilterListProviderType specifies how the filter list is retrieved.
	// Supported types are `static` and `download`.
	FilterListProviderType FilterListProviderType

	// StaticFilterList contains the static filter list.
	// Only used for provider type `static`.
	StaticFilterList []Filter

	// DownloaderConfig contains the configuration for the filter list downloader.
	// Only used for provider type `download`.
	DownloaderConfig *DownloaderConfig
}

// FilterListProviderType
type FilterListProviderType string

const (
	// FilterListProviderTypeStatic is the provider type for static filter list
	FilterListProviderTypeStatic FilterListProviderType = "static"
	// FilterListProviderTypeDownload is the provider type for downloading the filter list from an URL
	FilterListProviderTypeDownload FilterListProviderType = "download"
)

// Policy is the access policy
type Policy string

const (
	// PolicyAllowAccess is the `ALLOW_ACCESS` policy
	PolicyAllowAccess Policy = "ALLOW_ACCESS"
	// PolicyBlockAccess is the `BLOCK_ACCESS` policy
	PolicyBlockAccess Policy = "BLOCK_ACCESS"
)

// Filter specifies a network-CIDR policy pair.
type Filter struct {
	// Network is the network CIDR of the filter.
	Network string
	// Policy is the access policy (`BLOCK_ACCESS` or `ALLOW_ACCESS`).
	Policy Policy
}

// DownloaderConfig contains the configuration for the filter list downloader.
type DownloaderConfig struct {
	// Endpoint is the endpoint URL for downloading the filter list.
	Endpoint string
	// Authorization is static authorization header.
	Authorization *string
	// RefreshPeriod is interval for refreshing the filter list.
	// If unset, the filter list is only fetched on startup.
	RefreshPeriod *metav1.Duration
}
