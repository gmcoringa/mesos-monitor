package main

import(
	"fmt"
	"os"
	"time"
	"github.com/gmcoringa/mesos-monitor/config"
	"github.com/gmcoringa/mesos-monitor/mesos"
	"github.com/cactus/go-statsd-client/statsd"
	"github.com/jasonlvhit/gocron"
)

func main() {
	configuration := config.Load()
	statsd, err := statsd.NewBufferedClient(configuration.StatsDHost, configuration.Prefix, 100 *time.Millisecond, 512)
	mesos := mesos.NewMesos(configuration.MesosHost, configuration.ConnectionTimeoutMS)

	if err != nil {
		fmt.Println("Failed to create StatsD Client: ", err)
		os.Exit(1)
	}

	gocron.Every(configuration.IntervalInSeconds).Seconds().Do(send, mesos, statsd)
	<-gocron.Start()
}

func send(mesos mesos.MesosCollector, statsd statsd.Statter) {
	stats := mesos.Collect()

	for k, v := range stats {
		statsd.Raw(k, fmt.Sprintf("%f|g", v), 1.0)
	}
}
