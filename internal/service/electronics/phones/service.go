package phones

import "errors"

type ElectronicsPhonesService interface {
	List() []Phone
	Get(idx int) (*Phone, error)
	Remove(idx int) ([]Phone, error)
	New(phone Phone) ([]Phone, error)
	Edit(phone Phone, updPhone Phone) ([]Phone, error)
	GetLastID() uint64
}

type PhonesService struct{}

func NewService() *PhonesService {
	return &PhonesService{}
}

func (s *PhonesService) List() []Phone {
	return AllPhones
}

func (s *PhonesService) Get(ID uint64) (*Phone, error) {
	var result Phone

	for _, phone := range AllPhones {
		if phone.ID == ID {
			result = phone
			break
		}
	}

	return &result, nil
}

func (s *PhonesService) Remove(ID uint64) ([]Phone, error) {
	if len(AllPhones) == 0 {
		return nil, errors.New("phones are over")
	}

	var result []Phone

	for _, phone := range AllPhones {
		if phone.ID == ID {
			continue
		}

		result = append(result, phone)
	}

	AllPhones = result

	return AllPhones, nil
	/* Remove by index
	return append(AllPhones[:idx], AllPhones[idx+1:]...), nil
	*/
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

func (s *PhonesService) GetLastID() uint64 {
	return AllPhones[len(AllPhones)-1].ID
}
