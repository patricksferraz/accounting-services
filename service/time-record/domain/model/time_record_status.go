package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["timeRecordStatus"] = govalidator.Validator(func(str string) bool {
		res := str == Pending.String()
		res = res || str == Approved.String()
		return res
	})
}

type TimeRecordStatus int

const (
	Pending TimeRecordStatus = iota + 1
	Approved
)

func (t TimeRecordStatus) String() string {
	switch t {
	case Pending:
		return "pending"
	case Approved:
		return "approved"
	}
	return ""
}
