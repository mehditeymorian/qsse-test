

lint:
	golangci-lint run

build:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build main.go
	docker build -t $(IMAGE):latest .
	docker tag $(IMAGE):latest $(REGISTRY)/$(NAMESPACE)/$(IMAGE):latest
	docker push $(REGISTRY)/$(NAMESPACE)/$(IMAGE):latest
	docker rmi $(REGISTRY)/$(NAMESPACE)/$(IMAGE):latest
	docker rmi $(IMAGE):latest

server-i:
	helm install $(IMAGE)-server ./deployment/qsset-server

server-u:
	helm uninstall $(IMAGE)-server

generator-i:
	helm install $(IMAGE)-generator ./deployment/qsset-generator

generator-u:
	helm uninstall $(IMAGE)-generator

subscriber-i:
	helm install $(IMAGE)-subscriber ./deployment/qsset-subscriber

subscriber-u:
	helm uninstall $(IMAGE)-subscriber

