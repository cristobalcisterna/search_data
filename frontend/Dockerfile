FROM node:18.6.0 as build-stage
WORKDIR /nuxtapp
COPY . .
RUN npm install
RUN npm run build

FROM node:18.6.0 as prod-stage
WORKDIR /nuxtapp
COPY --from=build-stage /nuxtapp/.output/ ./.output/
COPY package*.json ./
RUN npm install --production --prefer-offline --pure-lockfile --non-interactive
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]