FROM golang:1.21.5-alpine3.19
WORKDIR /test
COPY . /test
RUN go build /test
EXPOSE 8081
ENTRYPOINT [ "./weather-tracker" ]