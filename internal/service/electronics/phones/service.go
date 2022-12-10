package phones

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Phone {
	return allPhones
}

func (s *Service) Get(idx int) (*Phone, error) {
	return &allPhones[idx], nil
}
