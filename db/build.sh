#!/bin/bash

set -x

script_path=$(dirname $0)
orig_path=$(pwd)

cd ${script_path}

source ./docker-info.sh

docker build -t ${image_tag} -f ${dockerfile} .

cd ${orig_path}
