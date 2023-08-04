FROM node:20.5.0-alpine3.17

WORKDIR /var/src
COPY . /var/src

RUN yarn install --frozen-lockfile

# live reload disabled because nginx-gateway-controller currently doesn't support websockets.
# Re-enable when https://github.com/nginxinc/nginx-kubernetes-gateway/issues/315 is fixed.
ENTRYPOINT ["yarn", "ng", "serve", "--live-reload", "false", "--host", "0.0.0.0", "--port", "80", "--disable-host-check"]