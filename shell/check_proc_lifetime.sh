#/bin/bash
ProcessName=$1
Thresiold=$2
OwnPid=$$ 

# このスクリプトを除く、ProcessNameにマッチするプロセスのプロセスIDを取得する
MatchProcessList=`ps -ef | grep "$ProcessName" | grep -v $OwnPid | grep -v grep | awk '{print $2}'`
if [ ! -n "$MatchProcessList" ]; then
    echo 'OK - Process not match'
    exit 0
fi

# マッチしたプロセスIDごとにしきい値を超過していないかチェックする
for i in $MatchProcessList
do
    # プロセスIDの起動時間を取得する
    TIME=`ps -o lstart --noheader -p $i`
    if [ -n "$TIME" ]; then
        # ps -o lstart の結果は Fri Jan 11 19:22:44 2019 出力なので、UnixTimeに戻す
        StartTime=`date +%s -d "$TIME"`
        CurrentTime=`date +%s`
        ElapsedTime=`expr $CurrentTime - $StartTime`
    else
        ElapsedTime=0
    fi
    # ElapsedTime: 経過時間がしきい値より大きければ、超過している
    if [ $ElapsedTime -gt $Thresiold ] ; then
        echo 'CRITICAL - Process id ' $i ' over run !! life time ' $ElapsedTime
        exit 2
    fi
done

# 最後まで処理を終えるのは、いずれのプロセスもしきい値を超えていないとき
echo 'OK - Any process not over run'
exit 0
