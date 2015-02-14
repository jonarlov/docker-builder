#Docker builder - Dobu

Dobu is short for **Do**cker **bu**ilder.

Dobu is your swiss knife for your Docker development and test environment.

Dobu is a small, lightweight application that can build, tag, stop and delete your Docker images.

>One of Dobu's greatest feature is that it enables you to build all the parent Docker images in a layered structure.  

>*#yourfutureself #dobu*

##Commands

- `dobu list` - List all images in a build chain  
- `dobu build` - Build all images in a build chain recursivly  
- `dobu stop` - Stop all running Docker containers by sending SIGTERM and then SIGKILL after a grace period  
- `dobu delete containers` - Delete all Docker containers
- `dobu delete images` - Delete all Docker images  
- `dobu delete all` - Delete all Docker containers and Docker images
- `dobu --help` - Show all available commands and arguments

##How to create and build a Dobu build chain

Say you have your Dockerfiles organized in a layered structure like this:

```
test
├── app1-images
│   ├── backend
│   │   └── Dockerfile
│   ├── dbms
│   │   └── Dockerfile
│   └── frontend
│       └── Dockerfile
└── base-images
    ├── debian-wheezy
    │   ├── Dockerfile
    │   ├── apache-2-4
    │   │   └── Dockerfile
    │   ├── apache-2.2
    │   │   ├── Dockerfile
    │   │   └── php-5
    │   │       ├── Dockerfile
    │   │       └── wordpress-latest
    │   │           └── Dockerfile
    │   ├── java-7
    │   │   ├── Dockerfile
    │   │   └── tomcat-7.0.59
    │   │       └── Dockerfile
    │   ├── java-8
    │   │   ├── Dockerfile
    │   │   └── tomcat-8.0.18
    │   │       └── Dockerfile
    │   └── nginx-latest
    │       ├── Dockerfile
    │       └── php-5
    │           ├── Dockerfile
    │           └── wordpress-latest
    │               └── Dockerfile
    └── ubuntu-trusty
        ├── Dockerfile
        ├── mariadb-10
        │   └── Dockerfile
        └── mysql-5.5
            └── Dockerfile
```

With Dobu, you can now build your app1 frontend image like this:
```
$> cd test/app1-images/frontend
$> dobu build
```

*This will build all parent images and the app1/frontend:latest image for you with one command*

You can also execute `dobu list -w test/app1-images/frontend` to see which images that will be built and the name the images will be tagged with in Docker.
```
$> dobu list -w test/app1-images/frontend/
These images will be built:
Image: company/debian:7 @ test/base-images/debian-wheezy
Image: company/apache:2.2 @ test/base-images/debian-wheezy/apache-2.2
Image: company/apache22-php:5 @ test/base-images/debian-wheezy/apache-2.2/php-5
Image: company/apache22-php5-wordpress:latest @ test/base-images/debian-wheezy/apache-2.2/php-5/wordpress-latest
Image: app1/frontend:latest @ test/app1-images/frontend
```
*How cool is that :-)*

All you need to do is to add a `dobu.yml` file next to your Dockerfiles in your build chain, that defines the path to the `parent` Dockerfile and the `dockertag` value to use when building the image.

```yaml
---
# relative path to the parent Dockerfile folder
parent: ../
# the tag to use when executing "docker build -t dockertag path"
dockertag: orby/app1-tomcat:8.0.18
```
*You can have a look at the [test](test/) folder to see all the files.*

Now you can automate your Docker images build procedure, and the best part is that you can start building from the bottom of the chain. Quite the opposite of what you can do today.

##Build instructions

Fork Dobu from GitHub  
https://github.com/orby/docker-builder/fork

Clone Dobu from GitHub
```
cd $GOPATH/src/github.com/username/
git clone git@github.com:username/docker-builder.git
```

Execute Dobu from source
```
go run $GOPATH/src/github.com/username/docker-builder/dobu/dobu.go
```

Install binary
```
go install github.com/username/docker-builder/dobu
```

Delete binary
```
go clean -i github.com/username/docker-builder/dobu
```

Add/fix/test/refactor code and create a pull request  
https://help.github.com/articles/using-pull-requests/
