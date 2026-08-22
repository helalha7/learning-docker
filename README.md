# learning-docker

Hands-on Docker and Kubernetes practice repository based on the Udemy course **Docker & Kubernetes: The Practical Guide**.

This repository contains small examples and exercises I build while learning Docker concepts and containerizing applications in different languages.

## Course

Udemy: **Docker & Kubernetes: The Practical Guide**

https://www.udemy.com/course/docker-kubernetes-the-practical-guide/

## What I'm Learning

So far, this repository covers:

* Docker images and containers
* Running pre-built images
* Building custom images with Dockerfiles
* `FROM`, `WORKDIR`, `COPY`, `RUN`, `CMD`, and `EXPOSE`
* Interactive containers with `-i` and `-t`
* Image layers and build caching
* Containerizing simple Node.js applications
* Containerizing Python applications
* Containerizing Go applications

More topics will be added as I continue through the course.

## Repository Structure

```text
learning-docker/
├── first-demo-starting-setup/
├── go-demo-server/
├── nodejs-app-starting-setup/
└── python-app-starting-setup/
```

Each directory contains a small project used to practice a specific Docker concept.

## Example

A simple Go Dockerfile:

```dockerfile
FROM golang

WORKDIR /app

COPY . .

RUN go build -o server .

CMD ["./server"]
```

Build the image:

```bash
docker build -t go-demo-server .
```

Run it:

```bash
docker run go-demo-server
```

## Goal

The goal of this repository is not to build one application, but to document and practice Docker concepts through small hands-on examples.

Eventually, I plan to cover:

* Volumes and persistent data
* Docker networking
* Environment variables
* Docker Compose
* Multi-stage builds
* Registries and image publishing
* Docker deployment
* Kubernetes basics
