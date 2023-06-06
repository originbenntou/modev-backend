FROM golang:1.20.1 as builder

WORKDIR /go/src/github.com/originbenntou/modev-backend

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $WORK_DIR
RUN ls /go/bin


FROM alpine:latest
RUN apk add ca-certificates
COPY --from=builder /go/bin/dataset_generator_back /main

# Lambda Web Adapter
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.6.0 /lambda-adapter /opt/extensions/lambda-adapter
EXPOSE 8080

ENTRYPOINT [ "/main" ]
