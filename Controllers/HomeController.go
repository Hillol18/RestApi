package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/nafi/oj-testing/Interfaces"
)

type HomeController struct {
	Interfaces.IHomeService
}

func (controller *HomeController) GetHome(req *fiber.Ctx) {
	req.Send(controller.GetHeading())
}
