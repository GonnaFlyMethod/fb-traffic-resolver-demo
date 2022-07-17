FROM golang:1.18 as backend
WORKDIR /backend

COPY ./backend /backend
RUN go build .
CMD ["./backend"]

FROM node:14.16.0 as frontend
WORKDIR /frontend

COPY ./frontend/package.json ./frontend/yarn.lock ./
COPY ./frontend .
RUN yarn
CMD ["yarn", "start"]
