package main

import (
	"bazel-golang-wasm-protoc/protos/api"
	"fmt"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	manager := &Manager{
		searchBar:     &SearchBar{},
		instanceTable: &InstanceTable{},
		pagination:    &Pagination{pageSize: 10, currentPage:1},
	}

	manager.searchBar.SetManager(manager)
	manager.instanceTable.SetManager(manager)
	manager.pagination.SetManager(manager)

	// init
	// manager.pagination.Init()

	app.Route("/", manager)
	app.Run()
}

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo
	searchBar     *SearchBar
	instanceTable *InstanceTable
	pagination    *Pagination
}

func (h *Manager) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				h.searchBar,
			),
		),
		app.Div().Class("container-fluid").Body(
			h.instanceTable,
		),
		app.Div().Body(
			h.pagination,
		),
	)
}

func (h *Manager) Search(q string, pageIndex int, pageSize int) []*api.Instance {
	instances, err := api.CallApiSearch(api.SearchRequest{
		Query:     q,
		PageIndex: int32(pageIndex),
		PageSize:  int32(pageSize),
	})

	if err != nil {
		fmt.Println("Search Error:", err)
		return []*api.Instance{}
	}

	return instances.Instances
}

func (h *Manager) UpdateInstances(q string, pageIndex, pageSize int) {
	instances := h.Search(q, pageIndex, pageSize)
	h.instanceTable.instances = instances
	h.instanceTable.Update()
}
