## Distributed File Storage Project

# Command to run Postgres
`docker run --name some-postgres -e POSTGRES_PASSWORD=goauthentication -p 5432:5432 -d postgres`

# Command for port-forward in K8s
* `kubectl port-forward service/nextjs-client-service 3000:3000`
* `kubectl port-forward <pod-name> 8000:8000`