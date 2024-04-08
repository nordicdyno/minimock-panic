package minipanic

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.2.2 -g -i github.com/nordicdyno/minimock-panic/pre-3.3.0/minipanic.Getter -o ./getter_minimock.go
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
