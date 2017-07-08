# Caso 1

## Sin proxy

ab -c 100 -n 100000 'http://127.0.0.1:8081/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   2.470 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      21899334 bytes
HTML transferred:       10100000 bytes
Requests per second:    40483.96 [#/sec] (mean)
Time per request:       2.470 [ms] (mean)
Time per request:       0.025 [ms] (mean, across all concurrent requests)
Transfer rate:          8657.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:     0    2   1.6      2      18
Waiting:        0    2   1.6      1      18
Total:          0    2   1.6      2      18

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      3
  80%      3
  90%      4
  95%      6
  98%      7
  99%      9
 100%     18 (longest request)


## Nginx

ab -c 100 -n 100000 'http://127.0.0.1:8001/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        nginx/1.4.6
Server Hostname:        127.0.0.1
Server Port:            8001

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   6.014 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      26799446 bytes
HTML transferred:       10100000 bytes
Requests per second:    16628.59 [#/sec] (mean)
Time per request:       6.014 [ms] (mean)
Time per request:       0.060 [ms] (mean, across all concurrent requests)
Transfer rate:          4351.92 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       5
Processing:     0    6   2.0      6      18
Waiting:        0    6   2.0      6      17
Total:          0    6   2.0      6      18

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      6
  75%      7
  80%      7
  90%      9
  95%     10
  98%     11
  99%     12
 100%     18 (longest request)

## Go-Balancer

ab -c 100 -n 100000 'http://127.0.0.1:8080/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   11.205 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      21899334 bytes
HTML transferred:       10100000 bytes
Requests per second:    8924.64 [#/sec] (mean)
Time per request:       11.205 [ms] (mean)
Time per request:       0.112 [ms] (mean, across all concurrent requests)
Transfer rate:          1908.63 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0   11   6.4     10     208
Waiting:        0   10   6.1      9     208
Total:          0   11   6.4     10     208

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     12
  75%     14
  80%     15
  90%     18
  95%     21
  98%     28
  99%     34
 100%    208 (longest request)

# Caso 2

Levantamos 5 servidores tests en los puertos 8081, 8082, 8083, 8084, 8085

## Nginx

ab -c 100 -n 100000 'http://127.0.0.1:8001/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        nginx/1.4.6
Server Hostname:        127.0.0.1
Server Port:            8001

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   5.590 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      26799728 bytes
HTML transferred:       10100000 bytes
Requests per second:    17890.34 [#/sec] (mean)
Time per request:       5.590 [ms] (mean)
Time per request:       0.056 [ms] (mean, across all concurrent requests)
Transfer rate:          4682.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       5
Processing:     0    6   2.4      5      39
Waiting:        0    5   2.4      5      39
Total:          0    6   2.3      5      39

Percentage of the requests served within a certain time (ms)
  50%      5
  66%      6
  75%      7
  80%      7
  90%      8
  95%      9
  98%     11
  99%     13
 100%     39 (longest request)

## Go-Balancer

ab -c 100 -n 100000 'http://127.0.0.1:8080/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   13.282 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      21899728 bytes
HTML transferred:       10100000 bytes
Requests per second:    7529.11 [#/sec] (mean)
Time per request:       13.282 [ms] (mean)
Time per request:       0.133 [ms] (mean, across all concurrent requests)
Transfer rate:          1610.21 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     0   13   7.6     12     138
Waiting:        0   12   7.4     11     137
Total:          0   13   7.6     12     138

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     15
  75%     17
  80%     18
  90%     22
  95%     26
  98%     32
  99%     38
 100%    138 (longest request)

# Caso 3

Levantamos 2 servidores tests en los puertos 8081, 8082

## Nginx

ab -c 100 -n 100000 'http://127.0.0.1:8001/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        nginx/1.4.6
Server Hostname:        127.0.0.1
Server Port:            8001

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   4.624 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      26799248 bytes
HTML transferred:       10100000 bytes
Requests per second:    21628.18 [#/sec] (mean)
Time per request:       4.624 [ms] (mean)
Time per request:       0.046 [ms] (mean, across all concurrent requests)
Transfer rate:          5660.34 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     0    5   2.0      4      28
Waiting:        0    5   2.0      4      28
Total:          0    5   2.0      4      28

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      5
  75%      5
  80%      6
  90%      7
  95%      8
  98%     11
  99%     12
 100%     28 (longest request)

## Go-Balancer

ab -c 100 -n 100000 'http://127.0.0.1:8080/foo?total=100'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /foo?total=100
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   12.276 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      21899238 bytes
HTML transferred:       10100000 bytes
Requests per second:    8145.95 [#/sec] (mean)
Time per request:       12.276 [ms] (mean)
Time per request:       0.123 [ms] (mean, across all concurrent requests)
Transfer rate:          1742.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0   12   6.6     11     108
Waiting:        0   11   6.4     11     108
Total:          0   12   6.6     11     108

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     14
  75%     15
  80%     16
  90%     20
  95%     23
  98%     28
  99%     35
 100%    108 (longest request)
