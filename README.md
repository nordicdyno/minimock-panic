`./run.sh` output:

```shell
# << PRE-3.3.0 >> :
## PANIC IS NOT HIDDEN, DEVELOPER UNDERSTANDS FAILURE REASON (PANIC IN TESTED CODE)
=== RUN   TestPanic
--- FAIL: TestPanic (0.00s)
panic: fail [recovered]
	panic: fail

goroutine 18 [running]:
testing.tRunner.func1.2({0x1007fff80, 0x14000104280})
	testing/testing.go:1631 +0x1c4
testing.tRunner.func1()
	testing/testing.go:1634 +0x33c
panic({0x1007fff80?, 0x14000104280?})
	runtime/panic.go:770 +0x124
github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic.GetterCaller.Run({{0x100824608?, 0x14000110500?}})
	github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic/minipanic.go:15 +0x60
github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic.TestPanic(0x14000126680)
	github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic/minipanic_test.go:17 +0x11c
testing.tRunner(0x14000126680, 0x1008230a0)
	testing/testing.go:1689 +0xec
created by testing.(*T).Run in goroutine 1
	testing/testing.go:1742 +0x318
FAIL	github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic	0.307s
FAIL


# << 3.3.0 >> :
## PANIC IS HIDDEN, DEVELOPER GOT OBSCURE ERROR, HAPPY DEBUGGING!
=== RUN   TestPanic
    safe_tester.go:19: Expected call to GetterMock.Get2
--- FAIL: TestPanic (0.00s)
FAIL
FAIL	github.com/nordicdyno/minimock-panic/after-3.3.0/minipanic	0.315s
FAIL


# << 3.3.12 (fix) >> :
## BAD BEHAVIOUR FIXED – HOORAY! https://github.com/gojuno/minimock/issues/91
=== RUN   TestPanic
    safe_tester.go:19: Expected call to GetterMock.Get2
--- FAIL: TestPanic (0.00s)
panic: fail [recovered]
	panic: fail

goroutine 34 [running]:
testing.tRunner.func1.2({0x1007c8120, 0x140001042b0})
	testing/testing.go:1631 +0x1c4
testing.tRunner.func1()
	testing/testing.go:1634 +0x33c
panic({0x1007c8120?, 0x140001042b0?})
	runtime/panic.go:770 +0x124
github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic.GetterCaller.Run({{0x1007ec9e8?, 0x1400011ec30?}})
	github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic/minipanic.go:16 +0x60
github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic.TestPanic(0x14000126820)
	github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic/minipanic_test.go:17 +0x9c
testing.tRunner(0x14000126820, 0x1007eb480)
	testing/testing.go:1689 +0xec
created by testing.(*T).Run in goroutine 1
	testing/testing.go:1742 +0x318
FAIL	github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic	0.300s
FAIL


# << 3.4.0 (latest) >> :
## BUG RETURNS – NOT SO HOORAY :(
=== RUN   TestPanic
    minipanic_test.go:13: fail
    getter_minimock.go:401: Expected call to GetterMock.Get2 at
        github.com/nordicdyno/minimock-panic/fix-3.4.0/minipanic/minipanic_test.go:21
--- FAIL: TestPanic (0.00s)
FAIL
FAIL	github.com/nordicdyno/minimock-panic/fix-3.4.0/minipanic	0.324s
FAIL
```

