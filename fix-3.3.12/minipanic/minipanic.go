package minipanic

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.3.12 -g -i github.com/nordicdyno/minimock-panic/fix-3.3.12/minipanic.Getter -o ./getter_minimock.go

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
