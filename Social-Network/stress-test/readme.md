# Test

Test with health check:

runs a benchmark for 30 seconds, using 12 threads, and keeping 400 HTTP connections open.

```bash
wrk -t12 -c400 -d30s --latency http://127.0.0.1:8080/health-check
```

### Result

```txt
Running 30s test @ http://127.0.0.1:8080/health-check
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.19ms    6.84ms  79.45ms   65.61%
    Req/Sec     5.47k   682.69     7.65k    68.11%
  Latency Distribution
     50%    5.76ms
     75%   10.05ms
     90%   15.50ms
     99%   31.93ms
  1964174 requests in 30.08s, 260.37MB read
  Socket errors: connect 0, read 371, write 0, timeout 0
Requests/sec:  65293.91
Transfer/sec:      8.66MB
```

# NewsFeed API

## Step 1:

Run :

```bash
go run cmd/seeds/main.go
```

## Step 2

### First test with runs a benchmark for 10 seconds, using 1 threads, and keeping 1 HTTP connections open.

```bash
 wrk http://127.0.0.1:8080/v1/newsfeeds -s ./stress-test/api-newfeed.lua --latency -t1 -c1 -d10s
```

Report :

```txt
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   209.25ms   13.96ms 258.20ms   75.00%
    Req/Sec     4.00      0.41     5.00     83.33%
  Latency Distribution
     50%  205.46ms
     75%  220.64ms
     90%  224.96ms
     99%  258.20ms
  48 requests in 10.09s, 45.84MB read
Requests/sec:      4.76
Transfer/sec:      4.54MB

```

### Next test with 30 seconds, using 12 threads, and keeping 400 HTTP connections open.

```bash
wrk http://127.0.0.1:8080/v1/newsfeeds -s ./stress-test/api-newfeed.lua --latency -t12 -c400 -d30s
```

Report :

```txt
Running 30s test @ http://127.0.0.1:8080/v1/newsfeeds
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.00us    0.00us   0.00us     nan%
    Req/Sec     0.00      0.00     0.00       nan%
  Latency Distribution
     50%    0.00us
     75%    0.00us
     90%    0.00us
     99%    0.00us
  0 requests in 30.10s, 0.00B read
  Socket errors: connect 0, read 320, write 0, timeout 0
Requests/sec:      0.00
Transfer/sec:       0.00B
```
