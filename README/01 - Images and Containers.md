
Images and containers are the two core concepts in Docker.

The simplest mental model is:

```text
Image
  ↓
docker run
  ↓
Container
```

An **image** is the reusable package.

A **container** is a running instance created from that image.

---

## Docker Images

A Docker image is a read-only template used to create containers.

An image can contain:

- operating system files
    
- runtimes like Go, Node, or Python
    
- application files
    
- dependencies
    
- configuration
    
- default startup commands
    

Examples of existing images:

```text
golang
node
python
mysql
ubuntu
```

These can be downloaded from a registry such as Docker Hub.

---

## Pulling an Image

To explicitly download an image:

```bash
docker pull node
```

Docker stores the image locally.

You can then create containers from it.

You often don't need to run `docker pull` manually.

If you do:

```bash
docker run node
```

and the image does not exist locally, Docker will try to pull it automatically.

The flow is:

```text
docker run node
      ↓
Is node image available locally?
      ↓
     no
      ↓
pull from Docker Hub
      ↓
create container
      ↓
start container
```

---

## Listing Images

To see local images:

```bash
docker image ls
```

or:

```bash
docker images
```

Example:

```text
REPOSITORY   TAG       IMAGE ID
node         latest    abc123
ubuntu       24.04     def456
```

### Untagged Images

Sometimes you may see:

```text
<none>
```

or:

```text
<untagged>
```

These are images that currently do not have a repository/tag pointing to them.

For example, rebuilding the same image tag can leave an older image untagged.

To show more images:

```bash
docker image ls -a
```

---

# Image Names and Tags

An image is commonly identified using:

```text
repository:tag
```

Example:

```text
node:22
```

Here:

```text
node
```

is the repository/image name.

```text
22
```

is the tag.

If no tag is specified:

```bash
docker run node
```

Docker normally interprets it as:

```text
node:latest
```

Tags can be used for versions:

```text
my-api:v1
my-api:v2
my-api:1.5.0
```

---

# Building My Own Image

Images do not have to come from Docker Hub.

I can build one myself using a Dockerfile.

Example:

```bash
docker build -t my-go-server .
```

This creates an image named:

```text
my-go-server:latest
```

Then:

```bash
docker run my-go-server
```

creates a container from it.

Related:

- [[02 - Dockerfiles]]
    

---

# Docker Containers

A container is an instance of an image.

For example:

```text
             node image
            /     |     \
           /      |      \
     container1 container2 container3
```

All three containers came from the same image.

But each container is independent.

---

## Image vs Container

A useful analogy:

```text
Image     = blueprint
Container = thing created from blueprint
```

Another analogy:

```text
Image     = class
Container = object
```

The analogy is not exact, but it helps with the basic idea.

---

# Running a Container

Basic command:

```bash
docker run IMAGE
```

Example:

```bash
docker run ubuntu
```

`docker run` roughly performs:

```text
find image
    ↓
create container
    ↓
start container
```

A new container is created every time `docker run` is executed.

So:

```bash
docker run ubuntu
docker run ubuntu
docker run ubuntu
```

creates three different containers.

---

# Why Some Containers Exit Immediately

A container runs as long as its main process is running.

If the main process exits:

```text
main process exits
      ↓
container stops
```

For example:

```bash
docker run node
```

may exit quickly if Node does not have interactive input or a program to keep running.

A container is not the same thing as a virtual machine that must stay alive forever.

The container's lifetime is tied to its main process.

---

# Running Interactively

Sometimes an application needs input.

Use:

```bash
docker run -i IMAGE
```

`-i` means:

```text
interactive
```

It keeps standard input open.

For example, a Python program using:

```python
input()
```

may need:

```bash
docker run -i my-python-app
```

---

## `-t`

`-t` allocates a pseudo-terminal.

Usually interactive terminal programs use:

```bash
docker run -it IMAGE
```

Mental model:

```text
-i = keep input open
-t = give the process a terminal
```

Example:

```bash
docker run -it ubuntu bash
```

This gives an interactive shell inside the Ubuntu container.

---

# Running in the Background

By default:

```bash
docker run IMAGE
```

attaches the terminal to the container.

To run in detached mode:

```bash
docker run -d IMAGE
```

`-d` means:

```text
detached
```

The container continues running in the background.

This is common for servers and databases.

Example:

```bash
docker run -d nginx
```

---

# Container Names

Docker automatically generates a name if I do not provide one.

I can choose one:

```bash
docker run --name my-server IMAGE
```

Then I can use:

```bash
docker stop my-server
```

instead of remembering the container ID.

---

# Listing Containers

Show currently running containers:

```bash
docker ps
```

or:

```bash
docker container ls
```

Show all containers, including stopped ones:

```bash
docker ps -a
```

Example:

```text
CONTAINER ID   IMAGE      STATUS
abc123         node       Up 2 minutes
def456         python     Exited
```

---

# Stopping a Container

Stop a running container:

```bash
docker stop CONTAINER
```

Example:

```bash
docker stop my-server
```

Stopping does not automatically remove the container.

It still exists.

---

# Starting a Stopped Container

A stopped container can be started again:

```bash
docker start CONTAINER
```

This is different from:

```bash
docker run IMAGE
```

because:

```text
docker start
→ starts an existing container

docker run
→ creates a new container and starts it
```

This distinction is important.

---

# Removing Containers

Remove a stopped container:

```bash
docker rm CONTAINER
```

Force-remove a running container:

```bash
docker rm -f CONTAINER
```

Removing the container does not remove its image.

Example:

```text
Image
 ├── Container A
 └── Container B
```

Deleting Container A:

```text
Image
 └── Container B
```

The image still exists.

---

# Removing Images

Remove an image:

```bash
docker image rm IMAGE
```

or:

```bash
docker rmi IMAGE
```

Example:

```bash
docker image rm my-app
```

Docker may refuse if containers still depend on that image.

---

# Container Filesystem

When a container is created, Docker gives it a writable layer on top of the image.

Conceptually:

```text
Container writable layer
────────────────────────
Image layers
────────────────────────
Base image
```

If I change a file inside a container, I am not changing the original image.

For example:

```text
Image
  ↓
Container A
  ↓
modify /app/file.txt
```

The image remains unchanged.

If I create another container:

```text
Image
  ↓
Container B
```

Container B starts from the original image state.

---

# Containers Are Disposable

A major Docker idea is that containers should generally be easy to replace.

Instead of treating a container as:

> this machine must live forever

think:

> if this container breaks, I should be able to create another one from the image

This becomes especially important for deployment.

Important persistent data should usually not depend entirely on one container's writable layer.

Related:

- [[04 - Data and Volumes]]
    

---

# Copying Files Into a Container

It is possible to modify an individual container manually:

```bash
docker cp file.txt CONTAINER:/app/file.txt
```

This changes that specific container.

It does **not** modify the image.

So:

```text
Image
 ├── Container A ← file copied here
 └── Container B ← unchanged
```

For application files that should always exist, they should normally be added while building the image using a Dockerfile instead.

---

# Port Publishing

A server can listen on a port inside the container.

Example:

```text
Go server inside container
→ port 8080
```

That port is not automatically exposed to my Mac.

I can publish it:

```bash
docker run -p 8080:8080 my-server
```

The format is:

```text
HOST_PORT:CONTAINER_PORT
```

Example:

```text
My Mac                 Container

localhost:8080  ───→   :8080
```

This lets traffic from the host reach the application inside the container.

Related:

- [[05 - Networking]]
    

---

# Image Sharing

Images can be shared through registries such as Docker Hub.

Typical flow:

```text
build image
    ↓
tag image
    ↓
push image
    ↓
Docker Hub
    ↓
another machine
    ↓
pull image
    ↓
run container
```

This means another machine does not necessarily need the source-code repository just to run the application.

Related:

- [[07 - Registries and Docker Hub]]
    

---

# Important Commands

## Images

```bash
docker pull IMAGE
docker image ls
docker image ls -a
docker image rm IMAGE
```

## Containers

```bash
docker run IMAGE
docker run -it IMAGE
docker run -d IMAGE

docker ps
docker ps -a

docker stop CONTAINER
docker start CONTAINER
docker rm CONTAINER
```

## Named Container

```bash
docker run --name my-container IMAGE
```

## Port Publishing

```bash
docker run -p HOST_PORT:CONTAINER_PORT IMAGE
```

---

# Important Distinctions

## `docker run` vs `docker start`

```text
docker run
→ create NEW container
→ start it

docker start
→ start EXISTING container
```

## Image vs Container

```text
Image
→ reusable template

Container
→ instance of image
```

## Stop vs Remove

```text
stop
→ container still exists

remove
→ container no longer exists
```

## Container Change vs Image Change

```text
change file inside container
→ image does not change
```

To change the reusable image, rebuild it.

---

## Related

- [[00 - Docker Overview]]
    
- [[02 - Dockerfiles]]
    
- [[03 - Image Layers and Caching]]
    
- [[04 - Data and Volumes]]
    
- [[05 - Networking]]
    
- [[07 - Registries and Docker Hub]]