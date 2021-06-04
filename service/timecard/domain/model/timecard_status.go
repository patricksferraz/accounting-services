package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["timecardStatus"] = govalidator.Validator(func(str string) bool {
		res := str == Open.String()
		res = res || str == WaitingConfirmation.String()
		res = res || str == Confirmed.String()
		return res
	})
}

type TimecardStatus int

const (
	Open TimecardStatus = iota + 1
	WaitingConfirmation
	Confirmed
)

func (o TimecardStatus) String() string {
	switch o {
	case Open:
		return "open"
	case WaitingConfirmation:
		return "waiting confirmation"
	case Confirmed:
		return "confirmed"
	}
	return ""
}
