# Go-api-demo-skeleton

This is the repo to track my golang learning steps. 

* Configuration
    - default config
    - Parameters/Flags from console to Override
    - From file config.yaml 

* API
    - / -> dummy response
    - /version 
    - /healthz
    - Authenticated Endpoints (by Groups):
        - /read
        - /basic
        - /admin
    


* Middleware
    - Header (NoCache, Options, Security)
    - RequestId
    - Authentication

* Global logging

* Hard coded Users
    - admin/admin (Role Admin)
    - basic/basic (Role Basic)
    - read/read (Role Read)

# Docker

* Build image
```
docker build -t my-api .
```

* Execute
```
docker run -p8088:8088 my-api
```


