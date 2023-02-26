FROM golang:1.16.2-alpine AS builder

LABEL stage=gobuilder

WORKDIR /build/zero

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on

COPY . .
ENV TZ Asia/Shanghai
#RUN go build -v -ldflags "-X 'xiudong/cli/cmd.version=$(go version)' \
#	-X 'xiudong/cli/cmd.commit=$(git show -s --format=%H)' \
#	-X 'xiudong/cli/cmd.date=$(date -Iseconds)'" \
#	-o /app/xiudong xiudong/cli
RUN go build -ldflags="-s -w" -o /app/xiudong xiudong/cli


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/xiudong /app/xiudong

ENTRYPOINT ["./xiudong"]
