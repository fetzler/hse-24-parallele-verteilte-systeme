FROM golang:1.22
WORKDIR /usr/src/app
COPY backend/go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./backend/...
CMD ["app"]