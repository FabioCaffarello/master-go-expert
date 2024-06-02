package gosd

import (
	"fmt"
	"os"
)

type ServiceDiscovery struct {
	UsesHostAddr bool
	services     map[string]string
}

func NewServiceDiscovery(
	usesHostAddr bool,
) *ServiceDiscovery {
	return &ServiceDiscovery{
		UsesHostAddr: usesHostAddr,
		services:     make(map[string]string),
	}
}

// RegisterService registers a service with a given hostname and port
func (sd *ServiceDiscovery) RegisterService(
	name string,
	port string,
) {
	var hostName string
	if !sd.UsesHostAddr {
		hostName = "localhost"
	} else {
		hostName = name
	}
	sd.services[name] = fmt.Sprintf("http://%s:%s", hostName, port)
}

// GetBaseURL returns the base URL for a registered service
func (sd *ServiceDiscovery) GetBaseURL(
	serviceName string,
) (string, error) {
	if url, exists := sd.services[serviceName]; exists {
		return url, nil
	}
	return "", fmt.Errorf("service %s not found", serviceName)
}

// GetBaseURLFromEnv constructs the base URL for a service from environment variables
func (sd *ServiceDiscovery) GetBaseURLFromEnv(
	serviceName string,
) (string, error) {
	hostName := os.Getenv(serviceName + "_HOST")
	if !sd.UsesHostAddr {
		hostName = "localhost"
	}
	port := os.Getenv(serviceName + "_PORT")
	if hostName == "" || port == "" {
		return "", fmt.Errorf("environment variables for %s not set", serviceName)
	}
	return fmt.Sprintf("http://%s:%s", hostName, port), nil
}
