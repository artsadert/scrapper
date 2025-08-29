FROM alpine:latest

#golang
RUN apk add go


COPY . ./
RUN go mod download

RUN go build -o scrapper ./cmd/

# nginx
RUN apk add nginx
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 8008

CMD ["nginx", "-g", "daemon off;"] 
