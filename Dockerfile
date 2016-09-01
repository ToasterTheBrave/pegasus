FROM golang:1.6.3-alpine

RUN apk --update add git

# Copy App
COPY . /go/src/github.com/truppert/pegasus

# Download Dependencies
RUN \
	cd /go/src/github.com/truppert/pegasus && \
	go get

# Build App
RUN \
	cd /go/src/github.com/truppert/pegasus && \
	go build

WORKDIR /go/src/github.com/truppert/pegasus

CMD ["go", "test", "github.com/truppert/pegasus", "github.com/truppert/pegasus/processor"]