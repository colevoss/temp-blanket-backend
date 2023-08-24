run:
	go run "$(CURDIR)/cmd/tempblanket/main.go"

docker-build:
	docker build -t colevoss/temp-blanket-backend .

docker-run:
	docker run \
		-it \
		--rm \
		-p 8080:8080 \
		--env SYNOPTIC_API_TOKEN=$(SYNOPTIC_API_TOKEN) \
		colevoss/temp-blanket-backend
