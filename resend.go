package tinkoff

import (
	"fmt"
)

type ResendRequest struct {
	BaseRequest
}

func (i ResendRequest) GetValuesForToken() map[string]string {
	return map[string]string{}
}

type ResendResponse struct {
	TerminalKey  string `json:"TerminalKey"`       // Идентификатор терминала, выдается Продавцу Банком
	Count        int    `json:"Count"`             // Количество сообщений, отправляемых повторно
	Success      bool   `json:"Success"`           // Успешность операции
	ErrorCode    string `json:"ErrorCode"`         // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"` // Краткое описание ошибки
	ErrorDetails string `json:"Details,omitempty"` // Подробное описание ошибки
}

func (c *Client) Resend(req *ResendRequest) (*ResendResponse, error) {
	response, err := c.postRequest("/Resend", req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res ResendResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorCode != "0" {
		return &res, fmt.Errorf(
			"while resend request: code %s - %s. %s",
			res.ErrorCode,
			res.ErrorMessage,
			res.ErrorDetails,
		)
	}
	return &res, nil
}
