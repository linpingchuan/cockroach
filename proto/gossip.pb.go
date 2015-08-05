// Code generated by protoc-gen-gogo.
// source: cockroach/proto/gossip.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// GossipRequest is the request struct passed with the Gossip RPC.
type GossipRequest struct {
	// Requesting node's ID.
	NodeID NodeID `protobuf:"varint,1,opt,name=node_id,casttype=NodeID" json:"node_id"`
	// Address of the requesting client.
	Addr Addr `protobuf:"bytes,2,opt,name=addr" json:"addr"`
	// Local address of client on requesting node (this is a kludge to
	// allow gossip to know when client connections are dropped).
	LAddr Addr `protobuf:"bytes,3,opt,name=l_addr" json:"l_addr"`
	// Maximum sequence number of gossip from this peer.
	MaxSeq int64 `protobuf:"varint,4,opt,name=max_seq" json:"max_seq"`
	// Reciprocal delta of new info since last gossip.
	Delta            []byte `protobuf:"bytes,5,opt,name=delta" json:"delta,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GossipRequest) Reset()         { *m = GossipRequest{} }
func (m *GossipRequest) String() string { return proto1.CompactTextString(m) }
func (*GossipRequest) ProtoMessage()    {}

func (m *GossipRequest) GetAddr() Addr {
	if m != nil {
		return m.Addr
	}
	return Addr{}
}

func (m *GossipRequest) GetLAddr() Addr {
	if m != nil {
		return m.LAddr
	}
	return Addr{}
}

func (m *GossipRequest) GetMaxSeq() int64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

func (m *GossipRequest) GetDelta() []byte {
	if m != nil {
		return m.Delta
	}
	return nil
}

// GossipResponse is returned from the Gossip.Gossip RPC.
// Delta will be nil in the event that Alternate is set.
type GossipResponse struct {
	// Requested delta of server's infostore.
	Delta []byte `protobuf:"bytes,1,opt,name=delta" json:"delta,omitempty"`
	// Non-nil means client should retry with this address.
	Alternate        *Addr  `protobuf:"bytes,2,opt,name=alternate" json:"alternate,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GossipResponse) Reset()         { *m = GossipResponse{} }
func (m *GossipResponse) String() string { return proto1.CompactTextString(m) }
func (*GossipResponse) ProtoMessage()    {}

func (m *GossipResponse) GetDelta() []byte {
	if m != nil {
		return m.Delta
	}
	return nil
}

func (m *GossipResponse) GetAlternate() *Addr {
	if m != nil {
		return m.Alternate
	}
	return nil
}
