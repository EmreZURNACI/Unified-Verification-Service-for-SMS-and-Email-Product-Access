package helpers

import (
	"encoding/json"
	"fmt"

	m "Models.go"
)

var dbMessage m.DatabaseMessage

func Response(_status bool, _statuscode int, _message string, _data []m.Product) string {
	res := m.StandartResponseModel{
		Status:     _status,
		StatusCode: _statuscode,
		Message:    _message,
		Data:       _data,
	}
	bs, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		panic(err)
	}
	return string(bs)
}
func ExtractStatuAndMessage(result string) (bool, string) {
	err := json.Unmarshal([]byte(result), &dbMessage)
	if err != nil {
		fmt.Println("Unmarshal hatası")
	}
	if dbMessage.Statu == "error" {
		return false, dbMessage.Message
	} else {
		return true, dbMessage.Message
	}
}
