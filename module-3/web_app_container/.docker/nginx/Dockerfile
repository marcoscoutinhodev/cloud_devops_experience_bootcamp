FROM alpine:latest

RUN apk upgrade && \
        apk add --no-cache nginx

COPY .docker/nginx/nginx.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off;"]

