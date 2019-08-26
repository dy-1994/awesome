package main

import (
	"testing"
	"net"
	"fmt"
	"time"
	"awesome/utils"
	"awesome/parser"
)

func Test_login(t *testing.T)  {
	err := login(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDial(t *testing.T) {
	con, err := net.Dial("tcp", ":1234")
	if err != nil {
		t.Fatal(err)
	}

	pr := parser.NewFrame(con)
	go func() {
		defer con.Close()
		for {
			msg, err := pr.ReadFrame()
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Println(string(msg))
		}
	}()

	if _, err = con.Write([]byte("12325345324")); err != nil {
		t.Fatal(err)
	}

	if err = pr.WriteFrame([]byte("test")); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second*10)
}

func TestUint32BytesConvert(t *testing.T) {
	test := uint32(123456)
	bs := utils.Uint322Bytes(test)
	t.Log(bs)
	t.Log(utils.Bytes2Uint32(bs))
	if test != utils.Bytes2Uint32(bs) {
		t.Fatal("failure")
	}
}