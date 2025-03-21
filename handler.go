package lab2

import (
	"bufio"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)
	expression, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	expression = strings.TrimSpace(expression)

	result, err := PrefixToInfix(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result))
	return err
}
