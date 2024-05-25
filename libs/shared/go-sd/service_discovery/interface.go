package gosd


type ServiceDiscoveryInterface interface {
	RegisterService(name, port string)
	GetBaseURL(serviceName string) (string, error)
	GetBaseURLFromEnv(serviceName string) (string, error)
}
