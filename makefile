TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG
test:
	go test .


build:
	go build  
pack: build
	docker build -t  .
upload:
	docker push 
deploy:
	envsubst < k8s/deployment.yml | kubectl apply -f -