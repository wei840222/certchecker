FROM golang:1.15-buster AS build
COPY . /src
WORKDIR /src
RUN go build -o app

FROM node:14.9.0-buster AS build-web
COPY ./web /src
WORKDIR /src
RUN yarn
RUN yarn build

FROM debian:buster
COPY --from=build /src/app /app
COPY --from=build-web /src/dist /web/dist
CMD /app