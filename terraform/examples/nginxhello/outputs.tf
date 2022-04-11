output "port" {
  value = resource.docker_container.nginx.ports[0].external
}