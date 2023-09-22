#!/bin/bash

apt install nfs-server -y

mkdir /mirror

sudo docker swarm init --advertise-addr=10.10.10.100
sudo docker swarm join-token worker | grep docker > /vagrant/dswarm_conf.sh

