FROM golang:1.18 as backend_build
WORKDIR /build

COPY ./backend .
RUN CGO_ENABLED=0 go build -o app main.go

FROM alpine:3.16 as backend_app
WORKDIR /backend

COPY --from=backend_build /build/app .
ENTRYPOINT ["./app"]

FROM node:14.16.0 as frontend_build
WORKDIR /frontend

COPY ./frontend .
RUN yarn && yarn run build

FROM gonnaflymethod/fb-traffic-resolver:1.0.0 as resolver
# Also you can pull the container with resolver from GitHub container registry:
#FROM ghcr.io/gonnaflymethod/fb-traffic-resolver:1.0.0 as resolver
COPY --from=frontend_build /frontend/build ./build
