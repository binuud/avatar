### Google Calendar MCP

For Demo we will be using Google Calendar MCP

- install gcloud
- gcloud auth login
- Run gcloud services enable calendar-json.googleapis.com --project=[PROJECTID]
- Run gcloud services enable calendarmcp.googleapis.com --project=[PROJECTID]

sampe configs/mcp-servers.config.yaml

```bash
configs:
  - uuid: 72bef27b-e6b6-4471-990e-25e32199cd8c
    name: mcp-astro
    mcpType: MCP_TYPE_STDIO
    command: mcp-astro
  - uuid: e6546063-11d0-4da9-b485-15d1d396c143
    name: mcp-greeter
    mcpType: MCP_TYPE_STDIO
    command: mcp-greeter
  - uuid: e95c9eb6-40e1-4b76-b882-9cd3d0bb8db4
    name: mcp-weather
    mcpType: MCP_TYPE_STDIO
    command: mcp-weather
    authType: MCP_AUTH_TYPE_NONE
  - uuid: c0aeacac-11d8-416a-9bf0-acd23521c330
    name: yahoo-finance
    mcpType: MCP_TYPE_HTTP
    authType: MCP_AUTH_TYPE_NONE
    url: https://gateway.mcpservers.org/yahoo-finance/mcp
```
