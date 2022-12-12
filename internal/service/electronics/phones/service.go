package phones

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Phone {
	return AllPhones
}

func (s *Service) Get(idx int) (*Phone, error) {
	return &AllPhones[idx], nil
}

func (s *Service) Remove(idx int) ([]Phone, error) {
	if len(AllPhones) == 0 {
		return nil, errors.New("phones are over")
	}

	return append(AllPhones[:idx], AllPhones[idx+1:]...), nil
}
