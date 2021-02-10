package http

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

var (
	headers = map[string]string{
		"Content-Type":    "application/json",
		"X-Custom-Header": "application/json",
	}
	JsonMarshalResponse = events.APIGatewayProxyResponse{
		StatusCode:        INTERNAL_SERVER_ERROR.Code,
		Headers:           headers,
		MultiValueHeaders: nil,
		Body:              "Error Marshaling JSON Object!",
		IsBase64Encoded:   false,
	}
)

type ResponseEntity struct {
	ResponseObject interface{}
	Status         Status
	Error          error
}

func (entity ResponseEntity) toJson() (string, error) {
	bytes, err := json.Marshal(&entity.ResponseObject)
	return string(bytes), err
}

func (entity ResponseEntity) ToApiGatewayProxyResponse() events.APIGatewayProxyResponse {
	jsonStr, err := entity.toJson()

	if err != nil {
		return JsonMarshalResponse
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
