FROM golang:alpine as build

RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o rusty ./main.go

FROM alpine
RUN mkdir /app 
# RUN apk add --no-cache libstdc++

EXPOSE 8080
COPY --from=build /app/rusty /app/rusty
ENTRYPOINT ["/app/rusty"]
