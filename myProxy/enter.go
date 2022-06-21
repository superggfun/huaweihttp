package myProxy

import (
	"errors"
	"strings"

	"huaweicloud.com/go-runtime/events/apig"
)

type APIG struct {
	apig.APIGTriggerEvent
}

func NewAPIG(i apig.APIGTriggerEvent) *APIG {
	var apig APIG
	apig.HttpMethod = i.HttpMethod
	apig.Path = i.Path
	apig.Body = i.Body
	apig.PathParameters = i.PathParameters
	apig.RequestContext = i.RequestContext
	apig.Headers = i.Headers
	apig.QueryStringParameters = i.QueryStringParameters
	apig.UserData = i.UserData
	return &apig
}

func (m *APIG) Start() (interface{}, error) {
	if _, ok := m.QueryStringParameters["url"]; !ok {
		return returnMsg(m.noUrl())
	}
	switch m.HttpMethod {
	case "GET":
		return returnMsg(m.myGet())
	case "POST":
		return returnMsg(m.myPost())
	default:
		return m.HttpMethod, errors.New("httpMethod error")
	}
}

func (m *APIG) getUrl() string {
	url := m.QueryStringParameters["url"]
	delete(m.QueryStringParameters, "url")
	Parameters := "?"
	if len(m.QueryStringParameters) != 0 {
		for key, value := range m.QueryStringParameters {
			Parameters = strings.Join([]string{Parameters, key, "=", value, "&"}, "")
		}
		url += Parameters[:len(Parameters)-1]
	} else {
		url += "/"
	}
	return url
}

func (m *APIG) delHeader() {
	delete(m.Headers, "host")
	delete(m.Headers, "x-forwarded-host")
	delete(m.Headers, "x-forwarded-port")
	delete(m.Headers, "x-request-id")
	delete(m.Headers, "x-forwarded-for")
	delete(m.Headers, "x-forwarded-proto")
	delete(m.Headers, "x-real-ip")
	delete(m.Headers, "accept-encoding")
}
