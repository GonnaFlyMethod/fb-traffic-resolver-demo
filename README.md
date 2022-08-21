# 🕹️ fb-traffic-resolver-demo
This is a simple demo application that implements a CRUD functionality that is related to users(Create users, Read users, Update users and Delete users).
The demo shows basic workflow with [fb-traffic-resolver](https://github.com/GonnaFlyMethod/fb-traffic-resolver)

# 🧩 Components of the demo
* Backend API running in inner docker network on port **8000**
* Resolver running on http://localhost:80 or simply http://localhost. (See [docker-compose.yml](https://github.com/GonnaFlyMethod/fb-traffic-resolver-demo/blob/main/docker-compose.yaml) and [.env file](https://github.com/GonnaFlyMethod/fb-traffic-resolver-demo/blob/main/.env)).
Resolver will give a frontend(static) files on http://localhost:80. But also you can get access to API using the same host. For instance, try HTTP GET http://localhost:80/api/users. So, frontend can simply rely on the path `/api`. [See how it's implemented in the demo](https://github.com/GonnaFlyMethod/fb-traffic-resolver-demo/blob/main/frontend/src/services/api/user.ts).

As you can see from [Dockerfile](https://github.com/GonnaFlyMethod/fb-traffic-resolver-demo/blob/main/Dockerfile#L12), all we have to do is to run API and to move built frontend to the container with resolver (to the root path).

# ℹ️ Environment variables of fb traffic resolver
* **ADDRESS_OF_API** - as the name suggets it is the environment variable that specifies an address of API. In the demo we're using docker-compose, so it's easy to reference to backend API by the container's name that is declared in docker-compose.yml. In kubernetes it can be service's name instead.
* **PING_API_ON_START** - should resolver ping backend API on start or not. Possible values (True/False)
* **RESOLVER_PORT** - the port of resolver. In the demo it is **port 80**

# 🚀 How to run the demo
Just execute:
```bash
make run
```
to shut down the demo type:
```bash
make down
```

👍 Special thanks to [RomaZherko21](https://github.com/RomaZherko21) who gave me a frontend project for the demo.
