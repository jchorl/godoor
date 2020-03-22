default:
	docker run -it --rm \
		-e GOOS=linux \
		-e GOARCH=arm \
		-v $$PWD:/godoor \
		-w /godoor \
		golang:1.14.1 \
		go build
