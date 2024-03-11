FROM golang:1.22.0-alpine3.19

WORKDIR /try/backend
COPY . .

#RUN go mod init try-standard-layout
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
RUN mkdir test

CMD [ "air" ]