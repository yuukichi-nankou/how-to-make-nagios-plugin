#!/usr/bin/python
# -*- coding: utf-8 -*-
import argparse
import requests
import json
import os
from google.cloud import monitoring_v3
from google.cloud.monitoring_v3.query import Query
import datetime

# 引数の解析
parser = argparse.ArgumentParser()
parser.add_argument("--project")
parser.add_argument("--instance")
parser.add_argument("--critical")
parser.add_argument("--warning")
parser.add_argument("--key")
args = parser.parse_args()

# 環境変数
os.environ['GOOGLE_APPLICATION_CREDENTIALS'] = args.key

# GCP SDK
client = monitoring_v3.MetricServiceClient()
name = client.project_path(args.project)

# 値を取得期間を指定する
# 取得しようとしている値は最長 240 秒値が格納されないことがあるので、5分以前の値を取得する
# https://cloud.google.com/monitoring/api/metrics_gcp#gcp-compute
endtime = datetime.datetime.utcnow()
endtime = endtime - datetime.timedelta(minutes=5)

# メトリック取得クエリの発行
# https://googleapis.dev/python/monitoring/latest/query.html#google.cloud.monitoring_v3.query.Query
query = Query(
            client,
            args.project,
            'compute.googleapis.com/instance/uptime', # 取得するメトリック
            end_time=endtime,
            days=0,
            hours=0,
            minutes=int(args.critical), # 取得する分数を指定する
        )

# 指定したインスタンスのみに値を絞る
query = query.select_metrics(instance_name=args.instance) 

uptime = 0.0
for timeseries in query:
    for point in timeseries.points:
        uptime += point.value.double_value # UPTIMEを合算し、起動時間とする

# 起動時間が0の時は起動していない、CRITICALエラーをだす
if (uptime <= 0.0) :
    print("CRITICAL - Instance {} stopped for {} minutes.".format(args.instance, args.critical))
    exit(2)

# 起動時間がWARNINGしきい値未満の場合、WARNINGエラーをだす
if (uptime <= int(args.warning) * 60) :
    print("WARNING - Instance {} stopped for {} minutes.".format(args.instance, args.warning))
    exit(1)

print("OK - Instance {} is UP.".format(args.instance))
exit(0)