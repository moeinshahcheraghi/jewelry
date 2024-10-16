FROM node:18-alpine as build
WORKDIR /app
COPY package*.json ./
ARG http_proxy=http://45.94.254.18:43129
ARG https_proxy=http://45.94.254.18:43129
RUN npm install
COPY . .
RUN npm run build
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]

