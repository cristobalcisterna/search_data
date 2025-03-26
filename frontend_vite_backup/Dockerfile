# Etapa 1: Construcción
#FROM node:18 AS build-stage
#WORKDIR /app
#COPY package.json package-lock.json ./
#RUN npm install
#COPY . .
#RUN npm run build

# Etapa 2: Servir con Nginx
#FROM nginx:alpine AS production-stage
#COPY --from=build-stage /app/dist /usr/share/nginx/html
#EXPOSE 5173
#CMD ["nginx", "-g", "daemon off;"]

# frontend/Dockerfile

FROM node:18

WORKDIR /app
COPY . .

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]