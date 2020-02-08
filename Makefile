PROJECT_NAME=carrier
BUILD_DIR=build
BUILD_OUTPUT=$(BUILD_DIR)/$(PROJECT_NAME)
IMAGE_TAG?=latest
ECR_IMAGE=$(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(PROJECT_NAME):$(IMAGE_TAG)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: dep
dep:
	dep ensure

.PHONY: build
build: dep
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_OUTPUT) .

.PHONY: run
run:
	$(BUILD_OUTPUT)

.PHONY: image
image:
	docker build . -t $(PROJECT_NAME):$(IMAGE_TAG)

.PHONY: push-ecr
push-ecr:
	docker tag $(PROJECT_NAME):$(IMAGE_TAG) $(ECR_IMAGE)
	docker push $(ECR_IMAGE)