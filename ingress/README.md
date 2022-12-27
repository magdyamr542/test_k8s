# Testing kubernetes ingress

- There are 2 backends **service1** and **service2** which have a deployment and a service each.
- To use an ingress, an ingress-controller should be deployed. This was done using helm following [this guide](https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop). Running `helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
` should be enough.
- After starting the application, make sure to forward the port of the ingress service in order to talk to the ingress. Run `kubectl port-forward --namespace=ingress-nginx services/ingress-nginx-controller 1234:80`
- A request made to `localhost:1234/service1/` is proxied to the **service1** backend.
- A request made to `localhost:1234/service2/` is proxied to the **service2** backend.
