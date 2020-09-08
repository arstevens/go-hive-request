package hrequest

import (
	ma "github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

/* NewVerificationResponse marshals a verification response value
using a protocol buffer */
func NewVerificationResponse(sigMap map[string][]byte) ([]byte, error) {
	signatures := make([]*VerificationResponse_SignaturePair, 0, len(sigMap))
	for id, signature := range sigMap {
		newSignature := &VerificationResponse_SignaturePair{
			Id:        id,
			Signature: signature,
		}
		signatures = append(signatures, newSignature)
	}

	verfResponse := &VerificationResponse{Signatures: signatures}
	response, err := proto.Marshal(verfResponse)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewVerificationResponse"}}
	}
	return response, nil
}

/* NewSubnetResponse marshals a set of subnet response values using a
protocol buffer */
func NewSubnetResponse(workers []string, bystanders []string, addrs map[string][]ma.Multiaddr, reward float64) ([]byte, error) {
	addresses := make([]*SubnetResponse_AddressPair, 0, len(addrs))
	for id, addr := range addrs {
		strAddr := make([]string, len(addr))
		for i, maAddr := range addr {
			strAddr[i] = maAddr.String()
		}
		newAddr := &SubnetResponse_AddressPair{
			Id:      id,
			Address: strAddr,
		}
		addresses = append(addresses, newAddr)
	}

	subResponse := &SubnetResponse{
		Workers:    workers,
		Bystanders: bystanders,
		Addresses:  addresses,
		Reward:     reward,
	}
	response, err := proto.Marshal(subResponse)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewSubnetResponse"}}
	}
	return response, nil
}

func NewReturnCodeResponse(code int) ([]byte, error) {
	returnResp := &ReturnCodeResponse{ReturnCode: int32(code)}
	response, err := proto.Marshal(returnResp)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewReturnCodeResponse"}}
	}
	return response, nil
}
