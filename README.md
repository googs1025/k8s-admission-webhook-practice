# k8s-admission-webhook-practice


docker run --rm -it -v /root/k8s-admission-webhook:/app -w /app -e GOPROXY=https://goproxy.cn -e CGO_ENABLED=0  golang:1.18.7-alpine3.15 go build -o ./myhook .

