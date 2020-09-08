package hrequest

import (
	"crypto"

	snapshot "github.com/arstevens/go-snapshot"
	"github.com/gogo/protobuf/proto"
	ma "github.com/multiformats/go-multiaddr"
)

type PublicKeyMarshaler func(crypto.PublicKey) []byte

func NewVerificationRequest(snap *snapshot.SimpleSnapshot, keys map[string]crypto.PublicKey,
	keymarsh PublicKeyMarshaler) ([]byte, error) {
	serialSnap, err := snap.Marshal()
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewVerificationRequest"}}
	}

	keypairs := make([]*VerificationRequest_Keypair, 0, len(keys))
	for id, key := range keys {
		serialKey := keymarsh(key)
		pair := &VerificationRequest_Keypair{Id: id, Key: serialKey}
		keypairs = append(keypairs, pair)
	}

	verfRequest := &VerificationRequest{Snapshot: serialSnap, Keys: keypairs}
	request, err := proto.Marshal(verfRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewVerificationRequest"}}
	}
	return request, nil
}

func NewSubnetRequest(bystanders int, storage int, compute int, availability float32) ([]byte, error) {
	workerOpts := &WorkerParameters{
		Storage:      int32(storage),
		Compute:      int32(compute),
		Availability: availability,
	}

	subRequest := &SubnetRequest{
		TotalBystanders: int32(bystanders),
		WorkerOpts:      workerOpts,
	}
	request, err := proto.Marshal(subRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewSubnetRequest"}}
	}
	return request, nil
}

func NewConflictRequest(epoch *snapshot.SimpleEpochTriplet, serverSig string, senderSig string) ([]byte, error) {
	serialEpoch, err := epoch.Marshal()
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewConflictRequest"}}
	}

	confRequest := &ConflictRequest{
		ServerSignature: serverSig,
		SenderSignature: senderSig,
		Epoch:           serialEpoch,
	}
	request, err := proto.Marshal(confRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewConflictRequest"}}
	}
	return request, nil
}

func NewAvailabilitySetRequest(id string, storage int, compute int, availability float32) ([]byte, error) {
	availabilityOpts := &WorkerParameters{
		Storage:      int32(storage),
		Compute:      int32(compute),
		Availability: availability,
	}

	availRequest := &AvailabilitySetRequest{
		Id:               id,
		AvailabilityOpts: availabilityOpts,
	}
	request, err := proto.Marshal(availRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewAvailabilitySetRequest"}}
	}
	return request, nil
}

func NewMultiaddressSetRequest(id string, addrs []ma.Multiaddr) ([]byte, error) {
	strAddrs := make([]string, len(addrs))
	for i, addr := range addrs {
		strAddrs[i] = addr.String()
	}

	multiRequest := &MultiaddressSetRequest{
		Id:             id,
		Multiaddresses: strAddrs,
	}
	request, err := proto.Marshal(multiRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewMultiaddressSetRequest"}}
	}
	return request, nil
}

/* NewInitialRequest returns a new instance of InitialRequest
using an underlying protocol buffer */
func NewInitialRequest(types []int32, serial []byte) ([]byte, error) {
	initRequest := &InitialRequest{
		Types:  types,
		Serial: serial,
	}
	request, err := proto.Marshal(initRequest)
	if err != nil {
		return nil, &MarshalErr{simpleErr{err: err, msg: "MarshalErr in NewInitialRequest"}}
	}
	return request, nil
}
