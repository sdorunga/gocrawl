package console

type FakePrinter struct {
  Messages string
}

func (fakePrinter *FakePrinter) Print(message string) {
  fakePrinter.Messages = fakePrinter.Messages + "\n" + message
}

func (fakePrinter *FakePrinter) Verify(otherMessages string) bool {
  return fakePrinter.Messages == otherMessages
}
