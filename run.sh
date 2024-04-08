#!/usr/bin/env sh

echo "# << PRE-3.3.0 >> :"
echo "PANIC IS NOT HIDDEN, DEVELOPER UNDERSTANDS FAILURE REASON (PANIC IN TESTED CODE)"
pushd pre-3.3.0
go test -count=1 -v ./minipanic
popd

echo ""; echo ""

echo "# << AFTER-3.3.0 >> :"
echo "PANIC IS HIDDEN, DEVELOPER GOT OBSCURE ERROR, HAPPY DEBUGGING!"
pushd after-3.3.0
go test -count=1 -v ./minipanic
popd
