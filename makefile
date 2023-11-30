APPNAME=dummy-website
VERSION=0.0.1

build:
	go build -o ${APPNAME}

publish:
	docker build --pull -t "${APPNAME}:${VERSION}" --build-arg APP_VERSION=${VERSION} .

# deploy:
