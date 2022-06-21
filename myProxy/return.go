package myProxy

import "huaweicloud.com/go-runtime/events/apig"

func returnMsg(bodyText string, headers map[string]string, StatusCode int, err error) (apig.APIGTriggerResponse, error) {
	if err != nil {
		return apig.APIGTriggerResponse{}, err
	}
	return apig.APIGTriggerResponse{
		Body:       bodyText,
		Headers:    headers,
		StatusCode: StatusCode,
	}, nil
}
