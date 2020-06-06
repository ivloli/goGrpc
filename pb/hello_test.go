package hello

import (
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestHelloPb(t *testing.T) {
	req := &HelloRequest{
		Name: "Jack",
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		log.Printf("fail to marhsla req")
		return
	}

	req2 := &HelloRequest{}
	err = proto.Unmarshal(bytes, req2)
	if err != nil {
		log.Printf("fail to unmarshal req")
	}
	log.Printf("%v, %v", req2.Name, req.Name == req2.Name)

}
