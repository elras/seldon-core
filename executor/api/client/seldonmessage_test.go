package client

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	api "github.com/seldonio/seldon-core/executor/api/grpc"
	"testing"
)



func TestSum(t *testing.T) {
	var sm api.SeldonMessage
	var data = ` {"data":{"ndarray":[1.1,2.0]}}
`
	jsonpb.UnmarshalString(data, &sm)

	ma := jsonpb.Marshaler{}
	msgStr, _ := ma.MarshalToString(&sm)

	fmt.Printf("hello %s",msgStr)
}