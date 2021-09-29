FROM public.ecr.aws/bitnami/golang:1.17 as builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o web-server .

FROM gcr.io/distroless/base-debian10:latest

WORKDIR /app
COPY --from=builder /app/web-server /app/web-server

EXPOSE 8080
ENTRYPOINT ["/app/web-server"]