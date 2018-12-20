package models

import "testing"


func TestGetAccount(t *testing.T) {

    account, res := GetAccount("")
    if !res {
        t.Errorf("result is [%t]", res)
    }
    if account == nil {
        t.Errorf("account is [%v]", account)
    }

    t.Log("finish")
}
