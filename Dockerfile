FROM golang:1.25

WORKDIR /usr/src/attendance-api-gofiber

RUN go install github.com/air-verse/air@latest

COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

EXPOSE 3000