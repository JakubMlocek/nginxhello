variable "docker_image" {
  description = "Value of the Docker image"
  type        = string
  default     = "jakmlo/nginxhello"
}

variable "port" {
  description = "Value of the output port for hello website"
  type        = number
  default     = 8081
}