#!/bin/bash

useradd $1 -c "$2" -m -s /bin/bash -p $(openssl passwd $3) -G $4

echo "user $1 created..."
