package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
)

type RPCRequest struct {
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ParamTypes []string      `json:"param_types"`
	Id         int           `json:"id"`
}

type RPCResponse struct {
	Result     string `json:"result"`
	ResultType string `json:"result_type"`
	Id         int    `json:"id"`
}

type RPCFunction func([]interface{}) interface{}

func main() {
	sAddr := net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 8090,
		Zone: "",
	}

	tcpListener, err := net.ListenTCP("tcp", &sAddr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn *net.TCPConn) {
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var request RPCRequest
	err = json.Unmarshal(buff[:n], &request)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	functionRegistry := map[string]RPCFunction{
		"floor":        floor,
		"nroot":        nroot,
		"reverse":      reverse,
		"validAnagram": validAnagram,
		"sort":         sort,
	}

	result := functionRegistry[request.Method](request.Params)
	responce := RPCResponse{
		Result:     fmt.Sprintf("%v", result),
		ResultType: fmt.Sprintf("%T", result),
		Id:         request.Id,
	}

	marshaledResponse, err := json.Marshal(responce)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	conn.Write([]byte(marshaledResponse))
	conn.Close()
}

func floor(params []interface{}) interface{} {
	return math.Floor((params[0]).(float64))
}

func nroot(params []interface{}) interface{} {
	n := int(params[0].(float64))
	x := int(params[1].(float64))
	return math.Pow(float64(x), 1.0/float64(n))
}

func reverse(params []interface{}) interface{} {
	s := params[0].(string)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func validAnagram(params []interface{}) interface{} {
	str1 := params[0].(string)
	str2 := params[1].(string)

	if len(str1) != len(str2) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range str1 {
		charCount[char]++
	}

	for _, char := range str2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	return true
}

func sort(params []interface{}) interface{} {
	strSlice := make([]string, len(params))
	for i, v := range params {
		strSlice[i] = v.(string)
	}

	for i := 0; i < len(strSlice); i++ {
		for j := 0; j < len(strSlice)-1-i; j++ {
			if strSlice[j] > strSlice[j+1] {
				strSlice[j], strSlice[j+1] = strSlice[j+1], strSlice[j]
			}
		}
	}

	return strSlice
}
