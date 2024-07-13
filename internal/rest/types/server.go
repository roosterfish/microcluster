package types

import (
	"github.com/canonical/microcluster/internal/extensions"
	"github.com/canonical/microcluster/rest/types"
)

// Server represents server status information.
type Server struct {
	Name       string                `json:"name"    yaml:"name"`
	Address    types.AddrPort        `json:"address" yaml:"address"`
	Version    string                `json:"version" yaml:"version"`
	Ready      bool                  `json:"ready"   yaml:"ready"`
	Extensions extensions.Extensions `json:"extensions" yaml:"extensions"`
}

const (
	// PublicEndpoint - Internally managed APIs available without authentication.
	PublicEndpoint types.EndpointPrefix = "cluster/1.0"

	// InternalEndpoint - all endpoints restricted to trusted servers.
	InternalEndpoint types.EndpointPrefix = "cluster/internal"

	// ControlEndpoint - all endpoints available on the local unix socket.
	ControlEndpoint types.EndpointPrefix = "cluster/control"
)
