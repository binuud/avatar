# avatar

[Documentation](https://binuud.com/docs/avatar)

Please refer to above link for the latest documentation.

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

Pull the docker image and start the container. The docker image has a very small footprint.

```bash
docker pull binuud/avatar:latest
```

```bash
mkdir dataVolume
docker run --rm  -v ./dataVolume:/dataVolume \
  -e FILE_STORE_PATH=/dataVolume \
  -p 9090:9090 -p 9080:9080 binuud/avatar:latest
```
(Tested and working on version 0.1.0)

Access from UI
```bash
http://localhost:9080
```

##### Prebuilt Binaries

[Prebuilt binaries are available here](https://github.com/binuud/avatar/releases)

##### Using Go Install

```bash
go install github.com/binuud/avatar/cmd/avatar@latest
```

```bash
mkdir dataVolume
export FILE_STORE_PATH=`pwd`/dataVolume
export SERVABLE_DIRECTORY=`pwd`/dataVolume/static
avatar
```
(Tested and working on version 0.1.1)


## Local Models

For local models

* llama.cpp
* ollama
* mlx serve (only on macos)

##### Build from source

```bash
git clone https://github.com/binuud/avatar.git
cd avatar
make vendor
make build
```
(Tested and working on version 0.1.0)

##### Supported Environment Variables

* LLM_HOST - The LLM Host, the default value connects to port 11234 (mlx serve)
* MQTT_HOST - Not required in this release
* LLM_MODEL - The default LLM Model
* IMAGE_MODEL - The default Image model
* VISION_MODEL - The default vision model
* FILE_STORE_PATH - The path where the memories, avatars, generated assets are stored
* SERVABLE_DIRECTORY - The path from which static files can be served

Sample data for supported Environment variables
```
export LLM_HOST="http://localhost:11234"
export MQTT_HOST="tcp://localhost:1883"

export LLM_MODEL="lmstudio-community/Qwen3.5-9B-MLX-4bit"
export IMAGE_MODEL="Runpod/FLUX.2-klein-4B-mflux-4bit"
export VISION_MODEL="lmstudio-community/Qwen3.5-9B-MLX-4bit"

export FILE_STORE_PATH=$PWD/dataVolume
export SERVABLE_DIRECTORY=$PWD/dataVolume/static
```

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

