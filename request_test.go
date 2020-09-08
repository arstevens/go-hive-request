package hrequest

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestRequest(t *testing.T) {
	subreq, err := NewSubnetRequest(5, 6, 7, 9.0)
	if err != nil {
		panic(err)
	}
	req, err := NewInitialRequest([]int32{0, 1, 2}, subreq)
	if err != nil {
		panic(err)
	}
	fmt.Println(subreq)
	fmt.Println(req)

	var response SubnetRequest
	err = proto.Unmarshal(subreq, &response)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}
