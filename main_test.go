package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	url := "https://api.live.bilibili.com/msg/send"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("bubble", "0")
	_ = writer.WriteField("msg", "?")
	_ = writer.WriteField("color", "16777215")
	_ = writer.WriteField("mode", "1")
	_ = writer.WriteField("room_type", "0")
	_ = writer.WriteField("jumpfrom", "0")
	_ = writer.WriteField("reply_mid", "0")
	_ = writer.WriteField("reply_attr", "0")
	_ = writer.WriteField("replay_dmid", "")
	_ = writer.WriteField("statistics", "{\"appId\":100,\"platform\":5}")
	_ = writer.WriteField("reply_type", "0")
	_ = writer.WriteField("reply_uname", "")
	_ = writer.WriteField("data_extend", "{\"trackid\":\"-99998\"}")
	_ = writer.WriteField("fontsize", "25")
	_ = writer.WriteField("rnd", "177623")
	_ = writer.WriteField("roomid", "177623")
	_ = writer.WriteField("csrf", "8da3af03b4a04044d01249b756b5f825")
	_ = writer.WriteField("csrf_token", "8da3af03b4a04044d01249b756b5f825")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "SESSDATA=f42dc0a3%2C1769776771%2Cb4ac9%2A81CjAN4cVqoi9spS052Bw0N6M8z4ATH7_fnqEFhANpBbNWcNnSaQAOnngHPBSo4ToWLjsSVjlMcTIwMUsxRkE5SVcyaURqUVpVTmNjTktKT2YydEIyZXVYVDd5S0x4c004MVByYW9EYmtPOWxrSURtOFRwcjlIUDk0c0k5NHZ1bzYzZEcxSEtXNVFRIIEC")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
