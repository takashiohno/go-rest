package models

import (
    "bytes"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "github.com/takashiohno/go-rest/config"
    "github.com/takashiohno/go-rest/logs"
)


type EthRpcRequest struct {
    Jsonrpc string
    Method string
    Params string
    Id int
}

type EthBase struct {
    Id uint32 `json:"id"`
    Jsonrpc string `json:"jsonrpc"`
}

type EthAccount struct {
    EthBase
    Address string `json:"result"`
}

var defConf config.Config

func init() {
    defConf = config.DefaultConf()
}

func EthCoinbase() (string, error) {

    ereq := &EthRpcRequest{Jsonrpc: "2.0", Method: "eth_coinbase", Params: "", Id: 1}
    ethaccount := EthAccount{}
    err := ereq.rpcCall(&ethaccount)
    if err != nil {
        logs.Error.Println("rpc call failed ", err)
        return "", err
    }
    logs.Debug.Println("eth coinbase result is ", ethaccount)
    return ethaccount.Address, nil
}

func (ereq *EthRpcRequest) rpcCall(eb *EthAccount) (error) {

    parray := []string{
                `{"jsonrpc":"` + ereq.Jsonrpc + `"`,
                `"method":"` + ereq.Method + `"`,
                `"params":"` + ereq.Params + `"`,
                `"id":` + strconv.Itoa(ereq.Id) + `}`}
    reqp := strings.Join(parray, ",")
    logs.Debug.Println("request >>> ", reqp)

    resp, err := http.Post(defConf.RpcHost + ":" + strconv.Itoa(defConf.RpcPort),
                            "application/json",
                            bytes.NewBufferString(reqp))
    defer resp.Body.Close()
    if err != nil {
        return err
    }
    if errdec := json.NewDecoder(resp.Body).Decode(eb); errdec != nil {
        return errdec
    }
    logs.Debug.Println("response >>> ", eb)
    return nil
}
