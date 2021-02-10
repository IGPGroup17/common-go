package http

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

var (
	headers = map[string]string {
		"Content-Type": "application/json",
		"X-Custom-Header": "application/json",
}
	JSON_Marshal_Response = events.APIGatewayProxyResponse{
		StatusCode:        INTERNAL_SERVER_ERROR.Code,
		Headers:           headers,
		MultiValueHeaders: nil,
		Body:              "Error Marshaling JSON Object!",
		IsBase64Encoded:   false,
	}
)

type ResponseEntity struct {
	ResponseObject interface{}
	Status         status
	Error          error
}

func (entity ResponseEntity) toJson() (string, error) {
	bytes, err := json.Marshal(&entity.ResponseObject)
	return string(bytes), err
}

func (entity ResponseEntity) toApiGatewayProxyResponse() events.APIGatewayProxyResponse {
	jsonStr, err := entity.toJson()

	if err != nil {
		return JSON_Marshal_Response
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode:        entity.Status.Code,
			Headers:           headers,
			MultiValueHeaders: nil,
			Body:              jsonStr,
			IsBase64Encoded:   false,
		}
	}
}
