#Docker builder - Dobu

Dobu is short for **Do**cker **bu**ilder

Dobu is a small, lightweight application that will build and tag your Docker images recursively from any Docker image in a layered structure.

---

![Screenshot](doc/screenshots/manual-docker-build.png)

###Building Docker images with Dobu
```
$> cd ~/src/docker/app1/tomcat-8.0.18
$> dobu build
```

*This will build all parent images and the orby/app1-tomcat:8.0.18 image for you*

All you need, to start using Dobu, is to add a `dobu.yml` beside your Dockerfiles that defines the path to the `parent` and `dockertag` value to use when building the image.

```yaml
---
# relative path to the parent Dockerfile folder
parent: ../
# the tag to use when executing "docker build -t dockertag path"
dockertag: orby/app1-tomcat:8.0.18
```

Now you can automate your Docker images build procedure, and the best part is that you can start building from the bottom of the chain. Quite the opposite way of what you can do today.

---
## Command line options
```
$> dobu
usage: dobu [<flags>] <command> [<flags>] [<args> ...]

Dobu is a recursive Docker image builder.

Flags:
  --help     Show help.
  --version  Show application version.
  -w, --working-directory="."  
             Change working directory
  -f, --file="dobu.yml"  
             Alternate dobu.yml filename

Commands:
  help [<command>]
    Show help for a command.

  list
    List images in the build chain

  build
    Build images in the build chain recursivly

  stop [<flags>]
    Stop all running containers by sending SIGTERM and then SIGKILL after a grace period
```

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

Add/fix/test/refactor code and create a pull request:  
https://help.github.com/articles/using-pull-requests/
