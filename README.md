
使用帮助
========

## Step 1: 启动`gortcd` Server

```text
./gortcd
{"level":"info","ts":1568725431.6296182,"msg":"config file used","path":"/home/xiaolin/work/AdamEve/gostun_client/gortcd.yml"}
{"level":"info","ts":1568725431.6296756,"msg":"parsed credentials","n":0}
{"level":"info","ts":1568725431.629685,"msg":"realm","k":"gortc.io"}
{"level":"info","ts":1568725431.6297455,"logger":"filter.peer","msg":"default action set","action":"allow"}
{"level":"info","ts":1568725431.6297612,"logger":"filter.client","msg":"default action set","action":"allow"}
{"level":"info","ts":1568725431.629765,"msg":"will be sending SOFTWARE attribute","software":"gortcd"}
{"level":"info","ts":1568725431.6298447,"msg":"api listening","addr":"localhost:3257"}
{"level":"info","ts":1568725431.629856,"logger":"reload","msg":"subscribed to SIGUSR2"}
{"level":"info","ts":1568725431.629874,"msg":"got addr","addr":"0.0.0.0:3478"}
{"level":"warn","ts":1568725431.6298842,"msg":"running on all interfaces"}
{"level":"warn","ts":1568725431.6298914,"msg":"picking addr from ICE"}
{"level":"warn","ts":1568725431.6301045,"msg":"got","a":"fe80::42:b4ff:fecd:ff71 (zone br-46820db1e687) [40]"}
{"level":"warn","ts":1568725431.6301196,"msg":"got","a":"fe80::42:bbff:fe37:5d68 (zone docker0) [40]"}
{"level":"warn","ts":1568725431.630126,"msg":"got","a":"fe80::5c5a:52a8:e7c:b203 (zone wlp5s0) [40]"}
{"level":"warn","ts":1568725431.6301312,"msg":"got","a":"172.17.0.1 [35]"}
{"level":"warn","ts":1568725431.6301363,"msg":"using","a":"172.17.0.1 [35]"}
{"level":"warn","ts":1568725431.630146,"msg":"got","a":"172.18.0.1 [35]"}
{"level":"warn","ts":1568725431.6301513,"msg":"using","a":"172.18.0.1 [35]"}
{"level":"warn","ts":1568725431.6301565,"msg":"got","a":"192.168.1.15 [35]"}
{"level":"warn","ts":1568725431.630161,"msg":"using","a":"192.168.1.15 [35]"}
{"level":"info","ts":1568725431.6301944,"msg":"gortc/gortcd listening","addr":"192.168.1.15:3478","network":"udp"}
{"level":"info","ts":1568725431.630243,"msg":"gortc/gortcd listening","addr":"172.17.0.1:3478","network":"udp"}
{"level":"info","ts":1568725431.63027,"msg":"gortc/gortcd listening","addr":"172.18.0.1:3478","network":"udp"}
```

- 最后三行，有监听的地址，选择一个可以访问的，例如`192.168.1.15:3478`

## Step 2: 使用`stun-client`连接到server

```text
./stun-client

# 或者
./stun-client 192.168.1.15:3478 10
# 其中: 192.168.1.15:3478 是连接的目的地址(UDP)
#      10是每个数据包发送的间隔
```

- 你将看到如下的输出：

```text
target stun server                  :  stun1.l.google.com:19302
sleep interval for each iteration is:  0
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
your address on Gateway is, IP: 124.78.156.52 Port: 1588
```
