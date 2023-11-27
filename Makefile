.PHONY: all build move dockerbuild start
DOCKER_IMAGE_NAME=shubhamdixit863/gos3
K8YAML=replicasetnodeport.yaml
AWS_REGION=us-east-2
CLUSTER_NAME=cluster-eksCluster-966cf6b

all:  start  dockerbuild

deploy:
	@echo "deploying on k8 server...."
	 kubectl apply -f ${K8YAML}

start:
	@echo "starting the app server"
	 go run main.go

dockerbuild:
		@echo "building docker image"
		sudo docker build --no-cache -t $(DOCKER_IMAGE_NAME) .
		@echo "pushing docker image"
		docker push $(DOCKER_IMAGE_NAME)
configureaws:
	  aws eks --region $(AWS_REGION)  update-kubeconfig --name  $(CLUSTER_NAME)
