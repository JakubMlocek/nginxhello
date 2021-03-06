# jakmlo/nginxhello
## Terraform module for running a dockerized application on machine.
* Simple Dockerfile helo world project serving static HTML page using nginx. Docker image of this project: [jakmlo/nginxhello](https://hub.docker.com/repository/docker/jakmlo/nginxhello). 
* Project has a makefile which combines building and pushing an Docker image.
* Terraform module allows running container on local machine. It accepts as input variable of a Docker image and port on which website should be accesible on localhost. Default image is jakmlo/nginxhello and default port is set to 8081.
* Tests written in GoLang checks whether the server is configred correctly.


## Usage Examples
* Firstly we need to initialize terraform directory: `terraform init`
* To run module with default setting: `terraform apply`
* To run module with custom Docker image: `terraform apply -var "docker_image=NAME_OF_DESIRE_DOCKER_IMAGE"`
* To run module with custom Port: `terraform apply -var "port=VALUE_OF_DESIRE_PORT"`
* To run module with both custom Docker image and custom port:
`terraform apply -var "docker_image=NAME_OF_DESIRE_DOCKER_IMAGE"  -var "port=VALUE_OF_DESIRE_PORT"`

