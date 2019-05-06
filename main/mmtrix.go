package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"metric_monitor/common/pcrypto"
	"path/filepath"
)

// rc4 key
var (
	DataKey = []byte{0x1f, 0x32, 0xb0, 0x29, 0xd7, 0x3f, 0xba, 0x56, 0x22, 0x71, 0xa8, 0x4d, 0x52, 0xe9, 0xa2, 0x3f}
)

func readAgtConf() {
	policyPath := filepath.Join("/home/zhouning/go/src/metric_monitor/itom_agent/data", "agtconf.json")
	b, err := ioutil.ReadFile(policyPath)
	fmt.Println(err)

	plain, err := pcrypto.Rc4Crypt(DataKey, b)
	var j interface{}
	json.Unmarshal(plain, &j)
	fmt.Println(j, err)
	b, err = json.MarshalIndent(j, "", "  ")
	fmt.Println(string(b), err)
}
func main() {
	readAgtConf()
	return
	ip := "localhost:13388"
	a, _ := pcrypto.Rc4Crypt(DataKey, []byte(ip))
	fmt.Println(base64.StdEncoding.EncodeToString(a))
	//plainUri, err := pcrypto.DecryPasswd(DataKey, "T/9rvVnvO/XvUjWmEPg9MOQWgQ==")
	//fmt.Println(plainUri, err)
}