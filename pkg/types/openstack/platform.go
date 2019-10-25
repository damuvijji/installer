package openstack

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// Region specifies the OpenStack region where the cluster will be created.
	// Deprecated: this value is not used by the installer.
	Region string `json:"region"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on OpenStack for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// Cloud
	// Name of OpenStack cloud to use from clouds.yaml
	Cloud string `json:"cloud"`

	// ExternalNetwork
	// The OpenStack external network name to be used for installation.
	ExternalNetwork string `json:"externalNetwork"`

	// FlavorName
	// The OpenStack compute flavor to use for servers.
	FlavorName string `json:"computeFlavor"`

	// LbFloatingIP
	// Existing Floating IP to associate with the OpenStack load balancer.
	LbFloatingIP string `json:"lbFloatingIP"`

	// ExternalDNS holds the IP addresses of dns servers that will
	// be added to the dns resolution of all instances in the cluster.
	// +optional
	ExternalDNS []string `json:"externalDNS"`

	// TrunkSupport
	// Whether OpenStack ports can be trunked
	TrunkSupport string `json:"trunkSupport"`

	// OctaviaSupport
	// Whether OpenStack has Octavia support
	OctaviaSupport string `json:"octaviaSupport"`
}
