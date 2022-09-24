FROM node:16.17.0-alpine
WORKDIR /usr/src/app
COPY . .
RUN npm install
RUN npm run build
RUN npm prune --omit=dev
CMD ["npm", "start"]
EXPOSE 3000
