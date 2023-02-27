package fnoptions

import (
	"errors"

	pkgerrors "github.com/pkg/errors"
)

type Option[T any] func(*T) error

func (o Option[T]) ToMustOption() MustOption[T] {
	return func(t *T) {
		err := o(t)
		if err != nil {
			panic(err)
		}
	}
}

type MustOption[T any] func(*T)

func (o MustOption[T]) ToOption() Option[T] {
	return func(t *T) error {
		o(t)
		return nil
	}
}

func Apply[T any](input T, options ...Option[T]) error {
	var errs []error

	for _, o := range options {
		errs = append(errs, o(&input))
	}

	return pkgerrors.WithStack(errors.Join(errs...))
}

func MustApply[T any](input T, options ...MustOption[T]) {
	for _, o := range options {
		o(&input)
	}
}
