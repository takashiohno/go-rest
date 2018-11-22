package resources

import (
    "encoding/json"
    "net/http"
    "net/url"

    "github.com/takashiohno/go-rest/base"
    "github.com/takashiohno/go-rest/logs"
    "github.com/takashiohno/go-rest/models"
)

type account struct {
    base.ResourceBase
}

func init() {
    http.HandleFunc("/account", base.RequestHandler(account{}))
    http.HandleFunc("/account/", base.RequestHandler(account{}))
}

func (a account) Get(url string, values url.Values) (int, interface{}) {

    if id := url[len("/account/"):]; len(id) > 0 {
        account, result := models.GetAccount(id)
        if !result {
            return http.StatusOK, nil
        }
        return http.StatusOK, account
    }

    accounts, _, err := models.GetAccounts()
    if err != nil {
        logs.Error.Println("failed get account")
        return http.StatusInternalServerError, nil
    }
    return http.StatusOK, accounts
}

func (a account) Post(url string, values url.Values, decoder *json.Decoder) (int, interface{}) {

    newAccount := models.NewAccount{}
    if err := decoder.Decode(&newAccount); err != nil {
        logs.Error.Println("decode error")
        return http.StatusInternalServerError, nil
    }
    logs.Debug.Println("request body", newAccount.Password)

    nac, result := models.CreateAccount(newAccount)
    if !result {
        logs.Error.Println("failed create account")
        return http.StatusInternalServerError, nil
    }
    return http.StatusOK, nac
}
