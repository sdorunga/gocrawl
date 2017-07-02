package console

import (
  "fmt"
)

type Printer interface {
  Print(string)
}

type TerminalPrinter struct {
}

func (terminalPrinter *TerminalPrinter) Print(message string) {
  fmt.Println(message)
}
