## GoLang alpine docker problem

Here is an example of a golang sqlx connection issue with GoLang Alpine docker base image

When connecting with `sqlx.Connect("postgres", url)`, the statement get stuck while using *alpine:latest*.
Comparing this with just running the same program in *golang:latest* docker base image, there is no problem.

Why is this?


### How to reproduce this

#### With rake

For those of you having `rake` installed, you can simply run `rake`. This will build the images and then run the commands to reproduce the problem.


#### Without rake

 * Build the docker images:
   * `docker build -f Dockerfile.build -t mamarcus.org/project:build .`
   * `docker build -f Dockerfile.alpine -t mamarcus.org/project:alpine .`
 * Start postgres `docker-compose up -d postgres`
 * Make sure postgres is up and responding to commands
 * Run `docker-compose up build` (This will start the full golang image which can connect properly)
 * Run `docker-compose up alpine` (This will start the slim golang alpine image that does not connect...)

### Example output

The expected (but not wanted) output when running `rake run`:

```bash
docker-compose up -d postgres
Creating o2_postgres_1 ...
Creating o2_postgres_1 ... done
Checking connection to postgres ('psql' must be installed)...........
Connected

Running service on build image. This should work fine:
docker-compose up build
o2_postgres_1 is up-to-date
Creating o2_build_1 ...
Creating o2_build_1 ... done
Attaching to o2_build_1
build_1     | Postgres addr: postgres://postgres@postgres/postgres?sslmode=disable
build_1     | Connecting successful
o2_build_1 exited with code 0

But: Running the alpine image; this will not connect nor quit/timeout for 10 min + ...
docker-compose up alpine
o2_postgres_1 is up-to-date
Creating o2_alpine_1 ...
Creating o2_alpine_1 ... done
Attaching to o2_alpine_1
alpine_1    | Postgres addr: postgres://postgres@postgres/postgres?sslmode=disable
```
And here, it stays for a long time...