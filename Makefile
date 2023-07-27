dev:
	docker build -t renandotcorrea/go-micro-auth:latest ./auth/.
	docker build -t renandotcorrea/go-micro-gateway:latest ./gateway/.
	kubectl apply -f ./auth/auth-deploy.yaml
	kubectl apply -f ./auth/service.yaml
	kubectl apply -f ./gateway/gateway-deploy.yaml
	kubectl apply -f ./gateway/service.yaml