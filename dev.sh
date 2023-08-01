#!/bin/bash


# all in one dev script. Creates cluster, deploys, and tears down all in one command.
pushd $(dirname $0)

./devcluster/setup.sh
skaffold dev
./devcluster/teardown.sh

popd