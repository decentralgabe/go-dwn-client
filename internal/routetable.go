package internal

import (
	"errors"
	"fmt"
)

type RouteTable struct {
	routes map[string]string
}

func NewRouteTableFromConfig(config interface{}) *RouteTable {
	routes := make(map[string]string)
	if config != nil {
		for route, target := range config.(map[string]interface{}) {
			routes[route] = target.(string)
		}
		return &RouteTable{
			routes: routes,
		}
	}
	return &RouteTable{
		routes: routes,
	}
}

func (rt *RouteTable) AddRoute(did string, target string) error {
	if rt.routes == nil {
		return errors.New("RouteTable is not initialized")
	}
	if gotRoute, ok := rt.routes[did]; ok {
		return fmt.Errorf("did<%s> already exists: %s", did, gotRoute)
	}
	rt.routes[did] = target
	if err := SaveRouteTable(rt); err != nil {
		return err
	}
	rt.PrintRoutes()
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
	fmt.Println("All routes:")
	for route, target := range rt.routes {
		fmt.Printf("%s -> %s\n", route, target)
	}
	fmt.Println()
}

func (rt *RouteTable) PrintRoute(did string) {
	if rt.routes == nil {
		return
	}
	if route, ok := rt.routes[did]; ok {
		fmt.Printf("Route for did<%s>: %s\n", did, route)
		return
	}
	fmt.Printf("No route found for did<%s>, try adding one with 'rdr route add <did> <target>'\n", did)
}

func (rt *RouteTable) RemoveRoute(did string) error {
	if rt.routes == nil {
		return errors.New("RouteTable is not initialized")
	}
	if route, ok := rt.routes[did]; ok {
		delete(rt.routes, did)
		fmt.Printf("Removed route<%s> for did<%s>\n", route, did)
		if err := SaveRouteTable(rt); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("did<%s> not found", did)
}
