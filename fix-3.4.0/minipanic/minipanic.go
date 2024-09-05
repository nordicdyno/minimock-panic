package minipanic

import "fmt"

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0 -g -i github.com/nordicdyno/minimock-panic/fix-3.4.0/minipanic.Getter -o ./getter_minimock.go

type Getter interface {
	Get() error
	Get2()
}

type GetterCaller struct {
	Getter
}

func (gc GetterCaller) Run() {
	if err := gc.Getter.Get(); err != nil {
		panic(err)
	}
	gc.Getter.Get2()
}

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6 -g -i github.com/nordicdyno/minimock-panic/after-3.3.0/minipanic.Stringer -o ./stringer_minimock.go

type Stringer interface {
	String() string
}

type Printer struct {
	Stringer
}

func (gc Printer) String() {
	fmt.Println(gc.Stringer.String())
	panic("panic will be hidden by failed you")
}
