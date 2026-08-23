Docker is a tool for packaging and running applications in isolated environments called **containers**.

The main idea is to make an application run in a predictable environment regardless of the machine it runs on.

## Core Mental Model

```text
Dockerfile
    ↓
docker build
    ↓
Image
    ↓
docker run
    ↓
Container
```

### Dockerfile

A `Dockerfile` contains instructions for building an image.

Example:

```dockerfile
FROM golang

WORKDIR /app

COPY . .

RUN go build -o server .

CMD ["./server"]
```

### Image

An image is a reusable, read-only template.

It can contain:

- operating system files
    
- runtimes like Go, Node, or Python
    
- application files
    
- dependencies
    
- configuration
    
- startup commands
    

Examples:

```text
golang
node
python
mysql
ubuntu
```

### Container

A container is a running instance of an image.

One image can create multiple containers:

```text
          Image
        /   |   \
       /    |    \
Container Container Container
```

Each container is separate, while the original image stays unchanged.

---

## Why Docker Is Useful

Docker helps solve problems such as:

### Consistent environments

Instead of every developer manually installing and configuring:

```text
MySQL
Redis
Node
Python
specific versions
```

Docker can define those environments for the project.

### Isolation

Different projects can use different versions of the same software without interfering with each other.

Example:

```text
Project A → MySQL 8.0
Project B → MySQL 8.4
```

### Easy setup

A project can eventually be started with something like:

```bash
docker compose up
```

instead of manually installing every service.

### Deployment

The same application image can be built once and then run on another machine.

```text
Developer / CI
      ↓
build image
      ↓
push image
      ↓
container registry
      ↓
server pulls image
      ↓
run container
```

---

## Docker Hub

Docker Hub is a container image registry.

It contains existing images such as:

- `golang`
    
- `node`
    
- `python`
    
- `mysql`
    
- `ubuntu`
    

For example:

```bash
docker run node
```

Docker looks for the `node` image locally.

If it does not exist, Docker can pull it from Docker Hub automatically.

Later, I can also push my own images to Docker Hub.

Related:

- [[07 - Registries and Docker Hub]]
    

---

## Docker During Development

Docker does not mean that all code must always run inside containers.

A common development setup can be:

```text
My Mac
├── Go app running normally
└── Docker
    ├── MySQL
    ├── Redis
    └── other services
```

This keeps development simple while Docker manages external services.

Later, the Go app itself can also run in a container.

---

## Docker in a Team

Docker helps developers share the same environment.

Instead of telling every teammate:

```text
Install MySQL
Install Redis
Use this exact version
Configure these ports
Create this database
```

the repository can contain Docker configuration describing the environment.

Then developers can recreate it consistently.

Docker does not replace Git.

```text
Git
→ shares source code

Docker
→ standardizes how the application and its services run
```

---

## Docker and Deployment

A Linux server can run the same images used during development.

For example:

```text
Internet
    ↓
Linux VPS
├── Go container
├── MySQL container
└── Redis container
```

The server does not necessarily need the application's Git repository.

A built image can instead be pushed to a registry and pulled by the server.

---

## Docker vs Docker Compose

Docker can manage individual containers.

Example:

```bash
docker run mysql
```

Docker Compose is used to describe and run multiple related services together.

Example:

```text
Application
├── Go API
├── MySQL
├── Redis
└── MongoDB
```

Related:

- [[06 - Docker Compose]]
    

---

## Docker vs Kubernetes

Docker is mainly concerned with:

```text
building images
running containers
managing container environments
```

Kubernetes is mainly concerned with managing containers at larger scale.

```text
Docker
→ containers

Kubernetes
→ orchestration of many containers
```

Kubernetes is not required to use Docker or to deploy normal backend applications.

---

## Main Topics

- [[01 - Images and Containers]]
    
- [[02 - Dockerfiles]]
    
- [[03 - Image Layers and Caching]]
    
- [[04 - Data and Volumes]]
    
- [[05 - Networking]]
    
- [[06 - Docker Compose]]
    
- [[07 - Registries and Docker Hub]]
    
- [[08 - Utility Containers]]
    
- [[09 - Debugging and Commands]]
    
- [[10 - Deployment]]
    

## Core Idea

The most important mental model:

```text
Source Code
    +
Dockerfile
    ↓
docker build
    ↓
Image
    ↓
docker run
    ↓
Container
```

An **image** is the reusable package.

A **container** is a running instance of that package.