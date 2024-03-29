FROM golang as builder

# Add Maintainer Info
LABEL maintainer="Muhammad Indrawan <me@indra.codes>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/aboutindra/fokusngoding-compiler-golang

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN apk add --no-cache util-linux
ENTRYPOINT export UUID=`uuidgen` && echo $UUIDFROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/aboutindra/fokusngoding-compiler-golang /app/fokusngoding-compiler-golang
WORKDIR "/app/fokusngoding-compiler-golang"
EXPOSE 4000
ENTRYPOINT ./fc-golang
