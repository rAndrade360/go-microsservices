dev-up:
	docker build -t renandotcorrea/go-micro-auth:latest ./auth/.
	docker build -t renandotcorrea/go-micro-gateway:latest ./gateway/.
	docker build -t renandotcorrea/go-micro-mysql:latest ./auth/mysql/.
	kubectl apply -f ./auth/.
	kubectl apply -f ./auth/mysql/.
	kubectl apply -f ./gateway/.
dev-down:
	kubectl delete -f ./auth/.
	kubectl delete -f ./auth/mysql/.
	kubectl delete -f ./gateway/.