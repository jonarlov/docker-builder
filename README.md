#Docker builder - dobu

Dobu is short for **Do**cker **bu**ilder

Dobu is a small, lightweight application that will build and tag your Docker images recursively from any Docker image in a layered structure.

---

![Screenshot](doc/screenshots/manual-docker-build.png)

All you need to add is a `dobu.yml` beside your Dockerfile that defines the path to the `parent` and `dockertag` value to use when building the image.

`parent` - relative path to the parent Dockerfile folder  
`dockertag` - the tag to use when executing `docker build -t dockertag path`

Now you can automate your Docker images build procedure, and the best part is that you can start building from the bottom of the chain. Quite the opposite way of what you can do today.
