package project

import (
	"flag"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/urfave/cli.v1"
	"testing"
)

func TestLoadProject(t *testing.T) {

	Convey("Dado un archivo descriptor de proyecto inválido", t, func() {
		Convey("Debe entrar en panico", func() {
			So(func() { LoadProject("../test-data/project/proyecto_invalido") }, ShouldPanic)
		})

	})

	Convey("Dado un archivo descriptor de proyecto inexistente", t, func() {
		Convey("Debe entrar en panico", func() {
			So(func() { LoadProject("../test-data/project/proyecto_inexsistente") }, ShouldPanic)
		})

	})

	Convey("Dado un projecto existente", t, func() {
		p := LoadProject("../test-data/project/proyecto_valido")
		Convey("Se reconocen sus campos", func() {
			So(p.Project_name, ShouldEqual, "Proyecto Válido")
			So(p.Project_description, ShouldEqual, "Este es un proyecto de prueba válido")
			So(p.Project_purpose, ShouldEqual, "Probar que al cargar este proyecto no se entra en pánico")
			So(p.Programming_language, ShouldEqual, "Go")
			So(p.Organizational_unit, ShouldEqual, "Interno")
			So(p.Package_dependencies, ShouldHaveSameTypeAs, map[string][]string{})
		})
		Convey("Tiene una representación textual correcta", func() {
			So(fmt.Sprintf("%s", p), ShouldNotBeEmpty)
			So(fmt.Sprintf("%s", p), ShouldNotStartWith, "TEMPLATE_ERROR=")
		})

	})

}

func TestMetadataExists(t *testing.T) {
	Convey("Dado un projecto", t, func() {
		Convey("Detectar existencia de metadatos", func() {
			So(MetadataExists("../test-data/project/proyecto_valido"), ShouldEqual, true)
		})
		Convey("Detectar usencia de metadatos", func() {
			So(MetadataExists("../test-data/project/proyecto_inexistente"), ShouldEqual, false)
		})
	})
}

func TestNew(t *testing.T) {
	Convey("Dado un Context de cli con un FlagSet determinado", t, func() {
		fs := flag.NewFlagSet("", flag.ContinueOnError)
		fs.String("name", "project_name", "")
		fs.String("desc", "project_description", "")
		ctx := cli.NewContext(cli.NewApp(), fs, nil)
		Convey("El Context no debe ser nil", func() {
			So(ctx, ShouldNotBeNil)
		})
		Convey("El Project tiene las propiedades este Context", func() {
			p := New(ctx)
			Convey("Tales como el nombre del projecto", func() {
				So(p.Project_name, ShouldEqual, "project_name")
			})
			Convey("y la descripción del projecto", func() {
				So(p.Project_description, ShouldEqual, "project_description")
			})
		})
	})
}

func TestWriteFile(t *testing.T) {
	Convey("Dado un Context de cli con un FlagSet determinado", t, func() {
		fs := flag.NewFlagSet("", flag.ContinueOnError)
		fs.String("name", "project_name", "")
		fs.String("desc", "project_description", "")
		ctx := cli.NewContext(cli.NewApp(), fs, nil)
		Convey("El Context no debe ser nil", func() {
			So(ctx, ShouldNotBeNil)
		})
		Convey("Y se puede escribir un archivo de descripcion del proyecto sin problemas", func() {
			p := New(ctx)
			err := p.WriteFile("../test-data/project/proyecto_TestWriteFile")
			So(err, ShouldBeNil)
		})
	})
}
