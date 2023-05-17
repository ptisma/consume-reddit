FROM redis/redis-stack:6.2.6-v7

WORKDIR /usr/local/db

COPY redis.conf ./redis.conf

ENV PORT=6379
EXPOSE 6379

CMD [ "redis-stack-server", "redis.conf" ]