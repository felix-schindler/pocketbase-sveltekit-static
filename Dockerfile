FROM node:lts-alpine AS sveltekit
WORKDIR /app
COPY sveltekit/package.json sveltekit/package-lock.json ./
RUN npm install
COPY sveltekit ./
RUN npm run build

FROM golang:alpine AS pocketbase
WORKDIR /app
ENV CGO_ENABLED=0
COPY pocketbase/go.mod pocketbase/go.sum ./
RUN go mod download
COPY pocketbase ./
COPY --from=sveltekit /app/build /app/pb_public
RUN go build -o pocketbase main.go

FROM alpine:latest
WORKDIR /app
RUN apk --update add ca-certificates
COPY --from=pocketbase /app/pocketbase /app/pocketbase
EXPOSE 8090
CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090"]
