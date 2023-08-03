FROM node:20.5.0-alpine3.17

WORKDIR /var/src
COPY . /var/src

RUN yarn install --frozen-lockfile

ENTRYPOINT ["yarn", "ng", "serve", "--host", "0.0.0.0", "--port", "80", "--disable-host-check"]