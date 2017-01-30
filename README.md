![header](https://raw.githubusercontent.com/beyondns/gotips/master/gophers.jpg)

> **Golang short tips & tricks**

This list of short golang code tips & tricks will help keep collected knowledge in one place.

> **Support the project**

Send some ether 0x30FD8822D15081F3c98e6A37264F8dF37a2EB416


# Tips list


- 65 - [grpc](https://github.com/beyondns/gotips#65---grpc)
- 64 - [go vs rust](https://github.com/beyondns/gotips#64---go-vs-rust)
- 63 - [chan with timeout](https://github.com/beyondns/gotips/blob/master/tips64.md#63---chan-with-timeout)
- 62 - [contex vs chan](https://github.com/beyondns/gotips/blob/master/tips64.md#62---context-vs-chan)
- 61 - [log with line number and func name](https://github.com/beyondns/gotips/blob/master/tips64.md#61---log-with-line-number-and-func-name)
- 60 - [custom queue](https://github.com/beyondns/gotips/blob/master/tips64.md#60---custom-queue)
- 59 - [go tools usage](https://github.com/beyondns/gotips/blob/master/tips64.md#59---go-tools-usage)
- 58 - [set go env](https://github.com/beyondns/gotips/blob/master/tips64.md#58---set-go-env)
- 57 - [Multiple concurrent http requests with timeout](https://github.com/beyondns/gotips/blob/master/tips64.md#57---multiple-concurrent-http-requests-with-timeout)
- 56 - [Inplace struct](https://github.com/beyondns/gotips/blob/master/tips64.md#56---inplace-struct)
- 55 - [Dynamic time intervals](https://github.com/beyondns/gotips/blob/master/tips64.md#55---dynamic-time-intervals)
- 54 - [NoDB](https://github.com/beyondns/gotips/blob/master/tips64.md#54---nodb)
- 53 - [EC operations from bitcoin core](https://github.com/beyondns/gotips/blob/master/tips64.md#53---ec-operations-from-bitcoin-core)
- 52 - [nil](https://github.com/beyondns/gotips/blob/master/tips64.md#52---nil)
- 51 - [BBQ](https://github.com/beyondns/gotips/blob/master/tips64.md#51---bbq)
- 50 - [cool go links](https://github.com/beyondns/gotips/blob/master/tips64.md#50---cool-go-links)
- 49 - [custom wait group](https://github.com/beyondns/gotips/blob/master/tips64.md#49---custom-wait-group)
- 48 - [ev_loop vs co_routine](https://github.com/beyondns/gotips/blob/master/tips64.md#48---ev_loop-vs-co_routine)
- 47 - [Use C for computation intensive operations](https://github.com/beyondns/gotips/blob/master/tips64.md#47---use-c-for-computation-intensive-operations)
- 46 - [Unique key doc manipulation in MonGo](https://github.com/beyondns/gotips/blob/master/tips64.md#46---unique-key-doc-manipulation-in-mongo)
- 45 - [Go is more a framework than a language](https://github.com/beyondns/gotips/blob/master/tips64.md#45---go-is-more-a-framework-than-a-language)
- 44 - [Facebook messenger chatbot](https://github.com/beyondns/gotips/blob/master/tips64.md#44---facebook-messenger-chatbot)
- 43 - [Passing args as a map](https://github.com/beyondns/gotips/blob/master/tips64.md#43---passing-args-as-a-map)
- 42 - [Optimize Go](https://github.com/beyondns/gotips/blob/master/tips64.md#42---optimize-go)
- 41 - [Telegram bot](https://github.com/beyondns/gotips/blob/master/tips64.md#41---telegram-bot)
- 40 - [Distributed consensus](https://github.com/beyondns/gotips/blob/master/tips64.md#40---distributed-consensus)
- 39 - [Public private struct fields and funcs](https://github.com/beyondns/gotips/blob/master/tips64.md#39---public-private-struct-fields-and-funcs)
- 38 - [Marshal Unmarshal to byte slice using gob](https://github.com/beyondns/gotips/blob/master/tips64.md#38---marshal-unmarshal-to-byte-slice-using-gob)
- 37 - [Time series on leveldb](https://github.com/beyondns/gotips/blob/master/tips64.md#37---time-series-on-leveldb)
- 36 - [Custom type marshal unmarshal json](https://github.com/beyondns/gotips/blob/master/tips64.md#36---custom-type-marshal-unmarshal-json)
- 35 - [Test specific functions](https://github.com/beyondns/gotips/blob/master/tips64.md#35---test-specific-functions)
- 34 - [File path exists](https://github.com/beyondns/gotips/blob/master/tips64.md#34---file-path-exists)
- 33 - [Table driven tests](https://github.com/beyondns/gotips/blob/master/tips64.md#33---table-driven-tests)
- 32 - [Work with consul](https://github.com/beyondns/gotips/blob/master/tips64.md#32---work-with-consul)
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

## #65 - grpc
> 2017-01-26 by [@beyondns](https://github.com/beyondns)

```bash
syntax = "proto3";

package remote;

service Inter {
  rpc HandShake (HandShakeData) returns (HandShakeResult) {}
}

message HandShakeData {
  int32 no = 1;
  string query = 2;
}

message HandShakeResult {
  int32 no = 1;
  string res = 2;
}

```

```go

package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "./remote"
)

// go get google.golang.org/grpc
// go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
// protoc -I ./remote/ ./remote/rem.proto  --go_out=plugins=grpc:remote

type server struct{
	host *string
}

func (s *server) HandShake(ctx context.Context, in *pb.HandShakeData) (*pb.HandShakeResult, error) {
	log.Printf("HandShake %v", in)
	res := &pb.HandShakeResult{
		Res: fmt.Sprintf("Response from %s to query %s" + s.host,in.Query),
	}
	return res, nil
}

// ./node -host 127.0.0.1:50051
// ./node -host 127.0.0.1:50052

func main() {
	host := flag.String("host", "127.0.0.1:50051", "host host:port")
	nodes := flag.String("nodes", "127.0.0.1:50051,127.0.0.1:50052,127.0.0.1:50053",
		"[host:port,host:port,...]")

	flag.Parse()

	go func() {
		lis, err := net.Listen("tcp", *host)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterInterServer(s, &server{host:host})
		reflection.Register(s)
		log.Printf("Serve...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	for _, n := range strings.Split(*nodes, ",") {
		if n != *host {
			go func(n string) {

				conn, err := grpc.Dial(n, grpc.WithInsecure())
				if err != nil {
					log.Printf("did not connect:%s %v", n, err)
					return
				}
				defer conn.Close()
				c := pb.NewInterClient(conn)

				r, err := c.HandShake(context.Background(), 
					&pb.HandShakeData{No: 1, Query: "hello"})
				if err != nil {
					log.Printf("could not HandShake: %v", err)
					return
				}
				log.Printf("HandShake:%s %s", n, r.Res)

			}(n)
		}
	}

	select {}
}

```
* [grpc.io](grpc.io)


## #64 - go vs rust
> 2017-01-20 by [@beyondns](https://github.com/beyondns)


* [tokio](https://tokio.rs/)
* [rust vs go](https://blog.ntpsec.org/2017/01/18/rust-vs-go.html)
* [Hacker News Rust vs. Go (ntpsec.org)](https://news.ycombinator.com/item?id=13430108)
* [c2go](https://github.com/rsc/c2go)

[gotips#32..#63](https://github.com/beyondns/gotips/blob/master/tips64.md)  
[gotips #0..#31](https://github.com/beyondns/gotips/blob/master/tips32.md)  


### General Golang links
* [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

### Inspired by
* [jstips](https://github.com/loverajoel/jstips)

### License
[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png)](http://creativecommons.org/publicdomain/zero/1.0/)

