#!/bin/bash

apt-get update && apt-get upgrade -y

curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

