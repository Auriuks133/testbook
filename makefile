TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG
test:
	go test .


build:
	go build -ldflags "-X main.version=$(TAG)" -o news .
pack: build
	docker build -t $(TAG) .
upload:
	docker push $(TAG)
deploy:
	envsubst < k8s/deployment.yml | kubectl apply -f -