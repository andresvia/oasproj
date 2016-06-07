package errutil

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFirstOrNil(t *testing.T) {
	Convey("<nil> si el arreglo de error esta vacío", t, func() {
		errs := []error{}
		So(FirstOrNil(errs), ShouldBeNil)
	})
	Convey("error si el arreglo de error no está vacio", t, func() {
		errs := []error{errors.New("primer error"), errors.New("segundo error")}
		So(FirstOrNil(errs), ShouldNotBeNil)
	})
}
