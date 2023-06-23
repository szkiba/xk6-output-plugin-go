# xk6-output-plugin-go

Go plugin SDK for [xk6-output-plugin](https://github.com/szkiba/xk6-output-plugin).

## Documentation

[API documentation](https://pkg.go.dev/github.com/szkiba/xk6-output-plugin-go/output)

```bash
go get github.com/szkiba/xk6-output-plugin-go
```

## Example

```go
package main

import (
  "context"
  "time"

  "github.com/hashicorp/go-hclog"
  "github.com/szkiba/xk6-output-plugin-go/output"
)

type example struct{}

func (e *example) Init(ctx context.Context, params *output.Params) (*output.Info, error) {
  hclog.L().Info("init")

  return &output.Info{Description: "example-go plugin"}, nil // nolint:exhaustruct
}

func (e *example) Start(ctx context.Context) error {
  hclog.L().Info("start")

  return nil
}

func (e *example) Stop(ctx context.Context) error {
  hclog.L().Info("stop")

  return nil
}

func (e *example) AddMetrics(ctx context.Context, metrics []*output.Metric) error {
  hclog.L().Info("metrics")

  for _, metric := range metrics {
    hclog.L().Info(metric.Name,
      "metric.type", metric.Type.String(),
      "metric.contains", metric.Contains.String(),
    )
  }

  return nil
}

func (e *example) AddSamples(ctx context.Context, samples []*output.Sample) error {
  hclog.L().Info("samples")

  for _, sample := range samples {
    hclog.L().Info(sample.Metric,
      "sample.time", time.UnixMilli(sample.Time).Format(time.RFC3339),
      "sample.value", sample.Value,
    )
  }

  return nil
}

func main() {
  output.Serve(new(example))
}
```

<details>
  <summary>Output</summary>

```plain
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

INFO[0000] init                                          plugin=example
INFO[0000] start                                         plugin=example
  execution: local
     script: script.js
     output: example-go plugin

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0001] metrics                                       plugin=example
INFO[0001] http_reqs                                     metric.contains=DEFAULT metric.type=COUNTER plugin=example
INFO[0001] http_req_duration                             metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_blocked                              metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_connecting                           metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_tls_handshaking                      metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_sending                              metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_waiting                              metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_receiving                            metric.contains=TIME metric.type=TREND plugin=example
INFO[0001] http_req_failed                               metric.contains=DEFAULT metric.type=RATE plugin=example
INFO[0001] samples                                       plugin=example
INFO[0001] http_reqs                                     plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=1
INFO[0001] http_req_duration                             plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.270604
INFO[0001] http_req_blocked                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=149.204512
INFO[0001] http_req_connecting                           plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.481335
INFO[0001] http_req_tls_handshaking                      plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0
INFO[0001] http_req_sending                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0.062589
INFO[0001] http_req_waiting                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.104524
INFO[0001] http_req_receiving                            plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0.103491
INFO[0001] http_req_failed                               plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0
INFO[0001] http_reqs                                     plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=1
INFO[0001] http_req_duration                             plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.571097
INFO[0001] http_req_blocked                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=248.7757
INFO[0001] http_req_connecting                           plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=123.46025
INFO[0001] http_req_tls_handshaking                      plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.23569
INFO[0001] http_req_sending                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0.049269
INFO[0001] http_req_waiting                              plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=125.398706
INFO[0001] http_req_receiving                            plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0.123122
INFO[0001] http_req_failed                               plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=0
INFO[0002] metrics                                       plugin=example
INFO[0002] vus                                           metric.contains=DEFAULT metric.type=GAUGE plugin=example
INFO[0002] vus_max                                       metric.contains=DEFAULT metric.type=GAUGE plugin=example
INFO[0002] data_sent                                     metric.contains=DATA metric.type=COUNTER plugin=example
INFO[0002] data_received                                 metric.contains=DATA metric.type=COUNTER plugin=example
INFO[0002] iteration_duration                            metric.contains=TIME metric.type=TREND plugin=example
INFO[0002] iterations                                    metric.contains=DEFAULT metric.type=COUNTER plugin=example
INFO[0002] samples                                       plugin=example
INFO[0002] vus                                           plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=1
INFO[0002] vus_max                                       plugin=example sample.time="2023-06-23T09:35:02+02:00" sample.value=1
INFO[0002] data_sent                                     plugin=example sample.time="2023-06-23T09:35:03+02:00" sample.value=542
INFO[0002] data_received                                 plugin=example sample.time="2023-06-23T09:35:03+02:00" sample.value=17310
INFO[0002] iteration_duration                            plugin=example sample.time="2023-06-23T09:35:03+02:00" sample.value=1649.955992
INFO[0002] iterations                                    plugin=example sample.time="2023-06-23T09:35:03+02:00" sample.value=1
INFO[0002] stop                                          plugin=example

     data_received..................: 17 kB 11 kB/s
     data_sent......................: 542 B 329 B/s
     http_req_blocked...............: avg=198.99ms min=149.2ms  med=198.99ms max=248.77ms p(90)=238.81ms p(95)=243.79ms
     http_req_connecting............: avg=124.47ms min=123.46ms med=124.47ms max=125.48ms p(90)=125.27ms p(95)=125.38ms
   ✓ http_req_duration..............: avg=125.42ms min=125.27ms med=125.42ms max=125.57ms p(90)=125.54ms p(95)=125.55ms
       { expected_response:true }...: avg=125.42ms min=125.27ms med=125.42ms max=125.57ms p(90)=125.54ms p(95)=125.55ms
   ✓ http_req_failed................: 0.00% ✓ 0        ✗ 2  
     http_req_receiving.............: avg=113.3µs  min=103.49µs med=113.3µs  max=123.12µs p(90)=121.15µs p(95)=122.14µs
     http_req_sending...............: avg=55.92µs  min=49.26µs  med=55.92µs  max=62.58µs  p(90)=61.25µs  p(95)=61.92µs 
     http_req_tls_handshaking.......: avg=62.61ms  min=0s       med=62.61ms  max=125.23ms p(90)=112.71ms p(95)=118.97ms
     http_req_waiting...............: avg=125.25ms min=125.1ms  med=125.25ms max=125.39ms p(90)=125.36ms p(95)=125.38ms
     http_reqs......................: 2     1.212069/s
     iteration_duration.............: avg=1.64s    min=1.64s    med=1.64s    max=1.64s    p(90)=1.64s    p(95)=1.64s   
     iterations.....................: 1     0.606035/s
     vus............................: 1     min=1      max=1
     vus_max........................: 1     min=1      max=1


running (00m01.7s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m01.7s/10m0s  1/1 iters, 1 per VU
```
</details>