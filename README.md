# MGWAAS (Microgateway As a Service)

## Running Apigee Edge Microgateway in a Docker Container 

>Official Doc. [Link](https://apigee.com/about/tags/docker)

---

Pull docker image  

```bash
docker pull gcr.io/apigee-microgateway/edgemicro:latest
```

Run docker container

```bash
docker run --env-file ./<org>-<env>-env.list  -p 8000:8000 -v <directory containing configuration>:/root/.edgemicro -d -t emgw
```

* `--env-file ./<org>-<env>-env.list` *places the environment variables from your list into the container’s environment.*
* `-p 8000:8000` *maps the port on the host machine (the first parameter) to the exposed port on the container (the second parameter).*
* `-v <directory containing configuration>:/root/.edgemicro` *mounts the directory containing your configuration on the host machine (first parameter) to the specified directory in the container (second parameter).*
* `-d` *runs in the container in detached mod*
* `-t` *specifies the tag of the image you want to run.*