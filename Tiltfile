# Traefik
k8s_yaml('traefik/deployment.yml')
k8s_yaml('traefik/service.yml')
k8s_yaml('traefik/config.yml')

# Traefik Dashboard
k8s_yaml('traefik/dashboard/ingress.yml')
k8s_yaml('traefik/dashboard/service.yml')

# Hello Service
k8s_yaml('helloService/k8s/deployment.yml')
k8s_yaml('helloService/k8s/service.yml')
k8s_yaml('helloService/k8s/ingress.yml')
docker_build('helloservice', 'HelloService')

# Keycloak
k8s_yaml('keycloak/deployment.yml')
k8s_yaml('keycloak/service.yml')
k8s_yaml('keycloak/ingress.yml')

# Whoami
k8s_yaml('whoami/deployment.yml')
k8s_yaml('whoami/service.yml')
k8s_yaml('whoami/ingress.yml')