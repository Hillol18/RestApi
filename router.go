package main

import (
	"sync"

	"github.com/gofiber/fiber"
)

//IFiberRouter ...
type IFiberRouter interface {
	InitRouter() *fiber.Fiber
}

type Router struct{}

func (router *Router) InitRouter() *fiber.Fiber {

	r := fiber.New()

	homeController := ServiceContainer().InjectHomeController()

	r.Get("/", homeController.GetHome)

	return r
}

var (
	m          *Router
	routerOnce sync.Once
)

//FiberRouter ...
func FiberRouter() IFiberRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &Router{}
		})
	}
	return m
}
