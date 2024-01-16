package service

import (
	"github.com/stretchr/testify/suite"
	"simple-cart/domain"
	"testing"
)

type CalculatorTestSuite struct {
	suite.Suite
	service Calculator
}

func (s *CalculatorTestSuite) SetupSuite() {
	s.service = NewCalculator(domain.PromotionsPerProduct, domain.PromotionsTotal)
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

func (s *CalculatorTestSuite) Test1Item() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Red",
				Count: 2,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal(float32(100), res.TotalPrice)
}

func (s *CalculatorTestSuite) TestRedAndGreen() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Red",
				Count: 1,
			},
			domain.CartItemReq{
				Name:  "Green",
				Count: 1,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal(float32(90), res.TotalPrice)
}

func (s *CalculatorTestSuite) Test2Item() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Red",
				Count: 2,
			},
			domain.CartItemReq{
				Name:  "Blue",
				Count: 4,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal(float32(220), res.TotalPrice)
}

func (s *CalculatorTestSuite) TestItemsWithMember() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Red",
				Count: 2,
			},
			domain.CartItemReq{
				Name:  "Blue",
				Count: 4,
			},
		},
		IsMember: true,
	}

	res := s.service.Execute(req)
	s.Equal((float32(198)), res.TotalPrice)
	s.Equal(float32(22), res.TotalDiscount)
}

func (s *CalculatorTestSuite) TestDoublePromotion2Item() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Orange",
				Count: 2,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal((float32(228)), res.TotalPrice)
	s.Equal(float32(12), res.TotalDiscount)
}

func (s *CalculatorTestSuite) TestDoublePromotion2ItemWithNormalItem() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Orange",
				Count: 2,
			},
			domain.CartItemReq{
				Name:  "Yellow",
				Count: 1,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal((float32(278)), res.TotalPrice)
	s.Equal(float32(12), res.TotalDiscount)
}

func (s *CalculatorTestSuite) TestDoublePromotion2ItemWithMemberPromotion() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Orange",
				Count: 2,
			},
			domain.CartItemReq{
				Name:  "Yellow",
				Count: 1,
			},
		},
		IsMember: true,
	}

	res := s.service.Execute(req)
	s.Equal((float32(250.2)), res.TotalPrice)
}

func (s *CalculatorTestSuite) TestDoublePromotion3Item() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Orange",
				Count: 3,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal((float32(342)), res.TotalPrice)
}

func (s *CalculatorTestSuite) TestDoublePromotion5ItemWithNormalItem() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Pink",
				Count: 5,
			},
			domain.CartItemReq{
				Name:  "Yellow",
				Count: 1,
			},
		},
		IsMember: false,
	}

	res := s.service.Execute(req)
	s.Equal((float32(430)), res.TotalPrice)
}

func (s *CalculatorTestSuite) TestDoublePromotion5ItemWithMemberPromotion() {
	req := domain.CalculatorReq{
		CartItemReq: []domain.CartItemReq{
			domain.CartItemReq{
				Name:  "Green",
				Count: 5,
			},
			domain.CartItemReq{
				Name:  "Yellow",
				Count: 1,
			},
		},
		IsMember: true,
	}

	res := s.service.Execute(req)
	s.Equal((float32(216)), res.TotalPrice)
}
