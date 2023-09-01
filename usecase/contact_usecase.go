package usecase

import "tdd-go-fiber/domain"

type RealContactUseCase struct{}

func NewContactUseCase() *RealContactUseCase {
	return &RealContactUseCase{}
}

func (cu *RealContactUseCase) GetContacts() ([]domain.Contact, error) {
	return []domain.Contact{}, nil
}
