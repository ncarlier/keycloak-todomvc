#########################################
# Build stage
#########################################
FROM openjdk:8 AS built-api
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Artifact definition
ARG ARTIFACT_NAME=todomvc-api

# Copy suources into the container
ADD . /usr/src/$ARTIFACT_NAME

# Set working directory
WORKDIR /usr/src/$ARTIFACT_NAME

# Build the artifact
RUN ./gradlew clean build

#########################################
# Final stage
#########################################
FROM openjdk:8-jre
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Artifact definition
ARG ARTIFACT_NAME=todomvc-api
ARG ARTIFACT_VERSION=0.0.1-SNAPSHOT

# Define env used by run scripts
ENV ARTIFACT_NAME $ARTIFACT_NAME
ENV ARTIFACT_VERSION $ARTIFACT_VERSION

# Copy built artifact into the container
COPY --from=built-api /usr/src/$ARTIFACT_NAME/build/libs/$ARTIFACT_NAME-$ARTIFACT_VERSION.jar /opt/

# Link the artifact
RUN ln -s /opt/${ARTIFACT_NAME}-${ARTIFACT_VERSION}.jar /opt/app.jar

# Exposed port(s)
EXPOSE 8080

# Define entrypoint
ENTRYPOINT ["java","-jar","/opt/app.jar"]

