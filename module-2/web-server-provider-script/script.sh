#!/bin/bash

apt update && \
        apt upgrade -y

apt install unzip -y
apt install apache2 -y

wget https://github.com/denilsonbonatti/linux-site-dio/archive/refs/heads/main.zip -P /tmp

unzip /tmp/main.zip && \
        mv -f ./linux-site-dio*/* /var/www/html/

rm /tmp/main.zip
rm -rf ./linux-site-dio*

