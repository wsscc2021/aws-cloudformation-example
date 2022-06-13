# Build stage
FROM public.ecr.aws/docker/library/golang:1.17.9-alpine3.15 AS build

WORKDIR /source

COPY src/ ./

RUN go mod tidy \
 && go mod download \
 && go build -o ./main

# Runtime stage
FROM public.ecr.aws/docker/library/golang:1.17.9-alpine3.15

WORKDIR /app

COPY --from=build /source/main ./

RUN chmod +x ./main

ENTRYPOINT ["./main"]