FROM php:7.4.3-apache

WORKDIR /var/www/html

COPY . .

RUN a2enmod rewrite

ARG PORT=80
ENV PORT=${PORT}
EXPOSE ${PORT}

ARG DATA_HOST=http://192.168.68.103:9999/
ENV DATA_HOST=${DATA_HOST}

CMD ["apachectl", "-D", "FOREGROUND"]