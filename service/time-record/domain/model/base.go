package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `bson:"_id" valid:"uuid"`
	CreatedAt time.Time `bson:"created_at" valid:"required"`
	UpdatedAt time.Time `bson:"updated_at" valid:"-"`
}
