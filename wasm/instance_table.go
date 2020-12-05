package main

import (
	"bazel-golang-wasm-protoc/protos/api"
	"fmt"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type InstanceTable struct {
	app.Compo
	manager   *Manager
	instances []*api.Instance
}

func (p *InstanceTable) SetManager(manager *Manager) {
	p.manager = manager
}

func (p *InstanceTable) Render() app.UI {

	nodes := []app.Node{}
	for _, i := range p.instances {
		nodes = append(nodes, app.Tr().Body(
			app.Td().Body(app.Text(i.ZoneID)),
			// app.Td().Body(app.Text(i.NetworkTypes)),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.ZoneStatue))),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.EngineVersion))),
			app.Td().Body(app.Text(i.Category)),
			app.Td().Body(app.Text(i.StorageType)),
			app.Td().Body(app.Text(i.DBInstanceClass)),
		))
	}

	return app.Table().Class("table").Body(
		app.Tr().Body(
			app.Th().Scope("col").Body(app.Text("ZoneID")),
			// app.Th().Scope("col").Body(app.Text("Network Type")),
			app.Th().Scope("col").Body(app.Text("ECU")),
			app.Th().Scope("col").Body(app.Text("Mem")),
			app.Th().Scope("col").Body(app.Text("Network")),
			app.Th().Scope("col").Body(app.Text("DBInstanceClass")),
			app.Th().Scope("col").Body(app.Text("Price")),
		),
		app.TBody().Body(nodes...),
	)

}
