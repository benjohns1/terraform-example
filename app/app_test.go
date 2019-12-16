package main

import (
	"github.com/aws/aws-lambda-go/events"
	"regexp"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name          string
		args          args
		wantMatchBody []byte
		wantErr       bool
	}{
		{
			name: "should return a Hello, World! message",
			args: args{
				request: events.APIGatewayProxyRequest{
					QueryStringParameters: map[string]string{"name": "World"},
				},
			},
			wantMatchBody: []byte(`"message":"Hello, World!"`),
		},
		{
			name: "should return the default hello message",
			args: args{
				request: events.APIGatewayProxyRequest{},
			},
			wantMatchBody: []byte(`"message":"Hello, my pretties!"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assertMatch(t, tt.wantMatchBody, got.Body)
		})
	}
}

func assertMatch(t *testing.T, wantMatch []byte, got string) {
	t.Helper()

	if wantMatch == nil {
		return
	}

	if match, err := regexp.MatchString(string(wantMatch), got); !match {
		t.Errorf("%s did not match expected pattern %s", got, wantMatch)
		return
	} else if err != nil {
		t.Fatalf("error matching pattern %s against %s", wantMatch, got)
		return
	}
}
