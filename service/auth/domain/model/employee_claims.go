package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeeClaims struct {
	ID    string   `mapstructure:"sub" valid:"uuid"`
	Roles []string `mapstructure:"roles" valid:"-"`
}

func (e *EmployeeClaims) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewEmployeeClaims(id string, roles []string) (*EmployeeClaims, error) {

	employeeClaims := EmployeeClaims{
		ID:    id,
		Roles: roles,
	}

	err := employeeClaims.isValid()
	if err != nil {
		return nil, err
	}

	return &employeeClaims, nil
}
