package base

import (
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
)

const (
    get = "GET"
    post = "POST"
    put = "PUT"
    delete = "DELETE"
)

type Resource interface {
    Get(url string, values url.Values) (int, interface{})
    Post(url string, values url.Values, decoder *json.Decoder) (int, interface{})
    Put(url string, values url.Values, decoder *json.Decoder) (int, interface{})
    Delete(url string, values url.Values) (int, interface{})
}

type ResourceBase struct {}

func (ResourceBase) Get(url string, values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func (ResourceBase) Post(url string, values url.Values, decoder *json.Decoder) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func (ResourceBase) Put(url string, values url.Values, decoder *json.Decoder) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func (ResourceBase) Delete(url string, values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func RequestHandler(resource Resource) http.HandlerFunc {
    return func(rw http.ResponseWriter, rq *http.Request) {

        //b := bytes.NewBuffer(make([]byte, 0))
        //reader := io.TeeReader(rq.Body, b)
        //rq.Body = ioutil.NopCloser(b)
        decoder := json.NewDecoder(rq.Body)
        defer rq.Body.Close()

        rq.ParseForm()

        url := formatUrlString(rq.URL.Path)

        var code int
        var data interface{}

        switch rq.Method {
        case get:
            code, data = resource.Get(url, rq.Form)
        case post:
            code, data = resource.Post(url, rq.Form, decoder)
        case put:
            code, data = resource.Put(url, rq.Form, decoder)
        case delete:
            code, data = resource.Delete(url, rq.Form)
        default:
            rw.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        var content []byte
        var err error
        content, err = json.Marshal(data)
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        rw.Header().Set("Content-Type", "application/json")
        rw.WriteHeader(code)
        rw.Write(content)
    }
}

func formatUrlString(url string) (string) {

    result := strings.TrimSpace(url)
    if !strings.HasSuffix(result, "/") {
        result = result + "/"
    }
    return result
}
