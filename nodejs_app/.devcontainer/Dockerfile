# 次のイメージは deprecatedのようだ。
# https://github.com/microsoft/vscode-dev-containers
# → deprecatedとなっている。
# こちらが新 https://github.com/devcontainers
#FROM mcr.microsoft.com/vscode/devcontainers/javascript-node:16
FROM mcr.microsoft.com/devcontainers/javascript-node:20

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends nginx
