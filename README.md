# go-yanc
Yet Another Network Calculator


![Static Badge](https://img.shields.io/badge/Project-IN_PROGRESS-orange) ![Static Badge](https://img.shields.io/badge/Go-blue)

Simple CLI Application in Go to split network subnets or perform RIPE WHOIS query on the IP

Usage example in Windows: 

```
go-yanc.exe -n 1.1.1.1/25 -split /26
```
![Example](pictures/split.png)

```
go-yanc.exe -w 8.8.8.8
```
![Example](pictures/whois.png)
