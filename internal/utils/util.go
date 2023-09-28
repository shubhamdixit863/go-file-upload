package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func CreateGuid() string {

	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	guid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return guid

}
