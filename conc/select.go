package conc

import (
	"context"
	"github.com/guessg/goutils/containers"
	"reflect"
)

type ValueType[T any] struct {
	Value T
	Index int
}

// Select function helps caller reading data from multiple channel on limit time,
// based on the reflect.Select.
// And caller can get result came from which channel by mapping  ValueType.Index to input channels
// Example:
//
//	sources := make([]chan int32, 10)
//	output := routine.Select(ctx, sources)
//	for v := range output {
//	 fmt.Println(v.Index, v.Value)
//	}
func Select[T any](ctx context.Context, src []<-chan T) <-chan *ValueType[T] {
	inputs := containers.InterfaceSlice(src)

	cases := make([]reflect.SelectCase, len(inputs))
	for i := range cases {
		cases[i].Dir = reflect.SelectRecv
		cases[i].Chan = reflect.ValueOf(inputs[i])
	}

	output := make(chan *ValueType[T])
	go func() {
		defer close(output)
		for ctx.Err() == nil {
			if len(cases) == 0 {
				return
			}
			i, recv, recvOK := reflect.Select(cases)
			if recvOK {
				// mapping result to input channels
				value := &ValueType[T]{
					Index: i,
					Value: recv.Interface().(T),
				}
				select {
				case <-ctx.Done():
					return
				case output <- value:
				}
			} else {
				// remove closed channel
				cases = append(cases[:i], cases[i+1:]...)
			}
		}
	}()

	return output
}
