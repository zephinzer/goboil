# Golang Build Tooling Boilerplate
Opinionated build tooling boilerplate for developing Golang applications with.

# Scope
## Non-Technical (Value Propositions)
- Uses Docker images to eliminate development machine differences
- Work on Golang applications outside of the `GOPATH`
- Enables fast feedback cycle through live-reloading of app and tests
- Sticks to the native Go build tooling for compilation
- Integrates with common CI providers such as Travis and GitLab
- Integrates with common CI tools such as CodeClimate and SonarQube

## Technical
- [x] Start application
- [x] Start application with live-reloading
- [x] Dependency management
- [x] Environment configuration
- [x] Test application with coverage
- [x] Test application with live-reloading
- [x] Compile binary
- [x] Versioning capabilities
- [x] Create production Docker image bundle
- [ ] Travis integration
- [ ] GitLab integration
- [ ] CodeClimate integration
- [ ] SonarQube integration

# Tooling Included
- [Make for running common operations](https://www.gnu.org/software/make/)
- [Realize for application live-reloading](https://github.com/oxequa/realize)
- [Dep for dependency management](https://github.com/golang/dep)
- [Auto-run.py from GoConvey for test live-reloading](https://github.com/smartystreets/goconvey/wiki/Auto-test)

# Software Development Lifecycle

## Dependency Management

For all `.local` postfixed commands, a best-attempt will be tried by symlinking the currect directory into your host machine's `${GOPATH}/src` and using that symlinked directory to run `dep`.

### Initialisation
Run `make dep ARGS=init` to initialize the dependencies *(you likely will not need this if you are seeding a project by cloning this repository)*.

> To run `dep` on your hostmachine, use `make dep.local ARGS=init`. You will need Dep installed and your directory needs to be within a valid `${GOPATH}/src` to do this.

### Adding a Dependency
Run `make dep ARGS="ensure -add ${DEPENDENCIES}"` to add dependencies where `${DEPENDENCIES}` is a space-separated list of dependencies you wish to add.

> To run `dep` on your hostmachine, use `make dep.local ARGS="..."`. You will need Dep installed and your directory needs to be within a valid `${GOPATH}/src` to do this.


## Development

### Getting Started in Development

Copy `./sample.properties` into `./Makefile.properties` and set your required values there. These configurations are mostly for production image releasing so you may not need to change anything.

Run `make start` to start the application in development with live-reload.

> To run the application on your host machine, use `make start.local`. You will need Go installed and a valid `GOPATH` set to do this.

To debug the application with a shell, run `make shell` to create a shell logged in as `root` into the running development container.

### Configuring the Environment

Modify the `./.env` file to change the environment variables. A `.env` file looks like:

```
ENV_VAR_1="value 1"
ENV_VAR_2=value2
# ....
```

The above will add the environment variables `ENV_VAR_1` and `ENV_VAR_2` to your Docker image.

## Testing

Append a `.local` to run the below on your host machine (`test.local`, `testc.local`, `testw.local`). Note that this may or may not work depending on what you have available on your host machine.

### Standalone Tests

Run `make test` to run the tests once.

### Standalone Tests with Live-Reload

Run `make testw` to run the tests in automated live-reload mode.

> Requires `python` to be installed on your machine. The script is at `./.scripts/auto-run.py` courtesy of GoConvey.

### Standalone Tests with Coverage

Run `make testc` to run the tests once and output the coverage.

## Building

Run `make build` to create the binary.

To use your host machine's Go installation, run `make build.local`

## Releasing

To get the latest version of the application, use `make version.get`.

To bump the **patch** version, use `make version.bump`

To bump the **minor** version, use `make version.bump BUMP=minor`

To bump the **major** version, use `make version.bump BUMP=major`

## Containerisation

> Note: to ensure you publish to the correct repository, confirm that your settings in `./Makefile.properties` is correct.

### Creating the Image

To create a production image, use `make dkbuild`. An image should be created named with your directory name and tagged with `:latest`.

To create a development image, use `make dkbuild.dev`. An image should be created named with your directory name and tagged with `:dev-latest`.

### Publishing the Image

To publish the production image, use `make dkpublish`.

To publish the development image, use `make dkpublish.dev`.

# Methodology
## Dockerfile
The Dockerfile is split into three build stages:

1. **development**: in this stage, development dependencies and convenience tools are available. When built, this image contains the source code only without binaries. This image works for files being mounted onto the directory at `/go/src/${PROJECT_NAME}` where `${PROJECT_NAME}` will default to the name of your host directory. This eliminates the need for a valid `GOPATH` on the host machine.
1. **compile**: this stage includes everything from **development** and includes the binary. Useful for debugging the production environment.
1. **production**: this stage contains only production level stuff like your binary. No development convenience tools or dependencies are here and permissions are restricted.