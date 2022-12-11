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

//TODO make remove command for phones commander
func (s *Service) Remove(phones []Phone, idx int) []Phone {
	return append(phones[:idx], phones[idx+1:]...)
}
