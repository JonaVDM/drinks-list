FROM node:24-alpine AS frontend
WORKDIR /app

COPY web/package*.json ./
RUN npm install
COPY web/ .
RUN npm run build

FROM golang:1.25 AS build
WORKDIR /go/src/app

COPY . .
RUN go mod download
COPY --from=frontend /app/.output web/.output
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/app /
CMD ["/app", "serve", "--http=0.0.0.0:8090"]
