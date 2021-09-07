package ishm

import (
	"testing"
)

func TestDefaultConverter(t *testing.T) {
	codec := "default"
	RegisterConverter(codec, DefaultConverter{})
	ts := &TestStruct{
		Name: "Tencent",
		Desc: "666",
	}
	data, _ := Encode(ts)
	tsRecv := &TestStruct{}
	Decode(data, tsRecv)
}
