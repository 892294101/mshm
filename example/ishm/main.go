package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/892294101/mshm/ishm"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

const MAXSIZE = 1 << 30

func WriteReadSHMI() {
	sm, err := ishm.CreateWithKey(12, MAXSIZE)
	if err != nil {
		log.Fatal(err)
		sm.Destroy()
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) // will write to buf

	shmi := ishm.SHMInfo{}
	lll := ishm.SizeStruct(shmi)
	fmt.Printf("shmisize:%v ,,,sizof:%v\n", lll, unsafe.Sizeof(shmi))
	shmi.MaxSHMSize = 100
	shmi.MaxContentLen = 64
	shmi.MaxTopicLen = 128
	shmi.Count = 4
	for i, _ := range shmi.Key {
		shmi.Key[i] = rand.Int31()
	}
	fmt.Printf("shm org:%#v\n", shmi)
	encoder.Encode(shmi)
	sm.Write(buf.Bytes())
	shmilen := buf.Len() //unsafe.Sizeof(shmi)

	fmt.Printf("buferlen:%v\n", shmilen)
	od, err := sm.ReadChunk(int64(shmilen), 0)
	if err != nil {
		//log.Fatal(err)
	}
	buf.Reset()
	buf.Write(od)
	decoder := gob.NewDecoder(&buf) // will read from buf
	smrd := ishm.SHMInfo{}
	err = decoder.Decode(&smrd)
	if err != nil {
		//log.Fatal(err)
	}
	fmt.Printf("shm read:%#v\n", smrd)
	fmt.Printf("sm:%#v\n", sm)
	sm.Destroy()
}
func testReadSHMByDefaultSHMI() {
	shmi, err := ishm.GetShareMemoryInfo(999999, false)
	if err != nil {
		log.Fatal(err)
	}
	for i, k := range shmi.Key {
		if i == int(shmi.Count) {
			break
		}
		fmt.Printf("key:%v\r\n", k)
		if uint64(i) < shmi.Count-1 {
			go func() {
				ishm.Readtlv(int64(k))
			}()
		} else {
			ishm.Readtlv(int64(k))
		}
		i++
	}

}

type TestJsonData struct {
	Name       string `json:"name"`
	DataLength int    `json:"dataLength"`
	Content    string `json:"content"`
}

func Producer() {

	shmParam := ishm.CreateSHMParam{Key: 4567, Size: 2000, Create: true}
	ctx := ishm.UpdateContent{EventType: "data-event", Topic: "xxx", Content: "正则表达式是一种进行模式匹配和文本操纵的功能强大的工具"}
	i, err := ishm.UpdateCtx(shmParam, ctx)
	if err != nil {
		fmt.Println("UpdateCtx error: ", err)
	} else {
		fmt.Printf("UpdateCtx success: %v %v\n", i, err)
	}

	for {
		time.Sleep(time.Second * 2)
		_, err := ishm.GetShareMemoryInfo(999999, false)
		if err != nil {
			fmt.Println("ishm.GetShareMemoryInfo: ", err)
		}

		shmParam.Create = false
		ishm.UpdateCtx(shmParam, ctx)
		res, err := ishm.GetCtx(shmParam)
		if err != nil {
			fmt.Println("ishm.GetCtx(shmParam): ", err)
		} else {
			log.Printf("read data: %s\n", res)
		}

	}

}
func main() {

	Producer()
}

