FROM webdevops/php-nginx:alpine

WORKDIR /app/

COPY . .

ARG PORT=80
ENV PORT=${PORT}
EXPOSE ${PORT}

ARG DATA_HOST=http://host.docker.internal:9999/
ENV DATA_HOST=${DATA_HOST}

CMD ["supervisord"]