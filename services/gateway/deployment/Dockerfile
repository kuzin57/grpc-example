FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y build-essential wget
RUN apt-get install git
RUN apt-gen install gpg

# Install Go
RUN wget https://go.dev/dl/go1.22.6.linux-arm64.tar.gz
RUN tar -C /usr/local -xzf go1.22.6.linux-arm64.tar.gz
RUN rm -f go1.22.6.linux-arm64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin"
RUN go version

RUN mkdir /gateway
WORKDIR /gateway
COPY ./ ./

RUN apt-get update
RUN go mod download

# Install OpenJDK-8
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk
RUN apt-get install -y ant
# Fix certificate issues
RUN apt-get update
RUN apt-get install ca-certificates-java
RUN apt-get clean
RUN update-ca-certificates -f
# Setup JAVA_HOME -- useful for docker commandline
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64/
RUN export JAVA_HOME

# Install ScyllaDBManager


RUN go build ./services/gateway/cmd/main/*
CMD ["./main", "--port", "9090"]
