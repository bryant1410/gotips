![header](https://raw.githubusercontent.com/beyondns/gotips/master/gophers.jpg)

> **Golang short tips & tricks**

This list of short golang code tips & tricks will help keep collected knowledge in one place.

> **Support the project**

Send some ether 0x30FD8822D15081F3c98e6A37264F8dF37a2EB416


# Tips list

- 58 - [set go env](https://github.com/beyondns/gotips#58---set-go-env)
- 57 - [Multiple concurrent http requests with timeout](https://github.com/beyondns/gotips#57---multiple-concurrent-http-requests-with-timeout)
- 56 - [Inplace struct](https://github.com/beyondns/gotips#56---inplace-struct)
- 55 - [Dynamic time intervals](https://github.com/beyondns/gotips#55---dynamic-time-intervals)
- 54 - [NoDB](https://github.com/beyondns/gotips#54---nodb)
- 53 - [EC operations from bitcoin core](https://github.com/beyondns/gotips#53---ec-operations-from-bitcoin-core)
- 52 - [nil](https://github.com/beyondns/gotips#52---nil)
- 51 - [BBQ](https://github.com/beyondns/gotips#51---bbq)
- 50 - [cool go links](https://github.com/beyondns/gotips#50---cool-go-links)
- 49 - [custom wait group](https://github.com/beyondns/gotips#49---custom-wait-group)
- 48 - [ev_loop vs co_routine](https://github.com/beyondns/gotips#48---ev_loop-vs-co_routine)
- 47 - [Use C for computation intensive operations](https://github.com/beyondns/gotips#47---use-c-for-computation-intensive-operations)
- 46 - [Unique key doc manipulation in MonGo](https://github.com/beyondns/gotips#46---unique-key-doc-manipulation-in-mongo)
- 45 - [Go is more a framework than a language](https://github.com/beyondns/gotips#45---go-is-more-a-framework-than-a-language)
- 44 - [Facebook messenger chatbot](https://github.com/beyondns/gotips#44---facebook-messenger-chatbot)
- 43 - [Passing args as a map](https://github.com/beyondns/gotips#43---passing-args-as-a-map)
- 42 - [Optimize Go](https://github.com/beyondns/gotips#42---optimize-go)
- 41 - [Telegram bot](https://github.com/beyondns/gotips#41---telegram-bot)
- 40 - [Distributed consensus](https://github.com/beyondns/gotips#40---distributed-consensus)
- 39 - [Public private struct fields and funcs](https://github.com/beyondns/gotips#39---public-private-struct-fields-and-funcs)
- 38 - [Marshal Unmarshal to byte slice using gob](https://github.com/beyondns/gotips#38---marshal-unmarshal-to-byte-slice-using-gob)
- 37 - [Time series on leveldb](https://github.com/beyondns/gotips#37---time-series-on-leveldb)
- 36 - [Custom type marshal unmarshal json](https://github.com/beyondns/gotips#36---custom-type-marshal-unmarshal-json)
- 35 - [Test specific functions](https://github.com/beyondns/gotips#35---test-specific-functions)
- 34 - [File path exists](https://github.com/beyondns/gotips#34---file-path-exists)
- 33 - [Table driven tests](https://github.com/beyondns/gotips#33---table-driven-tests)
- 32 - [Work with consul](https://github.com/beyondns/gotips#32---work-with-consul)
- 31 - [Do not need any web framework](https://github.com/beyondns/gotips/blob/master/tips32.md#31---do-not-need-any-web-framework)
- 30 - [Heart beat and lead election with etcd](https://github.com/beyondns/gotips/blob/master/tips32.md#30---heart-beat-and-lead-election-with-etcd)
- 29 - [Partial json read](https://github.com/beyondns/gotips/blob/master/tips32.md#29---partial-json-read)
- 28 - [Interact with etcd with http.Request](https://github.com/beyondns/gotips/blob/master/tips32.md#28---interact-with-etcd-with-httprequest)
- 27 - [Go-style concurrency in C](https://github.com/beyondns/gotips/blob/master/tips32.md#27---go-style-concurrency-in-c)
- 26 - [Go channels are slow](https://github.com/beyondns/gotips/blob/master/tips32.md#26---go-channels-are-slow)
- 25 - [Avoid conversions with hidden alloc copy](https://github.com/beyondns/gotips/blob/master/tips32.md#25---avoid-conversions-with-hidden-alloc-copy)
- 24 - [Slice shuffle](https://github.com/beyondns/gotips/blob/master/tips32.md#24---slice-shuffle)
- 23 - [Defer is slow](https://github.com/beyondns/gotips/blob/master/tips32.md#23---defer-is-slow)
- 22 - [Any number of args](https://github.com/beyondns/gotips/blob/master/tips32.md#22---any-number-of-args)
- 21 - [Test bench http request handler](https://github.com/beyondns/gotips/blob/master/tips32.md#21---test-bench-http-request-handler)
- 20 - [Use Atomics or GOMAXPROCS=1](https://github.com/beyondns/gotips/blob/master/tips32.md#20---use-atomics-or-gomaxprocs1)
- 19 - [Chunked HTTP response with flusher](https://github.com/beyondns/gotips/blob/master/tips32.md#19---chunked-http-response-with-flusher)
- 18 - [Use for range for channels](https://github.com/beyondns/gotips/blob/master/tips32.md#18---use-for-range-for-channels)
- 17 - [Use context API](https://github.com/beyondns/gotips/blob/master/tips32.md#17---use-context-api)
- 16 - [Go routines syncronization](https://github.com/beyondns/gotips/blob/master/tips32.md#16---go-routines-syncronization)
- 15 - [Time interval measurement](https://github.com/beyondns/gotips/blob/master/tips32.md#15---time-interval-measurement)
- 14 - [Benchmark switch vs else if](https://github.com/beyondns/gotips/blob/master/tips32.md#14---benchmark-switch-vs-else-if)
- 13 - [Use ASM in Go Code](https://github.com/beyondns/gotips/blob/master/tips32.md#13---use-asm-in-go-code)
- 12 - [JSON with unknown structure](https://github.com/beyondns/gotips/blob/master/tips32.md#12---json-with-unknown-structure)
- 11 - [Websocket over HTTP2](https://github.com/beyondns/gotips/blob/master/tips32.md#11---websocket-over-http2)
- 10 - [HTTP2](https://github.com/beyondns/gotips/blob/master/tips32.md#10---http2)
-  9 - [Error handling](https://github.com/beyondns/gotips/blob/master/tips32.md#9---error-handling)
-  8 - [Memory management with pools of objects](https://github.com/beyondns/gotips/blob/master/tips32.md#8---memory-management-with-pools-of-objects)
-  7 - [Sort slice of time periods](https://github.com/beyondns/gotips/blob/master/tips32.md#7---sort-slice-of-time-periods)
-  6 - [Fast http server](https://github.com/beyondns/gotips/blob/master/tips32.md#6---fast-http-server)
-  5 - [Close channel to notify many](https://github.com/beyondns/gotips/blob/master/tips32.md#5---close-channel-to-notify-many)
-  4 - [Is channel closed?](https://github.com/beyondns/gotips/blob/master/tips32.md#4---is-channel-closed)
-  3 - [Http request/response with close notify and timeout](https://github.com/beyondns/gotips/blob/master/tips32.md#3---http-requestresponse-with-close-notify-and-timeout)
-  2 - [Import packages](https://github.com/beyondns/gotips/blob/master/tips32.md#2---import-packages)
-  1 - [Map](https://github.com/beyondns/gotips/blob/master/tips32.md#1---map)
-  0 - [Slices](https://github.com/beyondns/gotips/blob/master/tips32.md#0---slices)

## #58 - set go env
> 2016-31-08 by [@beyondns](https://github.com/beyondns)

```bash
export GOROOT=/opt/go1.7.1
export PATH=$GOROOT/bin:$PATH
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

```

## #57 - Multiple concurrent http requests with timeout
> 2016-16-09 by [@beyondns](https://github.com/beyondns)

```go

package main

import (
    _"encoding/json"
    "flag"
    "io/ioutil"
    "log"
    "net/http"
    _"sort"
    "sync/atomic"
    "time"
    "strings"
)

const (
    MaxRequestWaitTime = time.Millisecond * 15000
)

func main() {
    addr := flag.String("addr", ":8080", "http listen address")
    flag.Parse()

    http.HandleFunc("/do", doHandler)

    log.Printf("listen and serve %s", *addr)
    log.Fatal(http.ListenAndServe(*addr, nil))
}



func doHandler(w http.ResponseWriter, r *http.Request) {
    urls, ok := r.URL.Query()["u"]
    if !ok {
        http.Error(w, "u is not specified", http.StatusBadRequest)
        return
    }

    results:=make([][]byte,len(urls))

    var wgc uint32 = 0
    var wgn uint32 = uint32(len(urls))

    reqDoneFlags := make([]bool, len(urls))
    wgCanceled := false

    done := make(chan struct{})
    cancel := make(chan struct{})

    for i, u := range urls {
        go func(u string, i int) {

            defer func() {
                atomic.AddUint32(&wgc, 1)
                if wgc == wgn && !wgCanceled {
                    done <- struct{}{}
                }
            }()

            req, err := http.NewRequest("GET", u, nil)
            if err != nil {
                log.Printf("http.NewRequest %v error %v", u, err)
                return
            }
            req.Header.Add("Content-Type", "application/json")

            tr := &http.Transport{}
            client := &http.Client{Transport: tr}
            clientDone := false

            exit := make(chan struct{})
            defer close(exit)
            go func() {
                select {
                case <-cancel:
                    if !clientDone {
                        tr.CancelRequest(req)
                    }
                case <-exit:
                }
            }()

            resp, err := client.Do(req)
            clientDone = true
            if err != nil {
                log.Printf("http get %v error %v or timeout", u, err)
                return
            }
            defer func() {
                if resp!=nil && resp.Body != nil {
                    resp.Body.Close()
                }
            }()
            if resp.StatusCode != 200 {
                log.Printf("http get %v error status %v", u, resp.StatusCode)
                return
            }

            results[i], err = ioutil.ReadAll(resp.Body)
            if err != nil {
                log.Printf("url %s body error %v", u, err)
                return
            }
            reqDoneFlags[i] = !wgCanceled
            //log.Printf("%s %v", u,results[i])
            if reqDoneFlags[i]{
                log.Printf("%s ok", u)
            }
        }(u, i)
    }

    select {
    case <-time.After(MaxRequestWaitTime):
        wgCanceled = true
        close(cancel)
        //log.Printf("wg timeout")
    case <-done:
    }

    log.Printf("results: %v", results)

    var allResults []byte
    for i, r := range results {
        if reqDoneFlags[i] {
            allResults=append(allResults,r...)
        }
    }

    doneUrls:=""
    for i, u := range urls {
        if reqDoneFlags[i]{
            if len(doneUrls)>0{
                doneUrls+=","
            }
            doneUrls+=u[strings.LastIndex(u,"/")+1:]
        }
    }
    w.Header().Set("Done-Urls", doneUrls)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(allResults)

}

```


## #56 - Inplace struct
> 2016-15-09 by [@beyondns](https://github.com/beyondns)

```go

package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	name:="Hero"
	bin,err:=json.Marshal(&struct {Name string}{Name: name})
	fmt.Printf("%s %v",string(bin),err)
}

```

## #55 - Dynamic time intervals
> 2016-31-08 by [@beyondns](https://github.com/beyondns)


```go

package main

import (
	"log"
	"time"
)

func main() {

	const WorkInterval = time.Second

	counter := 0
	dynoInterval := WorkInterval
	for {
		select {
		case <-time.After(dynoInterval):
			counter++
			log.Printf("%d--enter %v", counter,dynoInterval)
			c, err := doSomeWork(counter)
			log.Printf("%d--exit %v",counter,dynoInterval)

			if err != nil {
				dynoInterval *= 2
				continue
			}
			if c == 0 {
				dynoInterval *= 2
			} else {
				dynoInterval = WorkInterval
			}
		}
	}

}

func doSomeWork(c int) (int, error) {
	if c%3==0{
		time.Sleep(2 * time.Second)
	return c,nil
	}
	return 0, nil
}

```
```bash
2016/08/31 11:58:13 1--enter 1s
2016/08/31 11:58:13 1--exit 1s
2016/08/31 11:58:15 2--enter 2s
2016/08/31 11:58:15 2--exit 2s
2016/08/31 11:58:19 3--enter 4s
2016/08/31 11:58:21 3--exit 4s
2016/08/31 11:58:22 4--enter 1s
2016/08/31 11:58:22 4--exit 1s
2016/08/31 11:58:24 5--enter 2s
2016/08/31 11:58:24 5--exit 2s
```



## #54 - NoDB
> 2016-18-08 by [@beyondns](https://github.com/beyondns)  

A lot of talks which db to use, sql, nosql, blah.  
NoDB = not only database, simple data structures are best friends!  

Go slice and map are awesome. Can be used together, map for fast random access, slice for ordered access.

* [Go Data Structures](https://github.com/emirpasic/gods)

Ok, embedded key/value storages with file persistence 

* [leveldb](https://github.com/syndtr/goleveldb)
* [boltdb](https://github.com/boltdb/bolt)
* [rocksdb](https://github.com/tecbot/gorocksdb)
* [buntdb](https://github.com/tidwall/buntdb)

* [Writing a Simple Persistent Key-Value Store in Go](https://www.opsdash.com/blog/persistent-key-value-store-golang.html)

## #53 - EC operations from bitcoin core
> 2016-28-07 by [@beyondns](https://github.com/beyondns)  

Use secp256k1 (C library for EC operations on curve secp256k1) in Go  

```bash
sudo apt-get install autoconf libtool
git clone https://github.com/bitcoin-core/secp256k1.git
./autogen.sh && ./configure && make
sudo make install
export LD_LIBRARY_PATH=/usr/local/lib

```

```go
/*
#include "secp256k1.h"

#cgo LDFLAGS: -lsecp256k1
*/
import "C"
import "unsafe"

import (
	crand "crypto/rand"
	"io"
	"log"
)

var (
	context *C.secp256k1_context
)

func init() {
	context = C.secp256k1_context_create(C.SECP256K1_CONTEXT_VERIFY | C.SECP256K1_CONTEXT_SIGN)
}


func getRandNum(n int) []byte {
	buf := make([]byte, n)
	_, err := io.ReadFull(crand.Reader, buf)
	if err != nil {
		log.Fatalf("io.ReadFull(crand.Reader, buf) %v",err)
	}
	return buf
}



func GenKeyPair() ([]byte, []byte, bool) {
	var priv []byte = make([]byte,32)
	_, err := io.ReadFull(crand.Reader, priv)
	if err != nil {
		log.Fatalf("io.ReadFull(crand.Reader, buf) %v",err)
	}
	var pub []byte = make([]byte, 64)
	ret := C.secp256k1_ec_pubkey_create(context,
		(*C.secp256k1_pubkey)(unsafe.Pointer(&pub[0])),
		(*C.uchar)(unsafe.Pointer(&priv[0])),
	)

	if ret != C.int(1) {
		return nil,nil,false
	}

	return pub, priv, true
}



func PrivkeyVerify(priv []byte) bool {
	return C.secp256k1_ec_seckey_verify(context,
		(*C.uchar)(unsafe.Pointer(&priv[0]))) == C.int(1)
}


func GenPrivkey() ([]byte, bool) {
	const privkeylen=32
	privkey := make([]byte, privkeylen)
	var r bool
	const attlim = 1024
	for i:=0;i<attlim;i++{
		io.ReadFull(crand.Reader, privkey)
		r=PrivkeyVerify(privkey)
		if !r{
			continue
		}
		break
	}

	if !r {
		return nil,r
	}

	return privkey, true
}

func GenPubkey(priv []byte) ([]byte, bool) {
	var pub []byte = make([]byte, 64)
	ret := C.secp256k1_ec_pubkey_create(context,
		(*C.secp256k1_pubkey)(unsafe.Pointer(&pub[0])),
		(*C.uchar)(unsafe.Pointer(&priv[0])),
	)
	if ret != C.int(1) {
		return nil,false
	}
	return pub,true
}

func PubkeySerialize(pub []byte) ([]byte,bool) {
	var pubComp []byte = make([]byte, 33)
	var output_len C.size_t = C.size_t(33) 
	ret := C.secp256k1_ec_pubkey_serialize(
		context,
		(*C.uchar)(unsafe.Pointer(&pubComp[0])),
		&output_len,
		(*C.secp256k1_pubkey)(unsafe.Pointer(&pub[0])),
		C.SECP256K1_EC_COMPRESSED,
	)
	if ret != C.int(1) {
		return nil,false
	}
	return pubComp,true
}


func PubkeyParse(in []byte) ([]byte, bool) {
	var pub []byte = make([]byte, 64)
	var in_len C.size_t = C.size_t(33) 
	ret := C.secp256k1_ec_pubkey_parse(
		context,
		(*C.secp256k1_pubkey)(unsafe.Pointer(&pub[0])),
		(*C.uchar)(unsafe.Pointer(&in[0])),
		in_len,
	)
	if ret != C.int(1) {
		return nil,false
	}
	return pub,true
}


func MsgSign(msg []byte, priv []byte, nonce []byte) ([]byte, bool) {
	const siglen = 64
	sig:=make([]byte,siglen)	

	ret :=C.secp256k1_ecdsa_sign(
		context,
		(*C.secp256k1_ecdsa_signature)(unsafe.Pointer(&sig[0])), 
		(*C.uchar)(unsafe.Pointer(&msg[0])), 
		(*C.uchar)(unsafe.Pointer(&priv[0])), 
		&(*C.secp256k1_nonce_function_default), 
		unsafe.Pointer(&nonce[0]),
	)

	if ret != C.int(1) {
		return nil,false
	}
	return sig,true
}


func MsgSignVerify(msg []byte, sig []byte, pub []byte) (bool) {
	ret :=C.secp256k1_ecdsa_verify(
		context,
		(*C.secp256k1_ecdsa_signature)(unsafe.Pointer(&sig[0])), 
		(*C.uchar)(unsafe.Pointer(&msg[0])), 
		(*C.secp256k1_pubkey)(unsafe.Pointer(&pub[0])),
	)

	if ret != C.int(1) {
		return false
	}
	return true
}

```

```go
package secp256k1_test

import (
	_"crypto/rand"
	_"crypto/sha256"
	_"io"
	"testing"
	"github.com/beyondns/hypers/crypto/secp256k1"
	"fmt"
	"log"
	"bytes"
	"io"
	crand "crypto/rand"

)

func TestGenKeyPair(t *testing.T) {
	var ok bool
	var pub,priv []byte

	for i:=0;i<1000;i++{
		pub,priv,ok=secp256k1.GenKeyPair()
		if !ok {
			log.Printf("secp256k1.GenKeyPair error")
			continue
		}
		r:=secp256k1.PrivkeyVerify(priv)
		if !r {
			log.Printf("secp256k1.PrivkeyVerify error")
			continue
		}
	}

	fmt.Printf("pub %x\n",pub)
	fmt.Printf("priv %x\n",priv)
	fmt.Printf("ok\n")
}

func TestGenKeyPair2(t *testing.T) {
	var ok bool
	var pub,pubComp,pub2,priv []byte

	for i:=0;i<1000;i++{

	priv,ok = secp256k1.GenPrivkey()
	if !ok{
		t.Fatalf("GenPrivkey()")		
	}
	pub,ok = secp256k1.GenPubkey(priv)
	if !ok{
		t.Fatalf("GenPubkey()")		
	}
	pubComp,ok = secp256k1.PubkeySerialize(pub)
	if !ok{
		t.Fatalf("secp256k1.PubkeySerialize()")		
	}

	pub2,ok = secp256k1.PubkeyParse(pubComp)
	if !ok{
		t.Fatalf("secp256k1.PubkeyParse()")		
	}

	if !bytes.Equal(pub,pub2){
		t.Fatalf("bytes.Equal(pub,pub2)")
	}

	} //

	fmt.Printf("priv %x\n",priv)
	fmt.Printf("pub %x\n",pub)
	fmt.Printf("pubComp %x\n",pubComp)
	fmt.Printf("pub2 %x\n",pub2)
 	fmt.Printf("ok\n")
}


func TestGenKeyPairSignVerify(t *testing.T) {
	var ok bool
	var pub,priv,sig,msg []byte
	pub,priv,ok=secp256k1.GenKeyPair()
	msg=make([]byte,32)
	io.ReadFull(crand.Reader, msg)
	nonce:=make([]byte,32)
	io.ReadFull(crand.Reader, nonce[:])

	sig,ok=secp256k1.MsgSign(msg,priv,nonce)
	if !ok{
		t.Fatalf("secp256k1.MsgSign()")		
	}

	ok=secp256k1.MsgSignVerify(msg,sig,pub)
	if !ok{
		t.Fatalf("secp256k1.MsgSignVerify")		
	}

	sig[1]=sig[1]^1

	ok=secp256k1.MsgSignVerify(msg,sig,pub)
	if ok{
		t.Fatalf("secp256k1.MsgSignVerify")		
	}

	fmt.Printf("priv %x\n",priv)
	fmt.Printf("pub %x\n",pub)
	fmt.Printf("sig %x\n",sig)
 	fmt.Printf("ok %t\n",ok)

}

```

## #52 - nil
> 2016-22-07 by [@beyondns](https://github.com/beyondns)  

nil is zero value for pointer, channel, func, interface, map, slice.

## #51 - BBQ
> 2016-19-07 by [@beyondns](https://github.com/beyondns)  

Block buffered queue (BBQ)

```go
package bbq

import (
	"errors"
	"sync"
)

// block buffered queue
type BlockBQ struct {
	B1, B1 *[]interface{} // 2 buffers pointers
	I      int            // current index
	M      int            // max index
	T      int64          // total number
	mx     sync.Mutex     // sync
}

func NewBBQ(l int) *BlockBQ {
	b1, b2 := make([]interface{}, l), make([]interface{}, l)
	bbq := &BlockBQ{B1: &b1, B2: &b2}
	return bbq
}

func (bbq *BlockBQ) Swap() int {
	var l int
	bbq.mx.Lock()
	t := bbq.B1
	bbq.B1 = bbq.B2
	bbq.B2 = t
	l = bbq.I
	bbq.I = 0
	bbq.T += int64(l)
	bbq.mx.Unlock()
	return l
}

func (bbq *BlockBQ) Push(p interface{}) error {
	var err error
	bbq.mx.Lock()
	if bbq.I >= len(*(bbq.B1)) {
		err = errors.New("BBQ overflow")
	} else {
		(*bbq.B1)[bbq.I] = p
		bbq.I++
	}
	if bbq.I > bbq.M {
		bbq.M = bbq.I
	}
	bbq.mx.Unlock()
	return err
}
```


```go
package bbq

import (
	_ "fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestBBQ(t *testing.T) {
	var data []int
	bbq := NewBBQ(1024)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 1000):
				l := bbq.Swap()
				bk := make([]interface{}, l)
				copy(bk, *(bbq.B2))
				go func(bk []interface{}) {
					for _, v := range bk {
						d := v.(int)
						data = append(data, d)
					}
					log.Printf("bbq swap len(bk)=%d, len(data)=%d", len(bk), len(data))
				}(bk)
			case <-done:
				log.Printf("Done 1")
				return
			}
		}
	}()

	go func() { // generate txs
		counter := 0
		for {
			select {
			case <-time.After(time.Millisecond * 100):
				n := 1 + rand.Intn(10)
				for i := 0; i < n; i++ {
					bbq.Push(counter)
					counter++
				}
			case <-done:
				log.Printf("Done 2")
				return
			}

		}
	}()

	<-time.After(time.Second * 5)
	close(done)
	log.Printf("len(data)=%d", len(data))
	log.Printf("%v", data)
	for i, d := range data {
		if i != d {
			t.Fatalf("i!=d %d %d", i, d)
		}
	}
}
```

## #50 - cool go links
> 2016-05-07 by [@beyondns](https://github.com/beyondns)  

* [channels](https://github.com/eapache/channels)
* [jsonparser](https://github.com/buger/jsonparser)
* [go data structures](https://github.com/emirpasic/gods)

## #49 - custom wait group
> 2016-15-06 by [@beyondns](https://github.com/beyondns)  

Custom wait group, it's easy to use wait groups to sync different go routines, if you need more control just do it custom.   

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Start")
	const N = 5
	var c uint32 = 0
	var n uint32 = N
	done := make(chan struct{})
	doneMap := make(map[int]struct{})

	for i := 0; i < N; i++ {
		go func(i int) {
			// do work, an http request for example
			if i == 3 {
				// simulate timeout
				time.Sleep(2 * time.Second)
			}
			doneMap[i] = struct{}{}
			fmt.Printf("done f%d\n", i)
			atomic.AddUint32(&c, 1)
			if c == n {
				done <- struct{}{}
			}
		}(i)
	}

	select {
	case <-time.After(time.Second * 1):
		for i := 0; i < N; i++ {
			_, ok := doneMap[i]
			if ok {
				fmt.Printf("%d done ok\n", i)
			} else {
				fmt.Printf("%d not done\n", i)
			}

		}
		panic("time out")
	case <-done:
		fmt.Printf("Done")
	}

}

```



## #48 - ev_loop vs co_routine
> 2016-07-06 by [@beyondns](https://github.com/beyondns)  

This tip was promised, stay tuned.

It's super simple to implement udp/tcp/http etc. servers in #Go but how it vs plain event loop in C.  

ev_loop.c
```c
```

co_routine.go
```go
```

## #47 - Use C for computation intensive operations
> 2016-29-05 by [@beyondns](https://github.com/beyondns)  

Usage of C in Go code pretty stright forward  

```go
package main

/*
#include <stdlib.h>
#include <openssl/sha.h>    
#cgo LDFLAGS: -lcrypto
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {

    const msg string = "Alice in wonderland"

    md := (*C.char)(unsafe.Pointer(C.SHA256( (*C.uchar)(unsafe.Pointer(C.CString(msg))), 
        C.size_t(len(msg)), (*C.uchar)(nil) )))

    res := C.GoString( md );

    fmt.Printf("%x",res)// 604e40d7814621e2412a42b8d04e37d56bdea4628f5d622fb782bb644be8722f
}
```

* [cgo](https://golang.org/cmd/cgo/)
* [SHA256](https://rosettacode.org/wiki/SHA-256#C)

## #46 - Unique key doc manipulation in MonGo 
> 2016-21-05 by [@beyondns](https://github.com/beyondns)  

Every mongo record has unique autogenerated Id, but if u need external unique human readable key, just use {_k:'hero'} and following wrapper methods:  

```go
package xid

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoServers = "localhost" //"server1.example.com,server2.example.com"
	mongoDBName  = "xid"
)

var (
	mongoSess mgo.Session
	mongoDB   *mgo.Database

	TestC, UserC, EventC *mgo.Collection
)

func init() {
	mongoSess, err := mgo.Dial(mongoServers)
	if err != nil {
		panic(err)
	}
	mongoSess.SetMode(mgo.Monotonic, true)
	mongoDB = mongoSess.DB(mongoDBName)

	TestC,UserC,EventC = mongoDB.C("test"),mongoDB.C("users"),mongoDB.C("events")
}

func SetObject(c *mgo.Collection, key string, data bson.M) error {
	_, err := c.Upsert(bson.M{"_k": key}, bson.M{"$set": data})
	return err
}

func GetObject(c *mgo.Collection, key string) (bson.M, error) {
	var res bson.M
	err := c.Find(bson.M{"_k": key}).One(&res)
	return res, err
}

func GetObjectSelect(c *mgo.Collection, key string,sel bson.M) (bson.M, error) {
	var res bson.M
	err := c.Find(bson.M{"_k": key}).Select(sel).One(&res)
	return res, err
}

func DelObject(c *mgo.Collection, key string) error {
	return c.Remove(bson.M{"_k": key})
}
```

```js
// little tip how to select sub node key
// data
{"_key":"key1",{"node":{"subnode":"subnodevalue"}}}
// query
{"_key" : "key1","node.subnode":{ $exists: true }}
```

```js
// little tip how to select by date, a bit js code here
// query
{'date':{'$gte' : new Date((new Date()).toISOString())}}
```

* [mgo.v2](https://godoc.org/gopkg.in/mgo.v2)

## #45 - Go is more a framework than a language
> 2016-08-05 by [@beyondns](https://github.com/beyondns)  

Go is more a framework than a language. Anything can be built on top of [libuv](https://github.com/libuv/libuv), [libev](http://software.schmorp.de/pkg/libev.html), [epoll](https://en.wikipedia.org/wiki/Epoll), etc. But it's not so easy.  
Many ideas have been accumulated in Go + awesome tools and libraries. That it. No magic. Ready to go in production.

* ["Go style concurency in C"](https://github.com/sustrik/libmill)

## #44 - Facebook messenger chatbot
> 2016-03-05 by [@beyondns](https://github.com/beyondns)  

Simple echo bot for facebook messenger

```go
package main

//
// https://developers.facebook.com/docs/messenger-platform/send-api-reference#guidelines
//

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"net/url"
	"encoding/json"
	"time"
	"bytes"
	"errors"
	"regexp"
)

const (
    API   = "https://graph.facebook.com/v2.6/me/messages"
	TOKEN = < YOUR TOKEN HERE>
)


type FBId string

type FBSender struct{
	Id FBId `json:"id"`
}

type FBRecipient struct{
	Id FBId `json:"id"`
}

type FBElementPostbackButton struct{
	Type string `json:"type"` // postback
	Title string `json:"title"`
	Payload string `json:"payload"`
}

type FBElementUrlButton struct{
	Type string `json:"type"` // web_url
	Title string `json:"title"`
	Url string `json:"url"`
}

type FBTemplatePayloadElement struct{
	Title string `json:"title"`
	Image_url string `json:"image_url"`
	Subtitle string `json:"subtitle"`
	Buttons []interface{} `json:"buttons"`
}

type FBTemplatePayload struct{
	Template_type string `json:"template_type"`
	Elements []FBTemplatePayloadElement `json:"elements"`
}


type FBImgPayload struct{
	Url string `json:"url"`
}

type FBAttachment struct{
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
}


type FBSendAttachmentMessage struct{
	Attachment FBAttachment `json:"attachment"`
}


type FBSendAttachementMsgEntry struct{
	Recipient FBRecipient `json:"recipient"`
    Message FBSendAttachmentMessage `json:"message"`
}


type FBSendTextMessage struct{
	Text string `json:"text"`
}

type FBSendTextMsgEntry struct{
	Recipient FBRecipient `json:"recipient"`
    Message FBSendTextMessage 	`json:"message"`
}

type FBMessagePostbackPayload struct{
	Payload string `json:"payload"`
}
type FBMessage struct{
	Text string `json:"text"`

}

type FBMessaging struct{
	Sender FBSender `json:"sender"`
	Recipient FBRecipient `json:"recipient"`
	TS int64 `json:"timestamp"`
	Message FBMessage `json:"message"`
	Postback FBMessagePostbackPayload `json:"postback"`
}

type FBEntry struct{
    M []FBMessaging `json:"messaging"`
}

type FBMsgEntry struct{
	E []FBEntry `json:"entry"`
}


type Forever struct{}

func (s *Forever) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	if r.Method == "GET" {
		if r.URL.Path == "/webhook" {
			if r.URL.Query().Get("hub.verify_token") == "verify_me_token1" {
				fmt.Fprintf(w, "%s", r.URL.Query().Get("hub.challenge"))
				return
			}
			fmt.Fprintf(w, "hi")
			return
		}
	}

	if r.Method == "POST" {
		if r.URL.Path == "/webhook" {

			d, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("bad reguest, error %v", err), http.StatusBadRequest)
				return
			}

			log.Printf("%s", string(d))

			fbe := FBMsgEntry{}
			if len(d) > 0 {
				err := json.Unmarshal(d, &fbe)
				if err != nil {
					log.Fatal(err)
				}
				for _,e:=range fbe.E{
					for _,m:=range e.M{
						MessageHandle(&m)
					}
				}
			}
			return
		}
	}

}


func MessageHandle(m *FBMessaging){

	if len(m.Postback.Payload)>0{
		log.Printf("Postback.Payload %d,%s",m.Sender,m.Postback.Payload)
		sendTextMessage(m.Sender.Id,"you awesome man!")
	}
	if len(m.Message.Text)>0{

		if m.Message.Text == "demo"{
			st,d,err:=sendTemplateMessage4(m.Sender.Id,[]string{
			"http://media1.santabanta.com/full1/Animals/Elephants/elephants-9a.jpg",
			"https://upload.wikimedia.org/wikipedia/commons/a/ad/Kruger_Elephant.JPG",
			"https://upload.wikimedia.org/wikipedia/commons/4/4f/Elephant%2C_Selous_Game_Reserve.jpg",
			})
			log.Printf("%d %s %v",st,string(d),err)
			return
		}

		log.Printf("Message.Text %d,%s",m.Sender,m.Message.Text)
		sendTextMessage(m.Sender.Id,"type demo for demo")//m.Message.Text)
	}
}

func main() {
	port := os.Getenv("OPENSHIFT_GO_PORT")
	host := os.Getenv("OPENSHIFT_GO_IP")
	if port == "" {
		port = "8080"
	}
	if host == "" {
		host = "127.0.0.1"
	}
	bind := fmt.Sprintf("%s:%s", host, port)

	log.Printf("listening on %s...", bind)
	log.Fatal(http.ListenAndServe(bind, &Forever{}))
}

func sendTemplateMessage4(senderId FBId, carouselUrls []string) (int, []byte, error){

    fbe:=&FBSendAttachementMsgEntry{}
    fbe.Recipient.Id=senderId
    fbe.Message.Attachment.Type="template"
    payload:=FBTemplatePayload{
    	Template_type:"generic",
    }
    for i,img_url:=range carouselUrls{
    	e:=FBTemplatePayloadElement{
			Title:fmt.Sprintf("title %d",i),
			Image_url: img_url,
			Subtitle:fmt.Sprintf("subtitle %d",i),
	   	}
    	e.Buttons = append(e.Buttons, FBElementUrlButton{
            Type:"web_url",
            Url: img_url,
            Title: "View Item",
    	})
    	e.Buttons = append(e.Buttons, FBElementPostbackButton{
            Type:"postback",
            Title: "Action",
            Payload: fmt.Sprintf("payload %d",i),
    	})
    	payload.Elements=append(payload.Elements,e)
    }

    fbe.Message.Attachment.Payload=payload

    data,err:=json.Marshal(*fbe)
    if err!=nil{
    	log.Fatal(err)
    }

    return sendMessage(senderId,data)
}


func sendImgMessage(senderId FBId,imgUrl string) (int, []byte, error){
    fbe:=&FBSendAttachementMsgEntry{}
    fbe.Recipient.Id=senderId
    fbe.Message.Attachment.Type="image"
    fbe.Message.Attachment.Payload=FBImgPayload{Url:imgUrl}
    data,err:=json.Marshal(*fbe)
    if err!=nil{
    	log.Fatal(err)
    }
    return sendMessage(senderId,data)
}

func sendTextMessage(senderId FBId,text string) (int, []byte, error){
    fbe:=&FBSendTextMsgEntry{}
    fbe.Recipient.Id=senderId
    fbe.Message.Text=text
    data,err:=json.Marshal(*fbe)
    if err!=nil{
    	log.Fatal(err)
    }
    return sendMessage(senderId,data)
}

func sendMessage(senderId FBId,data []byte) (int, []byte, error){
	v := url.Values{}
    v.Add("access_token", TOKEN)
    log.Printf("sendMessage to:%v data: %s",senderId, string(data))
	st,d,err:=httpRequest("POST", API+"?"+v.Encode(),[]byte(data),time.Second*15)	
	return st,d,err
}


func httpRequest(meth, u string, data []byte,
    timeLimit time.Duration) (int, []byte, error) {

    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    c := make(chan error, 1)

    var respStatus int
    var respBody []byte
    req, err := http.NewRequest(meth, u, bytes.NewBuffer(data))
    if err != nil {
        return 0, nil, err
    }

    req.Header.Add("Content-Type", "application/json; charset=utf-8")

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

```

* [messenger](https://developers.facebook.com/products/messenger/)

## #43 - Passing args as a map
> 2016-01-05 by [@beyondns](https://github.com/beyondns)  

In some cases passign arguments as map to a function is flexible and convenient way. It looks like js hack. 


```go
package main

import(
	//"fmt"
	"log"
)

func f2(args map[string]interface{})int{
	if _,ok:=args["y"];ok{
		log.Printf("y in args: %v",args["y"])
		return 1
	}
	return 0	
}
func f1(args map[string]interface{})int{
	if _,ok:=args["x"];ok{
		log.Printf("x in args: %v",args["x"])
	}
	args["y"]="zebro"
	return f2(args)
}

func main(){
	args:=make(map[string]interface{})
	args["x"]=12345
	log.Printf("result=%d",f1(args))

}
```


## #42 - Optimize Go 
> 2016-26-04 by [@beyondns](https://github.com/beyondns)  

[How to optimize ‪#‎golang‬ for really high performance](https://www.youtube.com/watch?v=ZuQcbqYK0BY)

## #41 - Telegram bot 
> 2016-06-04 by [@beyondns](https://github.com/beyondns)  

Simple echo Telegram bot no external dependencies

```go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// https://core.telegram.org/bots/api

const (
	API   = "https://api.telegram.org/bot"
	TOKEN = <BOT TOKEN>
)

type Chat struct {
	Id int `json:"id"`
}

type From struct {
	Id int `json:"id"`
}

type Message struct {
	Id   int    `json:"message_id"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
	Date uint32 `json:"date"`
}

type Update struct {
	Id  int     `json:"update_id"`
	Msg Message `json:"message"`
}

type Result struct {
	Ok  bool     `json:"ok"`
	Res []Update `json:"result"`
}

func main() {
	log.Printf("Me:%s", Me())
	uc := createUpdatesChan()
	for u := range uc {

		if u.Msg.Text == "yes" {
			SendMessage(u.Msg.Chat.Id, "You awesome, yes!")
			continue
		}
		if u.Msg.Text == "no" {
			SendMessage(u.Msg.Chat.Id, "No newer say no, say yes!")
			continue
		}

		SendMessageMarkup(u.Msg.Chat.Id,
			"Just answer",
			`{"one_time_keyboard":true,"resize_keyboard":true,"keyboard":[["yes","no"]]}`)
	}
}

func createUpdatesChan() <-chan *Update {
	uc := make(chan *Update, 1024)
	go func() {

		const limit = 16
		const timeout = 60
		var offset = 0
		for {
			res, err := Updates(offset, limit, timeout)
			if err != nil {
				log.Printf("Updates error %v", err)
				time.Sleep(time.Second * 3)
				continue
			}

			if !res.Ok {
				log.Printf("Updates !res.Ok %v", res)
				time.Sleep(time.Second * 3)
				continue
			}

			for i, u := range res.Res {
				if u.Id >= offset {
					log.Printf("u %v", u)
					offset = u.Id + 1
					uc <- &res.Res[i]
				}
			}
			//            time.Sleep(time.Millisecond * 10)
		}
	}()

	return uc
}

func Me() string {
	st, d, err := httpRequest("GET", API+TOKEN+"/getMe", nil, time.Second*15)
	if err != nil {
		log.Fatal(err)
	}
	if st != http.StatusOK {
		log.Fatalf("status %d", st)
	}
	return string(d)
}

func Updates(offset, limit, timeout int) (*Result, error) {
	v := url.Values{}
	if offset > 0 {
		v.Add("offset", strconv.Itoa(offset))
	}
	if limit > 0 {
		v.Add("limit", strconv.Itoa(limit))
	}
	if timeout > 0 {
		v.Add("timeout", strconv.Itoa(timeout))
	}
	st, d, err := httpRequest("GET", API+TOKEN+"/getUpdates?"+v.Encode(), nil,
		time.Second*(60+15))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if st != http.StatusOK {
		log.Printf("status %d %s\n", st, string(d))
		return nil, err
	}

	res := &Result{}
	if len(d) > 0 {
		err := json.Unmarshal(d, &res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return res, nil
}

func SendMessage(chat_id int, text string) (string, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(chat_id))
	v.Add("text", text)
	st, d, err := httpRequest("GET", API+TOKEN+"/sendMessage?"+v.Encode(),
		nil, time.Second*15)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if st != http.StatusOK {
		log.Printf("status %d %s\n", st, string(d))
		return "", err
	}
	return string(d), nil
}

func SendMessageMarkup(chat_id int, text string, reply_markup string) (string, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(chat_id))
	v.Add("text", text)
	v.Add("reply_markup", reply_markup)
	st, d, err := httpRequest("GET", API+TOKEN+"/sendMessage?"+v.Encode(),
		nil, time.Second*15)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if st != http.StatusOK {
		log.Printf("status %d %s\n", st, string(d))
		return "", err
	}
	return string(d), nil
}

func httpRequest(meth, u string, data []byte,
	timeLimit time.Duration) (int, []byte, error) {

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)

	var respStatus int
	var respBody []byte
	req, err := http.NewRequest(meth, u, bytes.NewBuffer(data))
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
```

* [api](https://core.telegram.org/bots/api)


## #40 - Distributed consensus 
> 2016-01-04 by [@beyondns](https://github.com/beyondns)  

Distributed consensus is a common shared state or logic support equal across multiple nodes.  
1-stage (semi or one-many) consensus: send data to nodes, get responses, find consensus on one node.  
2-stage (full or many-many) consensus: send data to nodes, nodes send all-to-all, find consensus on each node.  

```go
package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"strings"
	"time"
	"bytes"
	"errors"
	"io/ioutil"
	"sync"
)

var (
	Nodes []string
)

func sysHandler1(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("bad reguest, error %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("sysHandler1: %s",string(d))
	w.WriteHeader(http.StatusOK)
}

func sysHandler2(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("bad reguest, error %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("sysHandler2: %s",string(d))
	if localConsensus(d,"1",w,r){
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusOK)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("bad reguest, error %v", err), http.StatusBadRequest)
		return
	}

	log.Printf("apiHandler: %s",string(d))

	if localConsensus(d,"2",w,r){
		w.WriteHeader(http.StatusOK)
	}

}

func localConsensus(d []byte, st string, 
	w http.ResponseWriter, r *http.Request) bool{
	if len(d)>0{
		ok,err:=SendDataToWorld(d,st)
		if err!=nil {
			http.Error(w, fmt.Sprintf("SendDataToWorld error %v", err), http.StatusBadRequest)
			return false
		}
		if !ok {
			http.Error(w, "no consensus", http.StatusBadRequest)
			return false
		}
	}
	return true
}

func main() {
	nodes := flag.String("nodes", "127.0.0.1:12379,127.0.0.1:22379,127.0.0.1:32379", "nodes host0:port0,host1:port1,..")
	api := flag.String("api", "127.0.0.1:12379", "api host:port")
	flag.Parse()

	Nodes=strings.Split(*nodes,",")

	log.Printf("Nodes %v",Nodes)

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/sys1", sysHandler1)
	http.HandleFunc("/sys2", sysHandler2)
	log.Fatal(http.ListenAndServe(*api, nil))
}

func SendDataToWorld(data []byte, st string) (bool,error){

	var l = len(Nodes)
	var wg sync.WaitGroup
	wg.Add(l)

	var rd [][]byte = make([][]byte, l)
	var rs []int = make([]int, l)
	var re []bool = make([]bool, l)
	for i, u := range Nodes {
		go func(i int, u string) {
			var e error
			rs[i],rd[i],e=httpRequest("POST",
				"http://"+u+"/sys"+st,data,time.Second*15)
			re[i]=(e==nil)
			wg.Done()
		}(i, u)
	}
	wg.Wait()

	log.Printf("%v %v %v",rs,rd,re)

	return findConsensus(rs,rd,re)

}

func httpRequest(meth, u string, data []byte, 
	timeLimit time.Duration) (int, []byte, error) {

    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    c := make(chan error, 1)

    var respStatus int
    var respBody []byte
    req, err := http.NewRequest(meth, u, bytes.NewBuffer(data))
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

func findConsensus(rs []int,rd [][]byte,re []bool)(bool,error){
	if len(rs) != len(rd) || len(rs) == 0 || len(rd) == 0 {
		panic(fmt.Sprintf("FindConsensus error papams: %v,%v", rs, rd))
	}
	ok:=0
	n:=0
	for i, _ := range rs{
		if !re[i]{
			continue	
		}
		n++
		if rs[i] == http.StatusOK{
			ok++
		}
	}
	log.Printf("find consensus %d : %d",ok,n)
	// 51% ok
	if float32(ok)>float32(n)/float32(2){
		return true,nil
	}
	return false,nil
}
```

## #39 - Public private struct fields and funcs 
> 2016-26-03 by [@beyondns](https://github.com/beyondns)  

Go has specific approach of definition private/public package fields and funcs. All declarations within package is visible (different os files).  
But from other packages only names with 1st capital letter is visible (public)  

lib/l.go

```go
package mylib

import (
	"fmt"
)

const (
	Const1 = "c1"
	const2 = "c2"
)

var (
	Var1 = 1
	var2 = 2
)

type Mydata struct{
	X int // Public
	y int // private
}

func (md *Mydata) DoAPI(){
	fmt.Println("DoAPI")
}

func (md *Mydata) doAPI(){
	fmt.Println("doAPI")
}

```

main.go

```go
package main

import(
	"./lib"
	"fmt"
)

func main(){
	x:=mylib.Mydata{
		X:1,
		//y:2, // error private
	}

	x.DoAPI()

//	x.doAPI() // error private

	fmt.Println(x.X)

//	fmt.Print(x.y) // error private

	fmt.Println(mylib.Const1)
//	fmt.Println(mylib.const2) // error private

	fmt.Println(mylib.Var1)
//	fmt.Println(mylib.var2) // error private

}
```

## #38 - Marshal Unmarshal to byte slice usng gob 
> 2016-25-03 by [@beyondns](https://github.com/beyondns)  

```go
import (
	"bytes"
	"encoding/gob"
)

func gobMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gobUnmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(v)
}

```

## #37 - Time series on leveldb 
> 2016-22-03 by [@beyondns](https://github.com/beyondns)  

Use timestamp as a key in keyvalue storage, leveldb or any other. Just iterate over specific period of time.

```go
import (
	"github.com/syndtr/goleveldb/leveldb"
)
var (
	bcdb *leveldb.DB
)

func init(){
	var err error
	bcdb, err = leveldb.OpenFile("./db/bc.db", nil)
	if err!=nil{
		xlog.Error(err)
		return
	}
	xlog.Notice("db init ok")
}
```
Test

```go
func TestDbTS(t *testing.T) {
	err := bcdb.Put([]byte("t0"),[]byte("boom"), nil)
	if err!=nil{
		t.Error(err)
	}
	
	var timelabel int64
	batch := new(leveldb.Batch)
	for i:=0;i<16;i++{
		if i==8{
			timelabel=time.Now().UnixNano()
		}
		batch.Put([]byte(fmt.Sprintf("t%d",time.Now().UnixNano())), 
			[]byte(fmt.Sprintf("v%d",i)))
	}
	err = bcdb.Write(batch, nil)
	if err!=nil{
		t.Error(err)
	}

    iter := bcdb.NewIterator(nil, nil)
    for ok := iter.Seek([]byte(fmt.Sprintf("t%d",timelabel))); ok; ok = iter.Next() {
    	key,value := iter.Key(),iter.Value()
    	ks:=string(key)
    	if ks[0] == 't'{
			ks=ks[1:]    		
    	}
    	ki, err := strconv.ParseInt(ks, 10, 64)
    	if err!=nil{
    		xlog.Error(err)
    	}
    	xlog.Debugf("key time : %s | value: %s\n", time.Unix(0,ki), value)
	}
	iter.Release()
	err = iter.Error()
	if err!=nil{
		t.Error(err)
	}
}
```

## #36 - Custom type marshal unmarshal json 
> 2016-21-03 by [@beyondns](https://github.com/beyondns)  

Json marshalling can be defined for custom type by implementing MarshalJSON/UnmarshalJSON methods.

```go
type ByteSlice []byte

func (s *ByteSlice) MarshalJSON() ([]byte, error) {
	return []byte(`"`+hex.EncodeToString(*s)+`"`), nil
}

func (s *ByteSlice) UnmarshalJSON(data []byte) error {
	d:=string(data)
	if d[0]=='"'{d=d[1:]}
	l:=len(d)
 	if d[l-1]=='"'{d=d[:l-1]}
	var err error
	*s,err=hex.DecodeString(d)
	return err
}
```

Testing

```go

type ByteSliceHolder struct {
	Bin ByteSlice `json:"bin"`
}

func TestByteSliceJSON(t *testing.T) {
	bsh := &ByteSliceHolder{}
	_, err := fmt.Sscanf("082372739127341723ab", "%x", &bsh.Bin)
	if err != nil {
		t.Fatal("fmt.Sscanf failed")
	}
	bin, err := json.Marshal(bsh)

	if err != nil {
		t.Errorf("json.Marshal failed ", err)
	}
	//xlog.Debugf("%s", bin)

	bsh2 := &ByteSliceHolder{}
	err = json.Unmarshal(bin, &bsh2)

	if err != nil {
		t.Errorf("json.Marshal failed ", err)
	}

	if !bytes.Equal(bsh.Bin,bsh2.Bin){
		t.Errorf("json.Marshal failed ", err)
	
	}

}

```

## #35 - Test specific functions 
> 2016-21-03 by [@beyondns](https://github.com/beyondns)  

To test TestFunc1

```bash
go test -run Func1
```
run argument is a regex, "Func$" test all funcs with "Func" at beginning.

* [Run only specific test cases in golang go test](http://blog.gaku.net/match/)

## #34 - File path exists 
> 2016-15-03 by [@beyondns](https://github.com/beyondns)  

```go
if _, err := os.Stat(filepath); os.IsNotExist(err) {
	// doesn't exist
}

if _, err := os.Stat(filepath); err == nil {
	// exists
}
```

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


[gotips #0..#31](https://github.com/beyondns/gotips/blob/master/tips32.md)


### General Golang links
* [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

### Inspired by
* [jstips](https://github.com/loverajoel/jstips)

### License
[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png)](http://creativecommons.org/publicdomain/zero/1.0/)

