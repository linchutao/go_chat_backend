FROM alpine
MAINTAINER linchutao <linchutaomail@163.com>
RUN echo "Asia/Shanghai" > /etc/timezone
WORKDIR /opt

COPY backend /opt/

RUN chmod +x /opt/backend

RUN apk update \
&& apk upgrade \
&& apk add --no-cache \
ca-certificates \
&& update-ca-certificates 2>/dev/null || true

EXPOSE 8080

CMD ["./backend"]





