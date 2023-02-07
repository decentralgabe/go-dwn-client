package pkg

import (
	"errors"
	"fmt"
)

type RouteTable struct {
	routes map[string]string
}

func NewRouteTable() *RouteTable {
	return &RouteTable{
		routes: make(map[string]string),
	}
}

func NewRouteTableFromConfig(config map[string]interface{}) *RouteTable {
	routes := make(map[string]string)
	for route, target := range config {
		routes[route] = target.(string)
	}
	return &RouteTable{
		routes: routes,
	}
}

func (rt *RouteTable) AddRoute(route string, target string) error {
	if rt.routes == nil {
		return errors.New("RouteTable is not initialized")
	}
	if gotRoute, ok := rt.routes[route]; ok {
		return fmt.Errorf("route<%s> already exists: %s", route, gotRoute)
	}
	rt.routes[route] = target
	return nil
}

func (rt *RouteTable) GetRoute(route string) (string, error) {
	if rt.routes == nil {
		return "", errors.New("RouteTable is not initialized")
	}
	if gotRoute, ok := rt.routes[route]; ok {
		return gotRoute, nil
	}
	return "", fmt.Errorf("route<%s> not found", route)
}

func (rt *RouteTable) PrintRoutes() {
	if rt.routes == nil {
		return
	}
	fmt.Println("Routes:")
	for route, target := range rt.routes {
		fmt.Printf("%s -> %s\n", route, target)
	}
	fmt.Println()
}

func (rt *RouteTable) RemoveRoute(route string) error {
	if rt.routes == nil {
		return errors.New("RouteTable is not initialized")
	}
	if _, ok := rt.routes[route]; ok {
		delete(rt.routes, route)
		return nil
	}
	return fmt.Errorf("route<%s> not found", route)
}
