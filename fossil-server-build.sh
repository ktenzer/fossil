#!/bin/sh

PLUGIN_DIR="/home/ktenzer/plugins"

echo "Installing Dependencies"
$GOBIN/dep ensure

echo "Running Unit Tests"
go test fossil/src/engine/util
if [ $? != 0 ]; then exit 1; fi
go test fossil/src/engine/plugins/pluginUtil
if [ $? != 0 ]; then exit 1; fi

echo "Building Shared Libraries"
go build fossil/src/engine/util
if [ $? != 0 ]; then exit 1; fi
go build fossil/src/engine/client
if [ $? != 0 ]; then exit 1; fi
go build fossil/src/engine/client/k8s
if [ $? != 0 ]; then exit 1; fi
go build fossil/src/engine/plugins/pluginUtil
if [ $? != 0 ]; then exit 1; fi

echo "Building Server Service"
go install fossil/src/engine/server
if [ $? != 0 ]; then exit 1; fi

echo "Copying default configs"
cp -r $GOPATH/src/fossil/src/cli/configs/default $GOBIN/configs/default/default
if [ $? != 0 ]; then exit 1; fi

echo "Copying startup script"
cp $GOPATH/src/fossil/fossil-server-startup.sh $GOBIN
if [ $? != 0 ]; then exit 1; fi

echo "Server build completed successfully"
