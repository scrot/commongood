# syntax=docker/dockerfile:1
# https://docs.docker.com/language/golang/build-images/
# https://snyk.io/blog/containerizing-go-applications-with-docker/

# Builder image
FROM golang:1.22rc1-bookworm AS builder

RUN groupadd -r commongood && useradd --no-log-init -r -g commongood commongood

WORKDIR /app

COPY . ./
RUN go mod download &&\
  CGO_ENABLED=0 GOOS=linux go build -o=/app/commongood ./cmd/web

# Production image
FROM scratch

COPY --from=builder /app/commongood /
COPY --from=builder /etc/passwd /etc/passwd

USER commongood

ENV PORT=443
EXPOSE 443

CMD ["commongood"]
