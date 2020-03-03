package Repositories

import (
	"github.com/nafi/oj-testing/Interfaces"
)

type HomeRepository struct {
	Interfaces.IORMHandler
}

func (repository *HomeRepository) GetHeadingFromRepo() string {
	response := "No data found"

	return response
}
