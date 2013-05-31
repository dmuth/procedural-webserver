#!/bin/bash
#
# This script just provides a bit of a quick smoke test to make sure 
# that things are working as they should be.  It's mostly intended
# for development.
#

#
# Errors are fatal
#
set -e

#
# Nuke all running instances of Go (slightly unsafe!)
#
killall a.out || true

go run ./main.go --debug-level trace &

#
# Wait for the server to start up
#
echo "Waiting for server to start..."
sleep 2

function get() {
	URL=$1
	curl -I $1

}

get http://localhost:8080/
get http://localhost:8080/?code=404&delay=.5s
get http://localhost:8080/foobar?delay=1000ms&code=201
get http://localhost:8080/12345?delay=1

echo "Waiting for server to finish..."
sleep 3

#
# Nuke all running instances of Go (slightly unsafe!)
#
killall a.out


