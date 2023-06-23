# go-tcp-keepalive-example

This is a example of exchanging keep-alive probes by golang (SO_KEEPALIVE).

## Environments

```sh
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.2 LTS
Release:        22.04
Codename:       jammy
```

```sh
$ go version
go version go1.20.5 linux/amd64
```

The sample code in this repository runs on both Windows and Linux.

## Requirements

### Task (go-task)

```sh
$ go install github.com/go-task/task/v3/cmd/task@latest
```

### tcpdump

```sh
$ sudo apt update
$ sudo apt install -y tcpdump
```

## How to Run

Open three terminals. the first terminal, do the following

```sh
$ task capture
```

the second terminal, do the following

```sh
$ task server
```

the third terminal, do the following

```sh
$ task client
```

After a while, the first terminal outputs the following results.

The following is the result of a run on Gitpod.

```sh
$ task capture
task: [capture] sudo tcpdump -i lo 'tcp and port 12345 and tcp[13] & 16 != 0 and tcp[12] & 15 = 0'
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on lo, link-type EN10MB (Ethernet), snapshot length 262144 bytes
07:42:59.109181 IP localhost.12345 > localhost.57006: Flags [S.], seq 1096637414, ack 3661790707, win 43690, options [mss 65495,sackOK,TS val 1727619187 ecr 1727619187,nop,wscale 7], length 0
07:42:59.109200 IP localhost.57006 > localhost.12345: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727619187 ecr 1727619187], length 0
07:43:01.120277 IP localhost.57006 > localhost.12345: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727621199 ecr 1727619187], length 0
07:43:01.120283 IP localhost.12345 > localhost.57006: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727621199 ecr 1727619187], length 0
07:43:03.132273 IP localhost.57006 > localhost.12345: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727623211 ecr 1727621199], length 0
07:43:03.132283 IP localhost.12345 > localhost.57006: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727623211 ecr 1727619187], length 0
07:43:05.148284 IP localhost.57006 > localhost.12345: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727625227 ecr 1727623211], length 0
07:43:05.148295 IP localhost.12345 > localhost.57006: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727625227 ecr 1727619187], length 0
07:43:07.164286 IP localhost.57006 > localhost.12345: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727627243 ecr 1727625227], length 0
07:43:07.164292 IP localhost.12345 > localhost.57006: Flags [.], ack 1, win 342, options [nop,nop,TS val 1727627243 ecr 1727619187], length 0
07:43:09.116306 IP localhost.57006 > localhost.12345: Flags [F.], seq 1, ack 1, win 342, options [nop,nop,TS val 1727629195 ecr 1727627243], length 0
07:43:09.116396 IP localhost.12345 > localhost.57006: Flags [F.], seq 1, ack 2, win 342, options [nop,nop,TS val 1727629195 ecr 1727629195], length 0
07:43:09.116440 IP localhost.57006 > localhost.12345: Flags [.], ack 2, win 342, options [nop,nop,TS val 1727629195 ecr 1727629195], length 0
```
