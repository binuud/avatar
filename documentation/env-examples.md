# Env Settings

### Supported Environment Variables

* LLM_HOST - The LLM Host, the default value connects to port 11234 (mlx serve)
* MQTT_HOST - Not required in this release
* LLM_MODEL - The default LLM Model
* IMAGE_MODEL - The default Image model
* VISION_MODEL - The default vision model
* FILE_STORE_PATH - The path where the memories, avatars, generated assets are stored
* SERVABLE_DIRECTORY - The path from which static files can be served

## Env variables for file paths
```bash
export FILE_STORE_PATH=$PWD/dataVolume
export SERVABLE_DIRECTORY=$PWD/dataVolume/static
```

Sample data for supported Environment variables

## Env Variable settings for different vendors

### Local MLX-Serve

```bash
export LLM_ENDPOINT="http://localhost:11234/v1"
export LLM_MODEL="lmstudio-community/Qwen3.5-9B-MLX-4bit"
export IMAGE_MODEL="Runpod/FLUX.2-klein-4B-mflux-4bit"
export VISION_MODEL="lmstudio-community/Qwen3.5-9B-MLX-4bit"
```

### openrouter.ai

```bash
export LLM_ENDPOINT="https://openrouter.ai/api/v1"
export LLM_MODEL="qwen/qwen3.5-9b-20260310"
export IMAGE_MODEL="black-forest-labs/flux-3-image-klein"
export VISION_MODEL="qwen/qwen3.5-9b-20260310"
```