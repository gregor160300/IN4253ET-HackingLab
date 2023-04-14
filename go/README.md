# Proof of concept amplification attack

A go program that runs a DNS amplification attack using the provided nameservers and domains.

## How to run
Note: this requires root permissions, since it uses sockets.  
Target ip is the only required argument, the rest of the arguments are optional

```
go run . --target_ip=<target_ip> -i <filename> -n <number of threads> -d <duration> --iface=<network interface>
```

One can get a help menu by running the following:

```
go run . --help
```

## Attack simulation using docker
A simple attack simulation can be used by running

```
docker-compose up (-d)
```


