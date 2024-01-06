# Dummy-Website

This is a really small web server written in go. 
This is not the correct way on making a server, but serve only to try to experiment deployment on kubernetes.

## Command

Create namepaces
```bash
kubectl apply -f ./deployment/namespace.yaml
# or
kubectl create ns dummy-ns
```

Create every ressource (but you should probably create namespace first)
```bash
kubectl apply -f ./deployment
```

Or you can pass by the makefile