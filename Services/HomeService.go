package Services

import (
	"github.com/nafi/oj-testing/Interfaces"
)

type HomeService struct {
	Interfaces.IHomeRepository
}

func (service *HomeService) GetHeading() string {
	return service.GetHeadingFromRepo()
}
