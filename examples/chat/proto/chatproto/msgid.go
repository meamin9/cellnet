// Generated by cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: chat.proto

package chatproto

import (
	"cellnet"
	"reflect"
	_ "cellnet/codec/pb"
)

func init() {

	// chat.proto
	cellnet.RegisterMessageMeta("pb", "chatproto.ChatREQ", reflect.TypeOf((*ChatREQ)(nil)).Elem(), 3634688514)
	cellnet.RegisterMessageMeta("pb", "chatproto.ChatACK", reflect.TypeOf((*ChatACK)(nil)).Elem(), 3019771488)

}
