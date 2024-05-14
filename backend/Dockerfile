FROM golang:1.22-alpine3.19
WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./...

ENV POSTGRES_USER=postgres
ENV POSTGRES_HOST=db
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=todo_list
ENV POSTGRES_PORT=5432

CMD ["app"]