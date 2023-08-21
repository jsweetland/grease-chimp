#!/bin/bash

set -x

script_path=$(dirname $0)
orig_path=$(pwd)

cd ${script_path}

source ./docker-info.sh

docker run -it -p "5432:5432" ${image_tag}

cd ${orig_path}
