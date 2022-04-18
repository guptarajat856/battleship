package input

type IInputProvider interface {
	TakeInput() PlayerInput
}
