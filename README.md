# What is this?

Its a really simple go app that returns a clients IP address

# Why?

Its just something I wrote to demonstrate different ways of returning a clients IP address, either directly or through a proxy

# How does it work?

## Server
```
go run main.go
running http server on port 8080...
falling back to remote addr
no x-forwarded-for, next...
```

## Client
```
# direct
curl localhost:8080
{
  "ip": "127.0.0.1"
}
  
# through nginx
curl https://server.home.michaelc.dev/icanhazip
{
  "ip": "192.168.1.81"
}
```
