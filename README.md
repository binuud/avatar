# avatar

[Documentation](https://binuud.com/docs/avatar)

ai life form, physical robot with integrated conscious, memory and communication

Project status - open for alpha users

## Install

* Using go install
```
go install github.com/binuud/avatar/cmd/avatar@latest
```

* Download the avatar binary from the release folder (https://github.com/binuud/avatar/releases)

* Access UI, after running the binary, using address
```
http://localhost:9080
```

## Metrics

The ui displays system metrics for CPU, RAM and GPU usage. The GPU memory, is the memory block allocated to GPU.

### Docker Image

```
docker pull binuud/avatar:latest
```

To start the avatar container 
```
docker run --rm  -p 9090:9090 -p 9080:9080 \
  -v "[CODE ROOT]:/data" \
  binuud/avatar:latest
```

Access from UI
```
http://localhost:9080
```

## Local Models

For local models

* llama.cpp
* ollama
* mlx serve (only on macos)


## Avatar Project Collection

Avatar is a collection of the below projects

| For             | License     | Repo                                                      | 
| :---            | :----       | :---                                                      |
| Avatar          | MIT         | https://github.com/binuud/avatar                          |
| Conscious       | MIT         | https://github.com/binuud/avatar-conscious                |
| Development UI  | MIT         | https://github.com/binuud/avatar-dev-ui                   |
| Communication   | MIT         | https://github.com/binuud/avatar-communication-plane      |
| Physical Plane  | MIT         | https://github.com/binuud/avatar-physical-plane           |


## Credits

