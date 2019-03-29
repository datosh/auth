# Auth

# Ingress / Traefik

Traefik can easily be installed using the helm chart:

```
helm install stable/traefik --set dashboard.enabled=true,dashboard.domain=dashboard.localhost
```

I wen't for manual configuration using my own K8s files.
These are based on [Blog Articles][TraefikSetup], to get a better
understanding of the inner workings.


# Open Points

* How to automatically write `mkcert` root CA in Treafik / K8s Secret?
* Switch from Keycloak to Hydra & something like AuthBoss

# OAuth2 

[Pluralsight Course][OAuthPluralsight] covers basics of OAuth, such as the
different grant flows.
* Everyone in possesion of the Bearer Token is able to confsume the API
    * OAuth2 extension for Client Certificates binds the Bearer Token to
    the Client Certificate
* Authorization Flow
    * Client needs to be able to keep a secret, i.e., backend service
* Implicit Flow
    * Javascript clients cannot keep a secret
    * Less secure than Authorization Flow
        * PKCE & OpenIDConnect help with some problems

OAuth2 is specified in [RFC 6749][OAuth2RFC], and gives additional security
considerations in [RFC 6819][OAuth2RFCThreat].

## Keycloak

One implementation example

Standard Endpoints
https://stackoverflow.com/questions/28658735/what-are-keycloaks-oauth2-openid-connect-endpoints

OAuth2 support for Client Certificates only recently added to spec
https://connect2id.com/blog/connect2id-server-6-13
More details
https://connect2id.com/blog/mutual-tls-client-auth
spec
https://tools.ietf.org/id/draft-ietf-oauth-mtls-07.html


mTLS explained
https://medium.com/@technospace/mutual-tls-for-oauth-client-authentication-cdd595d4dcac

<!-- Bibliography-->

[TraefikSetup]: https://medium.com/@geraldcroes/kubernetes-traefik-101-when-simplicity-matters-957eeede2cf8

[OAuthPluralsight]: https://app.pluralsight.com/player?course=oauth-2-getting-started
[OAuth2RFC]: https://tools.ietf.org/html/rfc6749
[OAuth2RFCThreat]: https://tools.ietf.org/html/rfc6819