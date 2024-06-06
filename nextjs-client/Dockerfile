FROM node:21.7.3-alpine AS base

WORKDIR /app

COPY ./nextjs-client/package.json ./
COPY ./nextjs-client/package-lock.json ./

RUN npm install

COPY ./nextjs-client .

EXPOSE 3000

ENV NODE_ENV=development

CMD [ "npm","run","dev" ]

