run:
	go run "$(CURDIR)/cmd/tempblanket/main.go"

docker-build:
	"$(CURDIR)/scripts/docker-build.sh"
# docker build -t colevoss/temp-blanket-backend .

docker-run:
	"$(CURDIR)/scripts/docker-run.sh"

# docker run \
# 	-it \
# 	--rm \
# 	-p 8080:8080 \
# 	--env SYNOPTIC_API_TOKEN=$(SYNOPTIC_API_TOKEN) \
# 	colevoss/temp-blanket-backend

prepare-deployment:
	"$(CURDIR)/scripts/create-env-file.sh"

terraform-init:
	cd ./deploy && \
		terraform init \
			-backend-config=./config/gcs.tfbackend

terraform-plan:
	cd ./deploy && \
		terraform plan \
			-var-file=./config/input.tfvars
