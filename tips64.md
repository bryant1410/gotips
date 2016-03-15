## #32 - Work with consul 
> 2016-14-03 by [@beyondns](https://github.com/beyondns)  

[Consul](https://www.consul.io/) has its own [client](https://github.com/hashicorp/consul/tree/master/api) but just use net/http 

```bash
consul agent -dev -advertise=127.0.0.1
```

```go
package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	consulKV = "http://127.0.0.1:8500/v1/kv/"
)

func httpRequest(meth, u, val string, timeLimit time.Duration) (int, []byte, error) {

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)

	var respStatus int
	var respBody []byte

	req, err := http.NewRequest(meth, u, strings.NewReader(val))
	if err != nil {
		return 0, nil, err
	}

	if val != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	go func() {
		resp, err := client.Do(req)

		if err != nil {
			goto E
		}

		respStatus = resp.StatusCode

		respBody, err = ioutil.ReadAll(resp.Body)
	E:
		c <- err
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	select {
	case <-time.After(timeLimit):
		tr.CancelRequest(req)
		log.Printf("Request timeout")
		<-c // Wait for goroutine to return.
		return 0, nil, errors.New("request time out")
	case err := <-c:
		if err != nil {
			log.Printf("Error in request goroutine %v", err)
			return 0, nil, err
		}
		return respStatus, respBody, nil
	}

}

//curl -v -XPUT http://127.0.0.1:8500/v1/kv/boom -d foo

func consulSet(k, v string) (int, []byte, error) {
	return httpRequest("PUT", consulKV+k, v, time.Second)
}

//curl -v http://127.0.0.1:8500/v1/kv/boom

func consulGet(k string) (int, []byte, error) {
	return httpRequest("GET", consulKV+k, "", time.Second)
}

func main() {

	s, d, err := consulSet("foo", "bar")
	if err != nil {
		log.Fatalf("etcdSet error %v", err)
	}
	log.Printf("set %d %s", s, string(d))

	s, d, err = consulGet("foo")
	if err != nil {
		log.Fatalf("etcdGet error %v", err)
	}
	log.Printf("get %d %s", s, string(d))

}
```