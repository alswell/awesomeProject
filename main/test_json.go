package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func main() {
	jsonFile := filepath.Join("/home/zhouning/go/src/awesomeProject/main", "test_json.json")
	fmt.Println(jsonFile)
	str ,err := ioutil.ReadFile(jsonFile)
	fmt.Println(err, str)
	var jjj interface{}
	err = json.Unmarshal(str, &jjj)
	fmt.Println(err, jjj)
	fmt.Printf("%T\n", jjj.(map[string]interface{})["b"])
	cmd := exec.Command("bash", "-c", "pwd; sleep 2; pwd")
	stdout, err := cmd.StdoutPipe()
	fmt.Println(stdout, err)
	err = cmd.Start()
	fmt.Println(err)
	b, err := ioutil.ReadAll(stdout)
	fmt.Println(string(b), err)
}
