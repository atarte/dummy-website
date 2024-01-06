APPNAME=dummy-website
VERSION=0.0.1

# build:
# 	go build -o ${APPNAME}

# test:

# clean:

publish:
	docker build --pull -t "${APPNAME}:${VERSION}" --build-arg APP_VERSION=${VERSION} .

deploy:
	kubectl apply -f deployment/namespace.yaml
	kubectl apply -f deployment

undeploy:
	kubectl delete -f deployment
