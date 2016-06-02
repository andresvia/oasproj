package cliutil

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/urfave/cli.v1"
	"testing"
)

func TestProjectHome(t *testing.T) {
	Convey("Dado un cli Context", t, func() {
		ctx := cli.NewContext(cli.NewApp(), flag.NewFlagSet("for-test-1", flag.ContinueOnError), nil)
		Convey("El 'Home' del proyecto es el directorio punto '.'", func() {
			So(ProjectHome(ctx), ShouldEqual, ".")
		})
		fs := flag.NewFlagSet("for-test-2", flag.ContinueOnError)
		fs.Parse([]string{"test-home", "ignored"})
		ctx = cli.NewContext(cli.NewApp(), fs, nil)
		Convey("O el 'Home' del proyecto es el directorio pasado en el primer argumento", func() {
			So(ProjectHome(ctx), ShouldEqual, "test-home")
		})

	})
}
