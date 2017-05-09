# TODOMVC API
#
# VERSION 0.0.1

FROM openjdk:8

MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Artifact definition
ARG ARTIFACT_NAME=todomvc-api
ARG ARTIFACT_VERSION=0.0.1-SNAPSHOT

# Artifact default location (Nexus)
ARG ARTIFACT_LOCATION="./build/libs/$ARTIFACT_NAME-$ARTIFACT_VERSION.jar"

# Define env used by run scripts
ENV ARTIFACT_NAME $ARTIFACT_NAME
ENV ARTIFACT_VERSION $ARTIFACT_VERSION

# Add jar
ADD $ARTIFACT_LOCATION /opt/app/$ARTIFACT_NAME-$ARTIFACT_VERSION.jar

# Link jar
RUN ln -s /opt/app/${ARTIFACT_NAME}-${ARTIFACT_VERSION}.jar /opt/app/api.jar

# Exposed port(s)
EXPOSE 8080

# Define entrypoint
ENTRYPOINT ["java","-jar","/opt/app/api.jar"]

