package config

import "flag"

type MesosConfig struct {
	MesosHost     		string
	StatsDHost			string
	Prefix   			string
	IntervalInSeconds 	uint64
	ConnectionTimeoutMS uint64
}

func Load() (MesosConfig){
	mesosConfig := MesosConfig{}

	mesosHost 	:= flag.String("m", "http://localhost:5050", "Mesos Host, format: host:port")
	statsDHost 	:= flag.String("s", "localhost:8125", "StatsD Host, format: host:port")
	prefix 		:= flag.String("p", "", "Stats prefix for metrics")
	interval 	:= flag.Uint64("i", 60, "Interval in seconds to collect metrics from Mesos")
	timeout 	:= flag.Uint64("t", 100, "Connection timeout in miliseconds to retrieve metrics from Mesos")

	flag.Parse()

	mesosConfig.MesosHost = *mesosHost
	mesosConfig.StatsDHost = *statsDHost
	mesosConfig.Prefix = *prefix
	mesosConfig.IntervalInSeconds = *interval
	mesosConfig.ConnectionTimeoutMS = *timeout

	return mesosConfig
}
