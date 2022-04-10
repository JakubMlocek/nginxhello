# jakmlo/nginxhello
## Terraform module for running a dockerized application on machine.
* Simple Dockerfile helo world project serving static HTML page using nginx. Docker image of this project: [jakmlo/nginxhello](https://hub.docker.com/repository/docker/jakmlo/nginxhello).
* Terraform module allowing running container on local machine. Accepts as input variable of a Docker image and port on which website should be accesible on localhost. Default image is jakmlo/nginxhello and default port is set to 8081.


## Usage Examples


