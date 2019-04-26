FROM golang:alpine

ENV AWS_ACCESS_KEY_ID=YOUR_AKID
ENV AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
ENV AWS_SESSION_TOKEN=TOKEN

ADD ./src /go/src/app

WORKDIR /go/src/app

ENV PORT=3001
CMD ["go", "run", "main.go"]