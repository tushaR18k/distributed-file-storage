# metadata-server/Dockerfile
FROM node:14-alpine
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
EXPOSE 3002
CMD ["node", "index.js"]
