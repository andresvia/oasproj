package command

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckShow(t *testing.T) {
	test_funcs := []func(ctx *cli.Context) error{Show, Check}
	for _, test_func := range test_funcs {
		Convey("Dado un Context de cli con un FlagSet vacío", t, func() {
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			Convey("Para un proyecto existente", func() {
				fs.Parse([]string{"../test-data/project/proyecto_valido"})
				ctx := cli.NewContext(cli.NewApp(), fs, nil)
				Convey("No debe causar error", func() {
					So(test_func(ctx), ShouldBeNil)
				})
			})
			Convey("Para un proyecto inexistente", func() {
				fs.Parse([]string{"../test-data/project/proyecto_inexistente"})
				ctx := cli.NewContext(cli.NewApp(), fs, nil)
				Convey("Debe causar error", func() {
					So(test_func(ctx), ShouldNotBeNil)
				})
			})
		})
	}
}

func TestInit(t *testing.T) {
	Convey("Dado un Context de cli con un FlagSet vacío", t, func() {
		fs := flag.NewFlagSet("", flag.ContinueOnError)
		Convey("Para iniciar un proyecto", func() {
			Convey("No debe causar error", func() {
				dir, _ := ioutil.TempDir("/tmp", "oasproj")
				fs.Parse([]string{dir + "/TestInit"})
				ctx := cli.NewContext(cli.NewApp(), fs, nil)
				So(Init(ctx), ShouldEqual, nil)
				os.RemoveAll(dir)
			})
			Convey("A menos que se trate de un proyecto ya iniciado", func() {
				fs.Parse([]string{"../test-data/project/proyecto_valido"})
				ctx := cli.NewContext(cli.NewApp(), fs, nil)
				So(Init(ctx), ShouldNotEqual, nil)
			})
		})
	})
}
