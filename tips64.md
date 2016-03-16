## #33 - Table driven tests 
> 2016-15-03 by [@beyondns](https://github.com/beyondns)  

Table driven tests very simple but efficient. An example shamelessly stolen from [advanced-testing-with-go](https://speakerdeck.com/mitchellh/advanced-testing-with-go)

```go
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct{ A, B, Sum int }{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
		{3, 2, 1}, // error!!!
	}

	for _, c := range cases {
		a := c.A + c.B
		e := c.Sum
		if a != e {
			t.Errorf("%d + %d = %d, expected %d", c.A, c.B, a, e)
		}
	}

}

```
```bash
--- FAIL: TestAdd (0.00s)
	tdt_test.go:20: 3 + 2 = 5, expected 1
```
* [advanced-testing-with-go](https://speakerdeck.com/mitchellh/advanced-testing-with-go)

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
	"encoding/json"
)

// https://www.consul.io/docs/agent/http/kv.html

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

type KVPair struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       []byte
	Session     string
}

func main() {

	s, d, err := consulSet("foo", "hero")
	if err != nil {
		log.Fatalf("Set error %v", err)
	}
	log.Printf("set %d %s", s, string(d))

	s, d, err = consulGet("foo")
	if err != nil {
		log.Fatalf("Get error %v", err)
	}

	if s!=200{
		log.Fatalf("Get status %d", s)
	}

    var kv []KVPair	
	json.Unmarshal(d,&kv)

	if len(kv)==0{
		log.Fatalf("len(kv)==0")
	}

	log.Printf("get %d %s %s", s, string(d), kv[0].Value)

}
```
```bash
2016/03/15 20:49:54 set 200 true
2016/03/15 20:49:54 get 200 [{"LockIndex":0,"Key":"foo","Flags":0,"Value":"aGVybw==","CreateIndex":7,"ModifyIndex":252}] hero

```