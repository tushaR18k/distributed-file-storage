FROM node:22.3.0-alpine3.20 AS base

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install && npm install -D @swc/cli @swc/core

COPY ./ .

RUN chmod +x ./entrypoint.sh

EXPOSE 3000

ENV NODE_ENV=development

CMD ["./entrypoint.sh"]

