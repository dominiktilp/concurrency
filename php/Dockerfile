FROM php:7.4.3-apache

WORKDIR /var/www/html

COPY . .

RUN a2enmod rewrite

ARG PORT=80
ENV PORT=${PORT}
EXPOSE ${PORT}
CMD ["apachectl", "-D", "FOREGROUND"]