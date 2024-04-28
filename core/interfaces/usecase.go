package interfaces

type Command interface {
	NewCommand()
}

type Result interface {
	NewResult()
}

type UseCase[T Command, Out Result] interface {
	Handle(Command) Result
}
