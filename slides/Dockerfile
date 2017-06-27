#########################################
# Build stage
#########################################
FROM node:6-onbuild AS built-app
RUN make

#########################################
# Final stage
#########################################
FROM nginx:latest
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Install static files
COPY --from=built-app /usr/src/app/_build /usr/share/nginx/html

