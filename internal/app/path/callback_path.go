package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Electronics  string
	Phones       string
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "_", 4)

	if len(callbackParts) != 4 {
		return CallbackPath{}, ErrUnknownCallback
	}

	return CallbackPath{
		Electronics:  callbackParts[0],
		Phones:       callbackParts[1],
		CallbackName: callbackParts[2],
		CallbackData: callbackParts[3],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s_%s_%s_%s", p.Electronics, p.Phones, p.CallbackName, p.CallbackData)
}
