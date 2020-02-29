FROM node:lts-alpine
CMD apt-get 
WORKDIR /var/www/js

COPY package.json package-lock.json ./
RUN npm install

COPY . .

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}
ENTRYPOINT ["npm", "run", "start"]