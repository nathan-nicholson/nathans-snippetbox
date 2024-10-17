# nathans-snippetbox

Working through the book ["Let's Go"](https://lets-go.alexedwards.net/) by Alex Edwards. This is a simple web application that allows users to create, read, update, and delete snippets of text. I'm riffing on the book's project with the intent of making it a bit more robust and production-ready as an exercise.

## Getting Started

This project requires several tools. If you have a Mac, Homebrew is your friend. If not, well... you're on your own. This assumes you have

- [Go](https://golang.org/)
  - Language of choice for the project since, you know, it's what the book is about.
- [Docker](https://www.docker.com/)
  - Containerization and whale puns are the name of the game.
  - [Buildx](https://docs.docker.com/buildx/working-with-buildx/)
    - A script is included to download and install buildx for Macs. (scripts/install-buildx.sh)
    - Should largely work, but will need to change the OS referenced in the script if you're not on a Mac.
    - I think included with Docker Desktop, but I'm not sure since I don't use it.
  - A docker runtime running on your machine
    - [Colima](https://github.com/abiosoft/colima) is what I'm using.
- [Kubernetes](https://kubernetes.io/) and [kubectl](https://kubernetes.io/docs/tasks/tools/)
  - Kubernetes is a whole thing. Good luck! :joy:
- [Kind](https://kind.sigs.k8s.io/)
  - My favorite way to run a local Kubernetes cluster.
  - Makes it easy to spin up and tear down clusters on a whim.
- [Helm](https://helm.sh/)
  - Used for defining the resources that will be deployed to the Kubernetes cluster.
- Some kind of text editor
  - I generally use [VS Code](https://code.visualstudio.com/).
  - I assume you have a favorite, so use that.
- Some kind of terminal
  - I use [iTerm2](https://iterm2.com/) + [Oh My Zsh](https://ohmyz.sh/).

### Potential Gotchas

- If you're using Colima, make sure you're setting the DOCKER_HOST environment variable to the correct value. This is probably `unix:///$HOME/.colima/docker.sock`.
- If your deployment isn't updating, make sure you're loading your images into the cluster.
  - This is why I've included `scripts/deploy.sh` so I don't have to remember this stuff!

## Kubernetes

**TODO:**

- [ ] Include an alternative setup that uses a Gateway API instead of the NGINX Ingress Controller.
- [ ] Add observability with Prometheus and Grafana
- [ ] Add logging with Loki
- [ ] Better secrets management
- [ ] Seed the database with some initial data on setup

### Setup

This project uses `kind` to create a local Kubernetes cluster. The included config creates a 4-node cluster with 1 control plane and 3 worker nodes. This is to provide a playground that at least pretends to not be a toy cluster. The control plane has additional ports exposed for ingress in order to make working with the application more natural.

You can do an initial setup of the cluster with the `scripts/setup-cluster.sh` script. This will do the following:

1. Create a new kind cluster
2. Install the NGINX Ingress Controller
3. Build the `snippetbox` image
4. Install the `snippetbox` chart using Helm ([see below](#helm))
5. Waits for the application to be available

If everything is hooked up correctly, you should be able to access the application at `http://localhost`. (Yay, ingress!)

### Tear Down

You can tear down the cluster with the `scripts/teardown-cluster.sh` script. This will delete the kind cluster and everything on it, including database contents.

## Helm

**TODO:**

- [x] Helm chart for Kubernetes deployment of application
- [x] Integrate MySQL chart for database

### Snippetbox Chart

The `snippetbox` chart is included in the `charts` directory. This chart includes the following resources:

- Deployment
- HorizontalPodAutoscaler
- Service
- Ingress
- ConfigMap (MySQL initialization script)
- MySQL Chart
  - StatefulSet
  - Service

## Build & Deploy

One of the, call it flaws, of the "Let's Go!" book is no mention of testing/pipelines...at all. While I understand that it's emphasizing the basics of Go, IMNSHO testing is a foundational element to any language. I'll add some tests that will run as part of a build process and some acceptance tests that can be run in parallel against the deployed application. Since there is no proper environment for the application, acceptance tests will be geared towards running against the local cluster.

For more information around testing generally, I recommend [Dave Farley's Youtube channel](https://www.youtube.com/@ContinuousDelivery). Just a wealth of knowledge there. Enjoy!

**TODO:**

- [x] Dockerize the application
- [ ] Use a smaller base image for the final image.
- [ ] Create a pipeline with GitHub Actions for image and chart
- [ ] Commit tests
- [ ] Acceptance tests

### Build

You can run the docker build leveraging the `scripts/build.sh` script. This will build the image with the name `snippetbox` and tag it with `latest`.

The build script will try to assess the architecture of the machine and build the image accordingly. Far from comprehensive, but works between my machines. YMMV (Your Mileage May Vary).

It uses a multi-stage build to build the binary and copy it into a base ubuntu image along with static assets that the application uses. I went for speed of implementation over optimization.

### Deployment

You can deploy the application with the `scripts/deploy.sh` script. This will do the following:

1. Build the `snippetbox` image with any changes you've made
2. Load the image into the kind cluster
3. Upgrade the `snippetbox` chart using values from `config/helm/values.yaml`
4. Perform a rolling update of the deployment
   1. This will force the deployment to pull the new image
