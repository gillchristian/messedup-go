package services

import "fmt"

var services Services

// Find a service by name
func Find(name string) (Service, error) {
	for _, s := range services {
		if s.name == name {
			return s, nil
		}
	}
	return Service{}, fmt.Errorf("Could not find a Service with name <%s>", name)
}

// Add a service to the collection
func Add(s Service) Service {
	services.append(s)
	return s
}

// Remove a service from the collection
func Remove(name string) error {
	for i, s := range services {
		if s.name == name {
			services = append(services[:i], services[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find a Service with name <%s> to remove", name)
}
