package main

import (
	"os"
	"fmt"
	"strconv"
	flags "github.com/jessevdk/go-flags"
	"time"
	"google.golang.org/api/iterator"
	"golang.org/x/net/context"
	monitoring "cloud.google.com/go/monitoring/apiv3"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

type Options struct {
	Project  string `short:"p" long:"project"  description:"gcp project id"       required:"true"`
	Instance string `short:"i" long:"instance" description:"gce instance name"    required:"true"`
	Critical string `short:"c" long:"critical" description:"critical threshold"   required:"true"`
  	Warning  string `short:"w" long:"warning"  description:"warning threshold"    required:"true"`
  	Auth     string `short:"a" long:"auth"     description:"gcp authenticate key" required:"true"`
}

func main() {
    // 引数解析処理
    var opts Options
  	parser := flags.NewParser(&opts, flags.IgnoreUnknown)
  	_, err := parser.Parse()
    if err != nil {
		fmt.Println("Missing required arguments.")
        parser.WriteHelp(os.Stdout)
        os.Exit(1)
  	}

	// 定数定義
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", opts.Auth)

	// GCPのSDKを使う準備
	// https://godoc.org/cloud.google.com/go/monitoring/apiv3#NewMetricClient
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	
	// フィルタの生成
	var filter string = "metric.type = \"compute.googleapis.com/instance/uptime\" "
	filter += "AND metric.label.instance_name = \"" + opts.Instance + "\" "

	// 取得期間の生成
	// https://godoc.org/google.golang.org/genproto/googleapis/monitoring/v3#ListTimeSeriesRequest
	var critical int64
	critical, _ = strconv.ParseInt(opts.Critical, 10, 64)
	var warning int64
	warning, _ = strconv.ParseInt(opts.Warning, 10, 64)
	unixNow := time.Now().Unix()
	req := &monitoringpb.ListTimeSeriesRequest{
		Name : "projects/" + opts.Project,
		Filter : filter,
		Interval: &monitoringpb.TimeInterval{
			EndTime: &googlepb.Timestamp{
				Seconds: unixNow - 300,
			},
			StartTime: &googlepb.Timestamp{
				Seconds: unixNow - (critical * 60),
			},
		},
	}

	uptime := 0.0
	it := c.ListTimeSeries(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		for _, point := range resp.Points {
			uptime += point.GetValue().GetDoubleValue()
		}
		_ = resp
	}

	if (int64(uptime) <=  0) {
		output(2, "CRITICAL - Instance " + opts.Instance + " stopped for " + opts.Critical + " minutes.")
	}

	if (int64(uptime) <=  warning * 60) {
		output(1, "WARNING - Instance " + opts.Instance + " stopped for " + opts.Critical + " minutes.")
	}

	output(0, "OK - Instance " + opts.Instance + " is UP.")
}

func output(status int, message string) {
	fmt.Println(message)
	os.Exit(status)
}