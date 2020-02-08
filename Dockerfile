FROM golang:1.12 AS build-env

# Download and install dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy and build source
WORKDIR $GOPATH/src/github.com/bcelenza/carrier
COPY . ./
RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /carrier .

FROM scratch
COPY --from=build-env /carrier ./
ENTRYPOINT ["./carrier"]
