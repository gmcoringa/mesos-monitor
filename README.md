# Mesos Monitor

Mesos Monitor collects all metrics exposed by Mesos and send them as gauges to StatsD.

Uses goCron to schedule metric collection by a configurable interval.

## Building from Sources

You can also get from the releases.

### Requisites

Two libraries are required: [goCron](https://github.com/jasonlvhit/gocron) and [go-statsd-client](https://github.com/cactus/go-statsd-client).


```shell
go get github.com/jasonlvhit/gocron
go get github.com/cactus/go-statsd-client/statsd
```

### Compile

No secret just run:

```shell
go build mesos-monitor.go
```

## Configurations

* **-m**: Mesos host and port, default value ``http://localhost:5050``.
* **-s**: StatsD host and port, default value ``localhost:8125``.
* **--p**: Metric prefix to be used in StatsD.
* **-i**: Interval in seconds to collect metrics from Mesos, default value ``60``.
* **-t**: Connection timeout in miliseconds to retrieve metrics from Mesos, default value ``100``.

Example:

```shell
mesos-monitor -m http://mesos-master:5050 -s statsd:8125 -p "stats.mesos" -i 10 -t 50
```
