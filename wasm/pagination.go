package main

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type Pagination struct {
	app.Compo
	manager *Manager

	currentPage int
	pageSize    int
}

func (p *Pagination) SetManager(manager *Manager) {
	p.manager = manager
}

// func (p *Pagination) Init() {
// 	p.currentPage = 1
// 	p.pageSize = 10
// }

func (p *Pagination) getPageSize() int {
	return p.pageSize
}

func (p *Pagination) setPageSize(pageSize int) {
	p.pageSize = pageSize
}

func (p *Pagination) getCurrentPage() int {
	return p.currentPage
}

func (p *Pagination) setCurrentPage(currentPage int) {
	p.currentPage = currentPage
}

func (p *Pagination) Render() app.UI {

	return app.Nav().Style("float", "right").Body(
		// app.Ul().Class("pagination").Body(
		// 	app.Li().Class("page-item").Style(
		// 		"padding", "3px 0",
		// 	).Body(
		// 		app.A().Class("page-link").Body(
		// 			app.Span().Body(app.Text("prev")).OnClick(p.prev),
		// 		),
		// 	),

		// 	app.Li().Class("page-item").Body(
		// 		app.A().Class("page-link").Body(
		// 			app.Span().Body(app.Text("next")).OnClick(p.next),
		// 		),
		// 	),
		// ),
		app.Input().
			Class("form-control").
			Value(p.currentPage).
			Placeholder("t2.small").
			AutoFocus(true).
			OnKeyup(p.OnInputChange),
	)
}

func (p *Pagination) OnInputChange(src app.Value, e app.Event) {
	p.currentPage, _ = strconv.Atoi(src.Get("value").String())

	p.Update()
	p.manager.UpdateInstances(p.manager.searchBar.searchString, p.currentPage, p.pageSize)
}

func (p *Pagination) next(src app.Value, e app.Event) {
	fmt.Println(src)
	p.currentPage++
	// p.manager.searchBar.Update()
}

func (p *Pagination) prev(src app.Value, e app.Event) {
	fmt.Println(src)
	p.currentPage--
	// p.manager.searchBar.Update()
}
