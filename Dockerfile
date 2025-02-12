FROM golang

WORKDIR usr/src/app

COPY . ./

RUN go mod download

RUN go build -v -o docker_project_api ./cmd
RUN go build -v -o docker_project_api_migrate ./internal/db/migration

EXPOSE 8080

CMD ["./docker_project_api"]