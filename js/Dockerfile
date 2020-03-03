FROM node:lts-alpine
CMD apt-get 
WORKDIR /var/www/js

COPY package.json package-lock.json ./
RUN npm install

COPY . .

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

ARG DATA_HOST=http://host.docker.internal:9999/
ENV DATA_HOST=${DATA_HOST}

ENTRYPOINT ["npm", "run", "start"]