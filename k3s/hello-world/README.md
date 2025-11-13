# Hello World in k3s

Развертка приложения Hello world на k3s

## Проверка статусов

```
kubectl get pods
```
```
kubectl get services
```
```
kubectl get ingress
```


## Проверка работы

```bash
kubectl port-forward service/hello-world-service 8080:80
```

```bash
curl http://localhost:8080
```