#!/bin/bash

if [[ $1 == '--watch-dark' ]]; then
  tailwindcss -c ./assets/tailwind.config.js -i ./assets/css/tailwinds-dark.css -o ./assets/css/tailwinds-dark-pack.css --watch
elif [[ $1 == '--watch-light' ]]; then
  tailwindcss -c ./assets/tailwind.config.js -i ./assets/css/tailwinds-light.css -o ./assets/css/tailwinds-light-pack.css --watch
else
  tailwindcss -c ./assets/tailwind.config.js -i ./assets/css/tailwinds-dark.css -o ./assets/css/tailwinds-dark-pack.css
  tailwindcss -c ./assets/tailwind.config.js -i ./assets/css/tailwinds-light.css -o ./assets/css/tailwinds-light-pack.css
fi