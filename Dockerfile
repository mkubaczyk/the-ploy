FROM golang:1.11.1-alpine3.7 as builder
ARG BUILD_FILE
ARG PROJECT=$GOPATH/src/github.com/mkubaczyk/theploy
RUN mkdir -p $PROJECT \
    && apk add --update --no-cache curl git \
    && curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR $PROJECT
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure --vendor-only -v
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app $BUILD_FILE

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
