package models

type NewAccount struct {
    Password string `json:"password"`
}

type Account struct {
    Address string `json:"address"`
    Wei uint64 `json:"wei"`
    Ether uint32 `json:"ether"`
}

type Accounts []*Account

func GetAccounts() (Accounts, int32, error) {
    return nil, 0, nil
}

func GetAccount(addr string) (*Account, bool) {

    // test
    acc := Account{"0xd95d24b85e209b327e91c7e63f9e36df65bec65a", 2000000000000000000, 2}
    return &acc, true
}

func CreateAccount(nac NewAccount) (*Account, bool) {

    // test
    acc := Account{"0xnewaccount", 0, 0}
    return &acc, true
}
