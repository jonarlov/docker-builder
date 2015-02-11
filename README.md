#Docker builder - dobu

Dobu is short for **Do**cker **bu**ilder

Dobu will build and tag your Docker images recursively from any Docker image in a layered structure.

All you need to add is a `dobu.yml` beside your Dockerfile that defines the `parent` and `dockertag` values.

`parent` - relative path to the parent Dockerfile folder  
`dockertag` - the tag to use when executing `docker build -t dockertag path`

Now you can automate your Docker images build procedure, and the best part is that you can start building from the bottom of the chain. Quite the opposite of what you need to do today.
