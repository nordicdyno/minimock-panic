#!/usr/bin/env sh

echo "# << PRE-3.3.0 >> :"
echo "PANIC IS NOT HIDDEN, DEVELOPER UNDERSTANDS FAILURE REASON (PANIC IN TESTED CODE)"
pushd pre-3.3.0
go test -count=1 -v ./minipanic
popd

echo ""; echo ""
echo "# << 3.3.0 >> :"
echo "PANIC IS HIDDEN, DEVELOPER GOT OBSCURE ERROR, HAPPY DEBUGGING!"
pushd after-3.3.0
go test -count=1 -v ./minipanic
popd

echo ""; echo ""
echo "# << 3.3.12 (fix) >> :"
echo "BAD BEHAVIOUR FIXED – HOORAY! https://github.com/gojuno/minimock/issues/91"
pushd fix-3.3.12
go test -count=1 -v ./minipanic
popd


echo ""; echo ""
echo "# << 3.4.0 (latest) >> :"
echo "BUG RETURNS – NOT SO HOORAY :("
pushd fix-3.4.0
go test -count=1 -v ./minipanic
popd
