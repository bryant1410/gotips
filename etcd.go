package main

import (
	"errors"
	"log"
	"time"
	"io/ioutil"
	"strings"
	"net/http"
)

var (
	etcdKeys = "http://127.0.0.1:2379/v2/keys/"
)

func httpRequest(meth, u, val string) (int,[]byte, error) {

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)

	var respStatus int
	var respBody []byte

	req, err := http.NewRequest(meth, u, strings.NewReader(val))
	if err != nil {
		return 0,nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	go func() {
		resp, err := client.Do(req)

		if err != nil {
			goto E
		}

		respStatus = resp.StatusCode

		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			goto E
		}
	E:
		c <- err
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	select {
	case <-time.After(time.Second * 10):
		tr.CancelRequest(req)
		log.Printf("Request timeout")
		<-c // Wait for goroutine to return.
		return 0,nil, errors.New("request time out")
	case err := <-c:
		if err != nil {
			log.Printf("Error in request goroutine %v", err)
			return 0,nil, err
		}
		return respStatus,respBody, nil
	}

}

/*

curl -L -X PUT http://127.0.0.1:2379/v2/keys/message -d value="Hello"
{"action":"set","node":{"key":"/message","value":"Hello","modifiedIndex":4,"createdIndex":4}}

*/

func etcdSet(k, v string) (int,[]byte, error) {
	return httpRequest("PUT", etcdKeys+k, "value="+v)
}

func etcdGet(k string) (int,[]byte, error) {
	return httpRequest("GET", etcdKeys+k, "")
}

func main() {
	s,d, err := etcdSet("foo", "bar")
	if err != nil {
		log.Fatalf("etcdSet error %v", err)
	}
	log.Printf("set %d %s",s,string(d))
	s, d, err = etcdGet("foo")
	if err != nil {
		log.Fatalf("etcdGet error %v", err)
	}
	log.Printf("get %d %s",s,string(d))
}
