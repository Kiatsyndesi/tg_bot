package phones

import "errors"

type ElectronicsPhonesService interface {
	List() []Phone
	Get(idx int) (*Phone, error)
	Remove(idx int) ([]Phone, error)
	New(phone Phone) ([]Phone, error)
	Edit(phone Phone, updPhone Phone) ([]Phone, error)
}

type PhonesService struct{}

func NewService() *PhonesService {
	return &PhonesService{}
}

func (s *PhonesService) List() []Phone {
	return AllPhones
}

func (s *PhonesService) Get(idx int) (*Phone, error) {
	return &AllPhones[idx], nil
}

func (s *PhonesService) Remove(idx int) ([]Phone, error) {
	if len(AllPhones) == 0 {
		return nil, errors.New("phones are over")
	}

	return append(AllPhones[:idx], AllPhones[idx+1:]...), nil
}

func (s *PhonesService) New(phone Phone) ([]Phone, error) {
	var emptyPhone Phone

	if phone == emptyPhone {
		return nil, errors.New("new phone is empty")
	}

	return append(AllPhones, phone), nil
}

func (s *PhonesService) Edit(phone Phone, updPhone Phone) ([]Phone, error) {
	hasPhone := false
	var idxForEdit int

	for idx, existingPhone := range AllPhones {
		if existingPhone == phone {
			idxForEdit = idx
			hasPhone = true
		}
	}

	if !hasPhone {
		return nil, errors.New("service doesn't have this phone")
	}

	AllPhones[idxForEdit] = updPhone

	return AllPhones, nil
}
