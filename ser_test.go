package serbench

import (
	"encoding/json"
	"encoding/xml"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/ugorji/go/codec"
	msgpack1 "github.com/vmihailenco/msgpack"
	"log"
	"testing"
)

var user = User{
	Id:   10,
	Name: "FancyFancyFancyFancyFancy",
	Item: []string{"Table", "Door", "Ball", "Knife"},
}

var protoUser = UserProto{
	Id:   10,
	Name: "FancyFancyFancyFancyFancy",
	Item: []string{"Table", "Door", "Ball", "Knife"},
}

var gogoUser = GogoUserProto{
	Id:   10,
	Name: "FancyFancyFancyFancyFancy",
	Item: []string{"Table", "Door", "Ball", "Knife"},
}

func TestMarshalDataLen(t *testing.T) {
	log.SetFlags(log.LstdFlags)

	buf, _ := json.Marshal(user)
	t.Logf("json:\t\t\t\t %d bytes", len(buf))

	buf, _ = xml.Marshal(user)
	t.Logf("xml:\t\t\t\t %d bytes", len(buf))

	buf, _ = user.MarshalMsg(nil)
	t.Logf("msgp:\t\t\t\t %d bytes", len(buf))

	buf, _ = proto.Marshal(&protoUser)
	t.Logf("proto:\t\t\t\t %d bytes", len(buf))

	buf, _ = goproto.Marshal(&gogoUser)
	t.Logf("gogo:\t\t\t\t %d bytes", len(buf))
}

func BenchmarkMarshalJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(user)
	}
}
func BenchmarkUnmarshalJson(b *testing.B) {
	bytes, _ := json.Marshal(user)
	result := User{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalXml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xml.Marshal(user)
	}
}
func BenchmarkUnmarshalXml(b *testing.B) {
	bytes, _ := json.Marshal(user)
	result := User{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xml.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalMsgp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user.MarshalMsg(nil)
	}
}
func BenchmarkUnmarshalMsgp(b *testing.B) {
	bytes, _ := user.MarshalMsg(nil)
	result := User{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalMsg(bytes)
	}
}

func BenchmarkMarshalMsgp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack1.Marshal(&user)
	}
}
func BenchmarkUnmarshalMsgp1(b *testing.B) {
	bytes, _ := msgpack1.Marshal(&user)
	b.ResetTimer()
	var user1 User
	for i := 0; i < b.N; i++ {
		msgpack1.Unmarshal(bytes, &user1)
	}
}

func BenchmarkMarshalMsgp2(b *testing.B) {
	var mh codec.MsgpackHandle
	var byteBuf []byte
	encoder := codec.NewEncoderBytes(&byteBuf, &mh)
	for i := 0; i < b.N; i++ {
		encoder.Encode(user)
	}
}
func BenchmarkUnmarshalMsgp2(b *testing.B) {
	var mh codec.MsgpackHandle
	var byteBuf []byte
	encoder := codec.NewEncoderBytes(&byteBuf, &mh)
	encoder.Encode(user)

	var byteBuf_2 []byte
	decoder := codec.NewDecoderBytes(byteBuf_2, &mh)
	b.ResetTimer()
	var user1 User
	for i := 0; i < b.N; i++ {
		decoder.Decode(&user1)
	}
}

func BenchmarkMarshalProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(&protoUser)
	}
}
func BenchmarkUnmarshalProto(b *testing.B) {
	bytes, _ := proto.Marshal(&protoUser)
	result := UserProto{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalGogo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goproto.Marshal(&gogoUser)
	}
}
func BenchmarkUnmarshalGogo(b *testing.B) {
	bytes, _ := proto.Marshal(&gogoUser)
	result := GogoUserProto{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goproto.Unmarshal(bytes, &result)
	}
}
