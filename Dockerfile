FROM golang:alpine AS build-env
WORKDIR /go/src
COPY . /go/src/slark


RUN apk add --no-cache go git

ENV GIT_TERMINAL_PROMPT=1
WORKDIR /go/src/slark
RUN go get
RUN go build -o ./bin/main .


FROM alpine
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/slark/bin /app
COPY --from=build-env /go/src/slark/templates /app/templates
COPY --from=build-env /go/src/slark/static /app/static
COPY --from=build-env /go/src/slark/files /app/files

EXPOSE 8888
ENTRYPOINT [ "./main" ]