FROM node:latest as nodebuild
WORKDIR /node/open-bar
COPY client/ ./
RUN npm install 
RUN npm install -g rollup 
RUN npm run build


FROM golang:1.7.3 as gobuild
WORKDIR /go/src/github.com/e-kot-unamur/open-bar
COPY server/ ./
RUN go get -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest as server
WORKDIR /app/ 
COPY --from=nodebuild /node/open-bar ./client
COPY --from=gobuild /go/src/github.com/e-kot-unamur/open-bar ./server
EXPOSE 8000
CMD ["./server/app"]