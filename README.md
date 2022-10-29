# k8s-admission-webhook-practice
## 开发自定义准入控制器demo

### 步骤流程
#### 1. 当开发程序完成后，需要打包成可执行文件
```bigquery
docker run --rm -it -v /root/k8s-admission-webhook:/app -w /app -e GOPROXY=https://goproxy.cn -e CGO_ENABLED=0  golang:1.18.7-alpine3.15 go build -o ./myhook .

--rm 结束后就删除
-it
-v 挂载 /宿主机目录:/docker目录
-e 环境变量
```
#### 2. 布署到k8s集群中
```bigquery
# 顺序不能乱
kubectl apply -f deploy.yaml
kubectl apply -f admconfig.yaml
kubectl apply -f newpod.yaml
```

