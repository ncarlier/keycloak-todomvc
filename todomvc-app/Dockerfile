#########################################
# Build stage
#########################################
FROM node:6-onbuild AS built-app

#########################################
# Final stage
#########################################
FROM nginx:latest
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Install static files
COPY --from=built-app /usr/src/app /usr/share/nginx/html

