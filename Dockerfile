FROM node:latest as nodebuild
WORKDIR /node/open-bar
COPY package*.json ./
RUN npm install 
RUN npm install -g rollup 
COPY . .
RUN npm run build


FROM golang:1.7.3 as gobuild
WORKDIR /go/src/github.com/e-kot-unamur/open-bar
COPY main.go data.go server.go ./
RUN go get -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest as server
WORKDIR /app/ 
COPY --from=nodebuild /node/open-bar .
COPY --from=gobuild /go/src/github.com/e-kot-unamur/open-bar .
EXPOSE 8000
CMD ["./app"]