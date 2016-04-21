package all

import (
	_ "github.com/sythe21/telegraf/plugins/outputs/amon"
	_ "github.com/sythe21/telegraf/plugins/outputs/amqp"
	_ "github.com/sythe21/telegraf/plugins/outputs/cloudwatch"
	_ "github.com/sythe21/telegraf/plugins/outputs/datadog"
	_ "github.com/sythe21/telegraf/plugins/outputs/file"
	_ "github.com/sythe21/telegraf/plugins/outputs/graphite"
	_ "github.com/sythe21/telegraf/plugins/outputs/influxdb"
	_ "github.com/sythe21/telegraf/plugins/outputs/kafka"
	_ "github.com/sythe21/telegraf/plugins/outputs/kinesis"
	_ "github.com/sythe21/telegraf/plugins/outputs/librato"
	_ "github.com/sythe21/telegraf/plugins/outputs/mqtt"
	_ "github.com/sythe21/telegraf/plugins/outputs/nsq"
	_ "github.com/sythe21/telegraf/plugins/outputs/opentsdb"
	_ "github.com/sythe21/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/sythe21/telegraf/plugins/outputs/riemann"
)
