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
usage: dobu [<flags>] <command>

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
    List docker images that would be build

  build
    Build docker images recursivly

```
