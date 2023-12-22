FROM node:12

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

VOLUME /usr/src/app/data   # Point de montage pour les données du compteur

EXPOSE 8080

CMD [ "npm", "start" ]

